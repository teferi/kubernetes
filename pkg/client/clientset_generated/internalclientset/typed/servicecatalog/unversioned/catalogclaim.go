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

// CatalogClaimsGetter has a method to return a CatalogClaimInterface.
// A group's client should implement this interface.
type CatalogClaimsGetter interface {
	CatalogClaims(namespace string) CatalogClaimInterface
}

// CatalogClaimInterface has methods to work with CatalogClaim resources.
type CatalogClaimInterface interface {
	Create(*servicecatalog.CatalogClaim) (*servicecatalog.CatalogClaim, error)
	Update(*servicecatalog.CatalogClaim) (*servicecatalog.CatalogClaim, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*servicecatalog.CatalogClaim, error)
	List(opts api.ListOptions) (*servicecatalog.CatalogClaimList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	CatalogClaimExpansion
}

// catalogClaims implements CatalogClaimInterface
type catalogClaims struct {
	client *ServicecatalogClient
	ns     string
}

// newCatalogClaims returns a CatalogClaims
func newCatalogClaims(c *ServicecatalogClient, namespace string) *catalogClaims {
	return &catalogClaims{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a catalogClaim and creates it.  Returns the server's representation of the catalogClaim, and an error, if there is any.
func (c *catalogClaims) Create(catalogClaim *servicecatalog.CatalogClaim) (result *servicecatalog.CatalogClaim, err error) {
	result = &servicecatalog.CatalogClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("catalogclaims").
		Body(catalogClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a catalogClaim and updates it. Returns the server's representation of the catalogClaim, and an error, if there is any.
func (c *catalogClaims) Update(catalogClaim *servicecatalog.CatalogClaim) (result *servicecatalog.CatalogClaim, err error) {
	result = &servicecatalog.CatalogClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("catalogclaims").
		Name(catalogClaim.Name).
		Body(catalogClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the catalogClaim and deletes it. Returns an error if one occurs.
func (c *catalogClaims) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *catalogClaims) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("catalogclaims").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the catalogClaim, and returns the corresponding catalogClaim object, and an error if there is any.
func (c *catalogClaims) Get(name string) (result *servicecatalog.CatalogClaim, err error) {
	result = &servicecatalog.CatalogClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogclaims").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CatalogClaims that match those selectors.
func (c *catalogClaims) List(opts api.ListOptions) (result *servicecatalog.CatalogClaimList, err error) {
	result = &servicecatalog.CatalogClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("catalogclaims").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested catalogClaims.
func (c *catalogClaims) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("catalogclaims").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
