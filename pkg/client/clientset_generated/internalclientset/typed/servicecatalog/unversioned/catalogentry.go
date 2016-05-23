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

// CatalogEntriesGetter has a method to return a CatalogEntryInterface.
// A group's client should implement this interface.
type CatalogEntriesGetter interface {
	CatalogEntries() CatalogEntryInterface
}

// CatalogEntryInterface has methods to work with CatalogEntry resources.
type CatalogEntryInterface interface {
	Create(*servicecatalog.CatalogEntry) (*servicecatalog.CatalogEntry, error)
	Update(*servicecatalog.CatalogEntry) (*servicecatalog.CatalogEntry, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*servicecatalog.CatalogEntry, error)
	List(opts api.ListOptions) (*servicecatalog.CatalogEntryList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	CatalogEntryExpansion
}

// catalogEntries implements CatalogEntryInterface
type catalogEntries struct {
	client *ServicecatalogClient
}

// newCatalogEntries returns a CatalogEntries
func newCatalogEntries(c *ServicecatalogClient) *catalogEntries {
	return &catalogEntries{
		client: c,
	}
}

// Create takes the representation of a catalogEntry and creates it.  Returns the server's representation of the catalogEntry, and an error, if there is any.
func (c *catalogEntries) Create(catalogEntry *servicecatalog.CatalogEntry) (result *servicecatalog.CatalogEntry, err error) {
	result = &servicecatalog.CatalogEntry{}
	err = c.client.Post().
		Resource("catalogentries").
		Body(catalogEntry).
		Do().
		Into(result)
	return
}

// Update takes the representation of a catalogEntry and updates it. Returns the server's representation of the catalogEntry, and an error, if there is any.
func (c *catalogEntries) Update(catalogEntry *servicecatalog.CatalogEntry) (result *servicecatalog.CatalogEntry, err error) {
	result = &servicecatalog.CatalogEntry{}
	err = c.client.Put().
		Resource("catalogentries").
		Name(catalogEntry.Name).
		Body(catalogEntry).
		Do().
		Into(result)
	return
}

// Delete takes name of the catalogEntry and deletes it. Returns an error if one occurs.
func (c *catalogEntries) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Resource("catalogentries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *catalogEntries) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Resource("catalogentries").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the catalogEntry, and returns the corresponding catalogEntry object, and an error if there is any.
func (c *catalogEntries) Get(name string) (result *servicecatalog.CatalogEntry, err error) {
	result = &servicecatalog.CatalogEntry{}
	err = c.client.Get().
		Resource("catalogentries").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CatalogEntries that match those selectors.
func (c *catalogEntries) List(opts api.ListOptions) (result *servicecatalog.CatalogEntryList, err error) {
	result = &servicecatalog.CatalogEntryList{}
	err = c.client.Get().
		Resource("catalogentries").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested catalogEntries.
func (c *catalogEntries) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Resource("catalogentries").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
