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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	pkgapis "knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/triggermesh/aws-event-sources/pkg/apis"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (s *AWSCognitoIdentitySource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AWSCognitoIdentitySource")
}

// GetUntypedSpec implements apis.HasSpec.
func (s *AWSCognitoIdentitySource) GetUntypedSpec() interface{} {
	return s.Spec
}

// GetConditionSet implements duckv1.KRShaped.
func (s *AWSCognitoIdentitySource) GetConditionSet() pkgapis.ConditionSet {
	return awsEventSourceConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *AWSCognitoIdentitySource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements AWSEventSource.
func (s *AWSCognitoIdentitySource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetARN implements AWSEventSource.
func (s *AWSCognitoIdentitySource) GetARN() apis.ARN {
	return s.Spec.ARN
}

// GetSourceStatus implements AWSEventSource.
func (s *AWSCognitoIdentitySource) GetSourceStatus() *AWSEventSourceStatus {
	return &s.Status
}

// Supported event types
const (
	AWSCognitoIdentityGenericEventType = "sync_trigger"
)

// GetEventTypes implements AWSEventSource.
func (s *AWSCognitoIdentitySource) GetEventTypes() []string {
	return []string{
		AWSCognitoIdentityGenericEventType,
	}
}
