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

// FakeAWSCognitoUserPoolSources implements AWSCognitoUserPoolSourceInterface
type FakeAWSCognitoUserPoolSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var awscognitouserpoolsourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "awscognitouserpoolsources"}

var awscognitouserpoolsourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "AWSCognitoUserPoolSource"}

// Get takes name of the aWSCognitoUserPoolSource, and returns the corresponding aWSCognitoUserPoolSource object, and an error if there is any.
func (c *FakeAWSCognitoUserPoolSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscognitouserpoolsourcesResource, c.ns, name), &v1alpha1.AWSCognitoUserPoolSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), err
}

// List takes label and field selectors, and returns the list of AWSCognitoUserPoolSources that match those selectors.
func (c *FakeAWSCognitoUserPoolSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSCognitoUserPoolSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscognitouserpoolsourcesResource, awscognitouserpoolsourcesKind, c.ns, opts), &v1alpha1.AWSCognitoUserPoolSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSCognitoUserPoolSourceList{ListMeta: obj.(*v1alpha1.AWSCognitoUserPoolSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSCognitoUserPoolSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSCognitoUserPoolSources.
func (c *FakeAWSCognitoUserPoolSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscognitouserpoolsourcesResource, c.ns, opts))

}

// Create takes the representation of a aWSCognitoUserPoolSource and creates it.  Returns the server's representation of the aWSCognitoUserPoolSource, and an error, if there is any.
func (c *FakeAWSCognitoUserPoolSources) Create(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.CreateOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscognitouserpoolsourcesResource, c.ns, aWSCognitoUserPoolSource), &v1alpha1.AWSCognitoUserPoolSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), err
}

// Update takes the representation of a aWSCognitoUserPoolSource and updates it. Returns the server's representation of the aWSCognitoUserPoolSource, and an error, if there is any.
func (c *FakeAWSCognitoUserPoolSources) Update(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscognitouserpoolsourcesResource, c.ns, aWSCognitoUserPoolSource), &v1alpha1.AWSCognitoUserPoolSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSCognitoUserPoolSources) UpdateStatus(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (*v1alpha1.AWSCognitoUserPoolSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awscognitouserpoolsourcesResource, "status", c.ns, aWSCognitoUserPoolSource), &v1alpha1.AWSCognitoUserPoolSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), err
}

// Delete takes name of the aWSCognitoUserPoolSource and deletes it. Returns an error if one occurs.
func (c *FakeAWSCognitoUserPoolSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscognitouserpoolsourcesResource, c.ns, name), &v1alpha1.AWSCognitoUserPoolSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSCognitoUserPoolSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscognitouserpoolsourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSCognitoUserPoolSourceList{})
	return err
}

// Patch applies the patch and returns the patched aWSCognitoUserPoolSource.
func (c *FakeAWSCognitoUserPoolSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscognitouserpoolsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSCognitoUserPoolSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), err
}
