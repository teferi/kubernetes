/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	api "k8s.io/kubernetes/pkg/api"
	servicecatalog "k8s.io/kubernetes/pkg/apis/servicecatalog"
	watch "k8s.io/kubernetes/pkg/watch"
)

// CatalogPostingsGetter has a method to return a CatalogPostingInterface.
// A group's client should implement this interface.
type CatalogPostingsGetter interface {
	CatalogPostings(namespace string) CatalogPostingInterface
}

// CatalogPostingInterface has methods to work with CatalogPosting resources.
type CatalogPostingInterface interface {
	Create(*servicecatalog.CatalogPosting) (*servicecatalog.CatalogPosting, error)
	Update(*servicecatalog.CatalogPosting) (*servicecatalog.CatalogPosting, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*servicecatalog.CatalogPosting, error)
	List(opts api.ListOptions) (*servicecatalog.CatalogPostingList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	CatalogPostingExpansion
}

// catalogPostings implements CatalogPostingInterface
type catalogPostings struct {
	client *ServicecatalogClient
	ns     string
}

// newCatalogPostings returns a CatalogPostings
func newCatalogPostings(c *ServicecatalogClient, namespace string) *catalogPostings {
	return &catalogPostings{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a catalogPosting and creates it.  Returns the server's representation of the catalogPosting, and an error, if there is any.
func (c *catalogPostings) Create(catalogPosting *servicecatalog.CatalogPosting) (result *servicecatalog.CatalogPosting, err error) {
	result = &servicecatalog.CatalogPosting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("catalogpostings").
		Body(catalogPosting).
		Do().
		Into(result)
	return
}

// Update takes the representation of a catalogPosting and updates it. Returns the server's representation of the catalogPosting, and an error, if there is any.
func (c *catalogPostings) Update(catalogPosting *servicecatalog.CatalogPosting) (result *servicecatalog.CatalogPosting, err error) {
	result = &servicecatalog.CatalogPosting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("catalogpostings").
		Name(catalogPosting.Name).
		Body(catalogPosting).
		Do().
		Into(result)
	return
}

// Delete takes name of the catalogPosting and deletes it. Returns an error if one occurs.
func (c *catalogPostings) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogpostings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *catalogPostings) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogpostings").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the catalogPosting, and returns the corresponding catalogPosting object, and an error if there is any.
func (c *catalogPostings) Get(name string) (result *servicecatalog.CatalogPosting, err error) {
	result = &servicecatalog.CatalogPosting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogpostings").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CatalogPostings that match those selectors.
func (c *catalogPostings) List(opts api.ListOptions) (result *servicecatalog.CatalogPostingList, err error) {
	result = &servicecatalog.CatalogPostingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogpostings").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested catalogPostings.
func (c *catalogPostings) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("catalogpostings").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
