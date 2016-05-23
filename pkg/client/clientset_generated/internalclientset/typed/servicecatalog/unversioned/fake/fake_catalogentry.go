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

// FakeCatalogEntries implements CatalogEntryInterface
type FakeCatalogEntries struct {
	Fake *FakeServicecatalog
}

var catalogentriesResource = unversioned.GroupVersionResource{Group: "servicecatalog", Version: "", Resource: "catalogentries"}

func (c *FakeCatalogEntries) Create(catalogEntry *servicecatalog.CatalogEntry) (result *servicecatalog.CatalogEntry, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(catalogentriesResource, catalogEntry), &servicecatalog.CatalogEntry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogEntry), err
}

func (c *FakeCatalogEntries) Update(catalogEntry *servicecatalog.CatalogEntry) (result *servicecatalog.CatalogEntry, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(catalogentriesResource, catalogEntry), &servicecatalog.CatalogEntry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogEntry), err
}

func (c *FakeCatalogEntries) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(catalogentriesResource, name), &servicecatalog.CatalogEntry{})
	return err
}

func (c *FakeCatalogEntries) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(catalogentriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &servicecatalog.CatalogEntryList{})
	return err
}

func (c *FakeCatalogEntries) Get(name string) (result *servicecatalog.CatalogEntry, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(catalogentriesResource, name), &servicecatalog.CatalogEntry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogEntry), err
}

func (c *FakeCatalogEntries) List(opts api.ListOptions) (result *servicecatalog.CatalogEntryList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(catalogentriesResource, opts), &servicecatalog.CatalogEntryList{})
	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.CatalogEntryList{}
	for _, item := range obj.(*servicecatalog.CatalogEntryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogEntries.
func (c *FakeCatalogEntries) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(catalogentriesResource, opts))
}
