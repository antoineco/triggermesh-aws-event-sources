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

// FakeAWSSNSSources implements AWSSNSSourceInterface
type FakeAWSSNSSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var awssnssourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "awssnssources"}

var awssnssourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "AWSSNSSource"}

// Get takes name of the aWSSNSSource, and returns the corresponding aWSSNSSource object, and an error if there is any.
func (c *FakeAWSSNSSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSSNSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssnssourcesResource, c.ns, name), &v1alpha1.AWSSNSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSSNSSource), err
}

// List takes label and field selectors, and returns the list of AWSSNSSources that match those selectors.
func (c *FakeAWSSNSSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSSNSSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssnssourcesResource, awssnssourcesKind, c.ns, opts), &v1alpha1.AWSSNSSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSSNSSourceList{ListMeta: obj.(*v1alpha1.AWSSNSSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSSNSSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSSNSSources.
func (c *FakeAWSSNSSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssnssourcesResource, c.ns, opts))

}

// Create takes the representation of a aWSSNSSource and creates it.  Returns the server's representation of the aWSSNSSource, and an error, if there is any.
func (c *FakeAWSSNSSources) Create(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.CreateOptions) (result *v1alpha1.AWSSNSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssnssourcesResource, c.ns, aWSSNSSource), &v1alpha1.AWSSNSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSSNSSource), err
}

// Update takes the representation of a aWSSNSSource and updates it. Returns the server's representation of the aWSSNSSource, and an error, if there is any.
func (c *FakeAWSSNSSources) Update(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (result *v1alpha1.AWSSNSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssnssourcesResource, c.ns, aWSSNSSource), &v1alpha1.AWSSNSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSSNSSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSSNSSources) UpdateStatus(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (*v1alpha1.AWSSNSSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awssnssourcesResource, "status", c.ns, aWSSNSSource), &v1alpha1.AWSSNSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSSNSSource), err
}

// Delete takes name of the aWSSNSSource and deletes it. Returns an error if one occurs.
func (c *FakeAWSSNSSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssnssourcesResource, c.ns, name), &v1alpha1.AWSSNSSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSSNSSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssnssourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSSNSSourceList{})
	return err
}

// Patch applies the patch and returns the patched aWSSNSSource.
func (c *FakeAWSSNSSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSSNSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssnssourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSSNSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSSNSSource), err
}
