/*
Copyright (c) 2020-2021 TriggerMesh Inc.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/triggermesh/aws-event-sources/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSourcesV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSourcesV1alpha1) AWSCloudWatchLogsSources(namespace string) v1alpha1.AWSCloudWatchLogsSourceInterface {
	return &FakeAWSCloudWatchLogsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCloudWatchSources(namespace string) v1alpha1.AWSCloudWatchSourceInterface {
	return &FakeAWSCloudWatchSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCodeCommitSources(namespace string) v1alpha1.AWSCodeCommitSourceInterface {
	return &FakeAWSCodeCommitSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCognitoIdentitySources(namespace string) v1alpha1.AWSCognitoIdentitySourceInterface {
	return &FakeAWSCognitoIdentitySources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCognitoUserPoolSources(namespace string) v1alpha1.AWSCognitoUserPoolSourceInterface {
	return &FakeAWSCognitoUserPoolSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSDynamoDBSources(namespace string) v1alpha1.AWSDynamoDBSourceInterface {
	return &FakeAWSDynamoDBSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSIoTSources(namespace string) v1alpha1.AWSIoTSourceInterface {
	return &FakeAWSIoTSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSKinesisSources(namespace string) v1alpha1.AWSKinesisSourceInterface {
	return &FakeAWSKinesisSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSSNSSources(namespace string) v1alpha1.AWSSNSSourceInterface {
	return &FakeAWSSNSSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSSQSSources(namespace string) v1alpha1.AWSSQSSourceInterface {
	return &FakeAWSSQSSources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSourcesV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
