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

// AWSIoTSourcesGetter has a method to return a AWSIoTSourceInterface.
// A group's client should implement this interface.
type AWSIoTSourcesGetter interface {
	AWSIoTSources(namespace string) AWSIoTSourceInterface
}

// AWSIoTSourceInterface has methods to work with AWSIoTSource resources.
type AWSIoTSourceInterface interface {
	Create(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.CreateOptions) (*v1alpha1.AWSIoTSource, error)
	Update(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.UpdateOptions) (*v1alpha1.AWSIoTSource, error)
	UpdateStatus(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.UpdateOptions) (*v1alpha1.AWSIoTSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSIoTSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSIoTSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSIoTSource, err error)
	AWSIoTSourceExpansion
}

// aWSIoTSources implements AWSIoTSourceInterface
type aWSIoTSources struct {
	client rest.Interface
	ns     string
}

// newAWSIoTSources returns a AWSIoTSources
func newAWSIoTSources(c *SourcesV1alpha1Client, namespace string) *aWSIoTSources {
	return &aWSIoTSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSIoTSource, and returns the corresponding aWSIoTSource object, and an error if there is any.
func (c *aWSIoTSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSIoTSource, err error) {
	result = &v1alpha1.AWSIoTSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiotsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSIoTSources that match those selectors.
func (c *aWSIoTSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSIoTSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSIoTSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiotsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSIoTSources.
func (c *aWSIoTSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiotsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSIoTSource and creates it.  Returns the server's representation of the aWSIoTSource, and an error, if there is any.
func (c *aWSIoTSources) Create(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.CreateOptions) (result *v1alpha1.AWSIoTSource, err error) {
	result = &v1alpha1.AWSIoTSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiotsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSIoTSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSIoTSource and updates it. Returns the server's representation of the aWSIoTSource, and an error, if there is any.
func (c *aWSIoTSources) Update(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.UpdateOptions) (result *v1alpha1.AWSIoTSource, err error) {
	result = &v1alpha1.AWSIoTSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiotsources").
		Name(aWSIoTSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSIoTSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSIoTSources) UpdateStatus(ctx context.Context, aWSIoTSource *v1alpha1.AWSIoTSource, opts v1.UpdateOptions) (result *v1alpha1.AWSIoTSource, err error) {
	result = &v1alpha1.AWSIoTSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiotsources").
		Name(aWSIoTSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSIoTSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSIoTSource and deletes it. Returns an error if one occurs.
func (c *aWSIoTSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiotsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSIoTSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiotsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSIoTSource.
func (c *aWSIoTSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSIoTSource, err error) {
	result = &v1alpha1.AWSIoTSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiotsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
