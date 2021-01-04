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
	"context"

	v1alpha1 "github.com/triggermesh/aws-event-sources/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAWSCloudWatchLogsSources implements AWSCloudWatchLogsSourceInterface
type FakeAWSCloudWatchLogsSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var awscloudwatchlogssourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "awscloudwatchlogssources"}

var awscloudwatchlogssourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "AWSCloudWatchLogsSource"}

// Get takes name of the aWSCloudWatchLogsSource, and returns the corresponding aWSCloudWatchLogsSource object, and an error if there is any.
func (c *FakeAWSCloudWatchLogsSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSCloudWatchLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatchlogssourcesResource, c.ns, name), &v1alpha1.AWSCloudWatchLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCloudWatchLogsSource), err
}

// List takes label and field selectors, and returns the list of AWSCloudWatchLogsSources that match those selectors.
func (c *FakeAWSCloudWatchLogsSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSCloudWatchLogsSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatchlogssourcesResource, awscloudwatchlogssourcesKind, c.ns, opts), &v1alpha1.AWSCloudWatchLogsSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSCloudWatchLogsSourceList{ListMeta: obj.(*v1alpha1.AWSCloudWatchLogsSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSCloudWatchLogsSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSCloudWatchLogsSources.
func (c *FakeAWSCloudWatchLogsSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatchlogssourcesResource, c.ns, opts))

}

// Create takes the representation of a aWSCloudWatchLogsSource and creates it.  Returns the server's representation of the aWSCloudWatchLogsSource, and an error, if there is any.
func (c *FakeAWSCloudWatchLogsSources) Create(ctx context.Context, aWSCloudWatchLogsSource *v1alpha1.AWSCloudWatchLogsSource, opts v1.CreateOptions) (result *v1alpha1.AWSCloudWatchLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatchlogssourcesResource, c.ns, aWSCloudWatchLogsSource), &v1alpha1.AWSCloudWatchLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCloudWatchLogsSource), err
}

// Update takes the representation of a aWSCloudWatchLogsSource and updates it. Returns the server's representation of the aWSCloudWatchLogsSource, and an error, if there is any.
func (c *FakeAWSCloudWatchLogsSources) Update(ctx context.Context, aWSCloudWatchLogsSource *v1alpha1.AWSCloudWatchLogsSource, opts v1.UpdateOptions) (result *v1alpha1.AWSCloudWatchLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatchlogssourcesResource, c.ns, aWSCloudWatchLogsSource), &v1alpha1.AWSCloudWatchLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCloudWatchLogsSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSCloudWatchLogsSources) UpdateStatus(ctx context.Context, aWSCloudWatchLogsSource *v1alpha1.AWSCloudWatchLogsSource, opts v1.UpdateOptions) (*v1alpha1.AWSCloudWatchLogsSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awscloudwatchlogssourcesResource, "status", c.ns, aWSCloudWatchLogsSource), &v1alpha1.AWSCloudWatchLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCloudWatchLogsSource), err
}

// Delete takes name of the aWSCloudWatchLogsSource and deletes it. Returns an error if one occurs.
func (c *FakeAWSCloudWatchLogsSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatchlogssourcesResource, c.ns, name), &v1alpha1.AWSCloudWatchLogsSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSCloudWatchLogsSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatchlogssourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSCloudWatchLogsSourceList{})
	return err
}

// Patch applies the patch and returns the patched aWSCloudWatchLogsSource.
func (c *FakeAWSCloudWatchLogsSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSCloudWatchLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatchlogssourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSCloudWatchLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCloudWatchLogsSource), err
}