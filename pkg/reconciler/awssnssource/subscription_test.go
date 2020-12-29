/*
Copyright (c) 2020 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awssnssource

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/triggermesh/aws-event-sources/pkg/apis/sources/v1alpha1"
)

func TestAwsCredentials(t *testing.T) {
	const (
		ns = "fake-namespace"

		accessKeyIDKey     = "key-id"
		accessKeyIDVal     = "fake key ID"
		secretAccessKeyKey = "secret-key"
		secretAccessKeyVal = "fake secret"
	)

	testCases := []struct {
		name        string
		initSecrets []*corev1.Secret
		input       v1alpha1.AWSSecurityCredentials
		expect      *credentials.Value
		getRequests int
	}{
		{
			name: "Both from value",
			input: v1alpha1.AWSSecurityCredentials{
				AccessKeyID: v1alpha1.ValueFromField{
					Value: accessKeyIDVal,
				},
				SecretAccessKey: v1alpha1.ValueFromField{
					Value: secretAccessKeyVal,
				},
			},
			expect: &credentials.Value{
				AccessKeyID:     accessKeyIDVal,
				SecretAccessKey: secretAccessKeyVal,
			},
			getRequests: 0,
		},
		{
			name: "One from value, the other from secret",
			initSecrets: []*corev1.Secret{
				newSecret(ns, "secret1", map[string]string{
					secretAccessKeyKey: secretAccessKeyVal,
				}),
			},
			input: v1alpha1.AWSSecurityCredentials{
				AccessKeyID: v1alpha1.ValueFromField{
					Value: accessKeyIDVal,
				},
				SecretAccessKey: v1alpha1.ValueFromField{
					ValueFromSecret: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret1",
						},
						Key: secretAccessKeyKey,
					},
				},
			},
			expect: &credentials.Value{
				AccessKeyID:     accessKeyIDVal,
				SecretAccessKey: secretAccessKeyVal,
			},
			getRequests: 1,
		},
		{
			name: "Both from same secret",
			initSecrets: []*corev1.Secret{
				newSecret(ns, "secret1", map[string]string{
					accessKeyIDKey:     accessKeyIDVal,
					secretAccessKeyKey: secretAccessKeyVal,
				}),
			},
			input: v1alpha1.AWSSecurityCredentials{
				AccessKeyID: v1alpha1.ValueFromField{
					ValueFromSecret: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret1",
						},
						Key: accessKeyIDKey,
					},
				},
				SecretAccessKey: v1alpha1.ValueFromField{
					ValueFromSecret: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret1",
						},
						Key: secretAccessKeyKey,
					},
				},
			},
			expect: &credentials.Value{
				AccessKeyID:     accessKeyIDVal,
				SecretAccessKey: secretAccessKeyVal,
			},
			getRequests: 1,
		},
		{
			name: "Both from a different secret",
			initSecrets: []*corev1.Secret{
				newSecret(ns, "secret1", map[string]string{
					accessKeyIDKey: accessKeyIDVal,
				}),
				newSecret(ns, "secret2", map[string]string{
					secretAccessKeyKey: secretAccessKeyVal,
				}),
			},
			input: v1alpha1.AWSSecurityCredentials{
				AccessKeyID: v1alpha1.ValueFromField{
					ValueFromSecret: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret1",
						},
						Key: accessKeyIDKey,
					},
				},
				SecretAccessKey: v1alpha1.ValueFromField{
					ValueFromSecret: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret2",
						},
						Key: secretAccessKeyKey,
					},
				},
			},
			expect: &credentials.Value{
				AccessKeyID:     accessKeyIDVal,
				SecretAccessKey: secretAccessKeyVal,
			},
			getRequests: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			secrets := make([]runtime.Object, len(tc.initSecrets))
			for i, secret := range tc.initSecrets {
				secrets[i] = secret
			}

			cli := fake.NewSimpleClientset(secrets...)

			creds, err := awsCredentials(cli.CoreV1().Secrets(ns), &tc.input)

			require.NoError(t, err)

			assert.Equal(t, tc.expect.AccessKeyID, creds.AccessKeyID, "Value of access key ID")
			assert.Equal(t, tc.expect.SecretAccessKey, creds.SecretAccessKey, "Value of secret access key")
			assert.Equal(t, tc.getRequests, len(cli.Actions()), "Number of API requests")
		})
	}
}

func TestErrors(t *testing.T) {
	genericErr := assert.AnError
	genericAWSErr := awserr.New("TestError", "an error", assert.AnError)
	genericK8SErr := &apierrors.StatusError{}

	t.Run("AWS error", func(t *testing.T) {
		assert.True(t, isAWSError(genericAWSErr))
		assert.True(t, isAWSError(fmt.Errorf("wrapped: %w", genericAWSErr)))
		assert.False(t, isAWSError(genericErr))
	})

	t.Run("denied", func(t *testing.T) {
		deniedErr := awserr.New(sns.ErrCodeAuthorizationErrorException, "an error", assert.AnError)

		assert.True(t, isDenied(deniedErr))
		assert.True(t, isDenied(fmt.Errorf("wrapped: %w", deniedErr)))
		assert.False(t, isDenied(genericAWSErr))
		assert.False(t, isDenied(genericErr))
	})

	t.Run("not found", func(t *testing.T) {
		notFoundAWSErr := awserr.New(sns.ErrCodeNotFoundException, "an error", assert.AnError)
		notFoundK8SErr := apierrors.NewNotFound(schema.GroupResource{}, "")

		assert.True(t, isNotFound(notFoundAWSErr))
		assert.True(t, isNotFound(fmt.Errorf("wrapped: %w", notFoundAWSErr)))
		assert.True(t, isNotFound(notFoundK8SErr))
		assert.True(t, isNotFound(fmt.Errorf("wrapped: %w", notFoundK8SErr)))
		assert.False(t, isNotFound(genericAWSErr))
		assert.False(t, isNotFound(genericK8SErr))
		assert.False(t, isNotFound(genericErr))
	})
}

func newSecret(ns, name string, data map[string]string) *corev1.Secret {
	secr := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
		Data: make(map[string][]byte, len(data)),
	}

	for k, v := range data {
		secr.Data[k] = []byte(v)
	}

	return secr
}
