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

// FakeCatalogs implements CatalogInterface
type FakeCatalogs struct {
	Fake *FakeServicecatalog
}

var catalogsResource = unversioned.GroupVersionResource{Group: "servicecatalog", Version: "", Resource: "catalogs"}

func (c *FakeCatalogs) Create(catalog *servicecatalog.Catalog) (result *servicecatalog.Catalog, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(catalogsResource, catalog), &servicecatalog.Catalog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.Catalog), err
}

func (c *FakeCatalogs) Update(catalog *servicecatalog.Catalog) (result *servicecatalog.Catalog, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(catalogsResource, catalog), &servicecatalog.Catalog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.Catalog), err
}

func (c *FakeCatalogs) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(catalogsResource, name), &servicecatalog.Catalog{})
	return err
}

func (c *FakeCatalogs) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(catalogsResource, listOptions)

	_, err := c.Fake.Invokes(action, &servicecatalog.CatalogList{})
	return err
}

func (c *FakeCatalogs) Get(name string) (result *servicecatalog.Catalog, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(catalogsResource, name), &servicecatalog.Catalog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.Catalog), err
}

func (c *FakeCatalogs) List(opts api.ListOptions) (result *servicecatalog.CatalogList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(catalogsResource, opts), &servicecatalog.CatalogList{})
	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.CatalogList{}
	for _, item := range obj.(*servicecatalog.CatalogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogs.
func (c *FakeCatalogs) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(catalogsResource, opts))
}
