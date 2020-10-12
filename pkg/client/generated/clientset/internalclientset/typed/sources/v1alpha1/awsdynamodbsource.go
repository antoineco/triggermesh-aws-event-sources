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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/triggermesh/aws-event-sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/triggermesh/aws-event-sources/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSDynamoDBSourcesGetter has a method to return a AWSDynamoDBSourceInterface.
// A group's client should implement this interface.
type AWSDynamoDBSourcesGetter interface {
	AWSDynamoDBSources(namespace string) AWSDynamoDBSourceInterface
}

// AWSDynamoDBSourceInterface has methods to work with AWSDynamoDBSource resources.
type AWSDynamoDBSourceInterface interface {
	Create(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.CreateOptions) (*v1alpha1.AWSDynamoDBSource, error)
	Update(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.UpdateOptions) (*v1alpha1.AWSDynamoDBSource, error)
	UpdateStatus(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.UpdateOptions) (*v1alpha1.AWSDynamoDBSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSDynamoDBSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSDynamoDBSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSDynamoDBSource, err error)
	AWSDynamoDBSourceExpansion
}

// aWSDynamoDBSources implements AWSDynamoDBSourceInterface
type aWSDynamoDBSources struct {
	client rest.Interface
	ns     string
}

// newAWSDynamoDBSources returns a AWSDynamoDBSources
func newAWSDynamoDBSources(c *SourcesV1alpha1Client, namespace string) *aWSDynamoDBSources {
	return &aWSDynamoDBSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSDynamoDBSource, and returns the corresponding aWSDynamoDBSource object, and an error if there is any.
func (c *aWSDynamoDBSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSDynamoDBSource, err error) {
	result = &v1alpha1.AWSDynamoDBSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSDynamoDBSources that match those selectors.
func (c *aWSDynamoDBSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSDynamoDBSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSDynamoDBSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSDynamoDBSources.
func (c *aWSDynamoDBSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSDynamoDBSource and creates it.  Returns the server's representation of the aWSDynamoDBSource, and an error, if there is any.
func (c *aWSDynamoDBSources) Create(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.CreateOptions) (result *v1alpha1.AWSDynamoDBSource, err error) {
	result = &v1alpha1.AWSDynamoDBSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSDynamoDBSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSDynamoDBSource and updates it. Returns the server's representation of the aWSDynamoDBSource, and an error, if there is any.
func (c *aWSDynamoDBSources) Update(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.UpdateOptions) (result *v1alpha1.AWSDynamoDBSource, err error) {
	result = &v1alpha1.AWSDynamoDBSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		Name(aWSDynamoDBSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSDynamoDBSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSDynamoDBSources) UpdateStatus(ctx context.Context, aWSDynamoDBSource *v1alpha1.AWSDynamoDBSource, opts v1.UpdateOptions) (result *v1alpha1.AWSDynamoDBSource, err error) {
	result = &v1alpha1.AWSDynamoDBSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		Name(aWSDynamoDBSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSDynamoDBSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSDynamoDBSource and deletes it. Returns an error if one occurs.
func (c *aWSDynamoDBSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSDynamoDBSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSDynamoDBSource.
func (c *aWSDynamoDBSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSDynamoDBSource, err error) {
	result = &v1alpha1.AWSDynamoDBSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdynamodbsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
