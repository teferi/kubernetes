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
	unversioned "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/servicecatalog/unversioned"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeServicecatalog struct {
	*core.Fake
}

func (c *FakeServicecatalog) Catalogs() unversioned.CatalogInterface {
	return &FakeCatalogs{c}
}

func (c *FakeServicecatalog) CatalogClaims(namespace string) unversioned.CatalogClaimInterface {
	return &FakeCatalogClaims{c, namespace}
}

func (c *FakeServicecatalog) CatalogEntries() unversioned.CatalogEntryInterface {
	return &FakeCatalogEntries{c}
}

func (c *FakeServicecatalog) CatalogPostings(namespace string) unversioned.CatalogPostingInterface {
	return &FakeCatalogPostings{c, namespace}
}

// GetRESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicecatalog) GetRESTClient() *restclient.RESTClient {
	return nil
}
