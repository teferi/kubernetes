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

package fake

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	servicecatalog "k8s.io/kubernetes/pkg/apis/servicecatalog"
	core "k8s.io/kubernetes/pkg/client/testing/core"
	labels "k8s.io/kubernetes/pkg/labels"
	watch "k8s.io/kubernetes/pkg/watch"
)

// FakeCatalogClaims implements CatalogClaimInterface
type FakeCatalogClaims struct {
	Fake *FakeServicecatalog
	ns   string
}

var catalogclaimsResource = unversioned.GroupVersionResource{Group: "servicecatalog", Version: "", Resource: "catalogclaims"}

func (c *FakeCatalogClaims) Create(catalogClaim *servicecatalog.CatalogClaim) (result *servicecatalog.CatalogClaim, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction(catalogclaimsResource, c.ns, catalogClaim), &servicecatalog.CatalogClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogClaim), err
}

func (c *FakeCatalogClaims) Update(catalogClaim *servicecatalog.CatalogClaim) (result *servicecatalog.CatalogClaim, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction(catalogclaimsResource, c.ns, catalogClaim), &servicecatalog.CatalogClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogClaim), err
}

func (c *FakeCatalogClaims) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction(catalogclaimsResource, c.ns, name), &servicecatalog.CatalogClaim{})

	return err
}

func (c *FakeCatalogClaims) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewDeleteCollectionAction(catalogclaimsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &servicecatalog.CatalogClaimList{})
	return err
}

func (c *FakeCatalogClaims) Get(name string) (result *servicecatalog.CatalogClaim, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction(catalogclaimsResource, c.ns, name), &servicecatalog.CatalogClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogClaim), err
}

func (c *FakeCatalogClaims) List(opts api.ListOptions) (result *servicecatalog.CatalogClaimList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction(catalogclaimsResource, c.ns, opts), &servicecatalog.CatalogClaimList{})

	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.CatalogClaimList{}
	for _, item := range obj.(*servicecatalog.CatalogClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogClaims.
func (c *FakeCatalogClaims) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction(catalogclaimsResource, c.ns, opts))

}
