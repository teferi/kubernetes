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

// This file is generated by client-gen with arguments: --input=[api/,extensions/,catalog/]

package unversioned

import (
	api "k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/catalog"
	watch "k8s.io/kubernetes/pkg/watch"
)

// CatalogsGetter has a method to return a CatalogInterface.
// A group's client should implement this interface.
type CatalogsGetter interface {
	Catalogs(namespace string) CatalogInterface
}

// CatalogInterface has methods to work with Catalog resources.
type CatalogInterface interface {
	Create(*catalogapi.Catalog) (*catalogapi.Catalog, error)
	Update(*catalogapi.Catalog) (*catalogapi.Catalog, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*catalogapi.Catalog, error)
	List(opts api.ListOptions) (*catalogapi.CatalogList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	CatalogExpansion
}

// catalogs implements CatalogInterface
type catalogs struct {
	client *CatalogClient
	ns     string
}

// newCatalogs returns a Catalogs
func newCatalogs(c *CatalogClient, namespace string) *catalogs {
	return &catalogs{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a catalog and creates it.  Returns the server's representation of the catalog, and an error, if there is any.
func (c *catalogs) Create(catalog *catalogapi.Catalog) (result *catalogapi.Catalog, err error) {
	result = &catalogapi.Catalog{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("catalogs").
		Body(catalog).
		Do().
		Into(result)
	return
}

// Update takes the representation of a catalog and updates it. Returns the server's representation of the catalog, and an error, if there is any.
func (c *catalogs) Update(catalog *catalogapi.Catalog) (result *catalogapi.Catalog, err error) {
	result = &catalogapi.Catalog{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("catalogs").
		Name(catalog.Name).
		Body(catalog).
		Do().
		Into(result)
	return
}

// Delete takes name of the catalog and deletes it. Returns an error if one occurs.
func (c *catalogs) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *catalogs) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogs").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the catalog, and returns the corresponding catalog object, and an error if there is any.
func (c *catalogs) Get(name string) (result *catalogapi.Catalog, err error) {
	result = &catalogapi.Catalog{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogs").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Catalogs that match those selectors.
func (c *catalogs) List(opts api.ListOptions) (result *catalogapi.CatalogList, err error) {
	result = &catalogapi.CatalogList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogs").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested catalogs.
func (c *catalogs) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("catalogs").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
