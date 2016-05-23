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

// FakeCatalogPostings implements CatalogPostingInterface
type FakeCatalogPostings struct {
	Fake *FakeServicecatalog
	ns   string
}

var catalogpostingsResource = unversioned.GroupVersionResource{Group: "servicecatalog", Version: "", Resource: "catalogpostings"}

func (c *FakeCatalogPostings) Create(catalogPosting *servicecatalog.CatalogPosting) (result *servicecatalog.CatalogPosting, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction(catalogpostingsResource, c.ns, catalogPosting), &servicecatalog.CatalogPosting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogPosting), err
}

func (c *FakeCatalogPostings) Update(catalogPosting *servicecatalog.CatalogPosting) (result *servicecatalog.CatalogPosting, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction(catalogpostingsResource, c.ns, catalogPosting), &servicecatalog.CatalogPosting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogPosting), err
}

func (c *FakeCatalogPostings) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction(catalogpostingsResource, c.ns, name), &servicecatalog.CatalogPosting{})

	return err
}

func (c *FakeCatalogPostings) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewDeleteCollectionAction(catalogpostingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &servicecatalog.CatalogPostingList{})
	return err
}

func (c *FakeCatalogPostings) Get(name string) (result *servicecatalog.CatalogPosting, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction(catalogpostingsResource, c.ns, name), &servicecatalog.CatalogPosting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.CatalogPosting), err
}

func (c *FakeCatalogPostings) List(opts api.ListOptions) (result *servicecatalog.CatalogPostingList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction(catalogpostingsResource, c.ns, opts), &servicecatalog.CatalogPostingList{})

	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.CatalogPostingList{}
	for _, item := range obj.(*servicecatalog.CatalogPostingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogPostings.
func (c *FakeCatalogPostings) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction(catalogpostingsResource, c.ns, opts))

}
