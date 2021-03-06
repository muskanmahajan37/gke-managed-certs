/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/GoogleCloudPlatform/gke-managed-certs/pkg/apis/networking.gke.io/v1"
	scheme "github.com/GoogleCloudPlatform/gke-managed-certs/pkg/clientgen/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagedCertificatesGetter has a method to return a ManagedCertificateInterface.
// A group's client should implement this interface.
type ManagedCertificatesGetter interface {
	ManagedCertificates(namespace string) ManagedCertificateInterface
}

// ManagedCertificateInterface has methods to work with ManagedCertificate resources.
type ManagedCertificateInterface interface {
	Create(*v1.ManagedCertificate) (*v1.ManagedCertificate, error)
	Update(*v1.ManagedCertificate) (*v1.ManagedCertificate, error)
	UpdateStatus(*v1.ManagedCertificate) (*v1.ManagedCertificate, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ManagedCertificate, error)
	List(opts metav1.ListOptions) (*v1.ManagedCertificateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ManagedCertificate, err error)
	ManagedCertificateExpansion
}

// managedCertificates implements ManagedCertificateInterface
type managedCertificates struct {
	client rest.Interface
	ns     string
}

// newManagedCertificates returns a ManagedCertificates
func newManagedCertificates(c *NetworkingV1Client, namespace string) *managedCertificates {
	return &managedCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managedCertificate, and returns the corresponding managedCertificate object, and an error if there is any.
func (c *managedCertificates) Get(name string, options metav1.GetOptions) (result *v1.ManagedCertificate, err error) {
	result = &v1.ManagedCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagedCertificates that match those selectors.
func (c *managedCertificates) List(opts metav1.ListOptions) (result *v1.ManagedCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ManagedCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managedCertificates.
func (c *managedCertificates) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managedcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a managedCertificate and creates it.  Returns the server's representation of the managedCertificate, and an error, if there is any.
func (c *managedCertificates) Create(managedCertificate *v1.ManagedCertificate) (result *v1.ManagedCertificate, err error) {
	result = &v1.ManagedCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managedcertificates").
		Body(managedCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a managedCertificate and updates it. Returns the server's representation of the managedCertificate, and an error, if there is any.
func (c *managedCertificates) Update(managedCertificate *v1.ManagedCertificate) (result *v1.ManagedCertificate, err error) {
	result = &v1.ManagedCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedcertificates").
		Name(managedCertificate.Name).
		Body(managedCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *managedCertificates) UpdateStatus(managedCertificate *v1.ManagedCertificate) (result *v1.ManagedCertificate, err error) {
	result = &v1.ManagedCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedcertificates").
		Name(managedCertificate.Name).
		SubResource("status").
		Body(managedCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the managedCertificate and deletes it. Returns an error if one occurs.
func (c *managedCertificates) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managedCertificates) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched managedCertificate.
func (c *managedCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ManagedCertificate, err error) {
	result = &v1.ManagedCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managedcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
