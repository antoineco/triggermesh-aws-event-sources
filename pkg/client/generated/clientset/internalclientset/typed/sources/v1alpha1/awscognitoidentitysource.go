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
	"time"

	v1alpha1 "github.com/triggermesh/aws-event-sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/triggermesh/aws-event-sources/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSCognitoIdentitySourcesGetter has a method to return a AWSCognitoIdentitySourceInterface.
// A group's client should implement this interface.
type AWSCognitoIdentitySourcesGetter interface {
	AWSCognitoIdentitySources(namespace string) AWSCognitoIdentitySourceInterface
}

// AWSCognitoIdentitySourceInterface has methods to work with AWSCognitoIdentitySource resources.
type AWSCognitoIdentitySourceInterface interface {
	Create(*v1alpha1.AWSCognitoIdentitySource) (*v1alpha1.AWSCognitoIdentitySource, error)
	Update(*v1alpha1.AWSCognitoIdentitySource) (*v1alpha1.AWSCognitoIdentitySource, error)
	UpdateStatus(*v1alpha1.AWSCognitoIdentitySource) (*v1alpha1.AWSCognitoIdentitySource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AWSCognitoIdentitySource, error)
	List(opts v1.ListOptions) (*v1alpha1.AWSCognitoIdentitySourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AWSCognitoIdentitySource, err error)
	AWSCognitoIdentitySourceExpansion
}

// aWSCognitoIdentitySources implements AWSCognitoIdentitySourceInterface
type aWSCognitoIdentitySources struct {
	client rest.Interface
	ns     string
}

// newAWSCognitoIdentitySources returns a AWSCognitoIdentitySources
func newAWSCognitoIdentitySources(c *SourcesV1alpha1Client, namespace string) *aWSCognitoIdentitySources {
	return &aWSCognitoIdentitySources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSCognitoIdentitySource, and returns the corresponding aWSCognitoIdentitySource object, and an error if there is any.
func (c *aWSCognitoIdentitySources) Get(name string, options v1.GetOptions) (result *v1alpha1.AWSCognitoIdentitySource, err error) {
	result = &v1alpha1.AWSCognitoIdentitySource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSCognitoIdentitySources that match those selectors.
func (c *aWSCognitoIdentitySources) List(opts v1.ListOptions) (result *v1alpha1.AWSCognitoIdentitySourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSCognitoIdentitySourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSCognitoIdentitySources.
func (c *aWSCognitoIdentitySources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a aWSCognitoIdentitySource and creates it.  Returns the server's representation of the aWSCognitoIdentitySource, and an error, if there is any.
func (c *aWSCognitoIdentitySources) Create(aWSCognitoIdentitySource *v1alpha1.AWSCognitoIdentitySource) (result *v1alpha1.AWSCognitoIdentitySource, err error) {
	result = &v1alpha1.AWSCognitoIdentitySource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		Body(aWSCognitoIdentitySource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aWSCognitoIdentitySource and updates it. Returns the server's representation of the aWSCognitoIdentitySource, and an error, if there is any.
func (c *aWSCognitoIdentitySources) Update(aWSCognitoIdentitySource *v1alpha1.AWSCognitoIdentitySource) (result *v1alpha1.AWSCognitoIdentitySource, err error) {
	result = &v1alpha1.AWSCognitoIdentitySource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		Name(aWSCognitoIdentitySource.Name).
		Body(aWSCognitoIdentitySource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *aWSCognitoIdentitySources) UpdateStatus(aWSCognitoIdentitySource *v1alpha1.AWSCognitoIdentitySource) (result *v1alpha1.AWSCognitoIdentitySource, err error) {
	result = &v1alpha1.AWSCognitoIdentitySource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		Name(aWSCognitoIdentitySource.Name).
		SubResource("status").
		Body(aWSCognitoIdentitySource).
		Do().
		Into(result)
	return
}

// Delete takes name of the aWSCognitoIdentitySource and deletes it. Returns an error if one occurs.
func (c *aWSCognitoIdentitySources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSCognitoIdentitySources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aWSCognitoIdentitySource.
func (c *aWSCognitoIdentitySources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AWSCognitoIdentitySource, err error) {
	result = &v1alpha1.AWSCognitoIdentitySource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscognitoidentitysources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
