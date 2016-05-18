/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

// Package app does all of the work necessary to create a Kubernetes
// APIServer by binding together the API, master and APIServer infrastructure.
// It can be configured and called directly or via the hyperkube framework.
package app

import (
	"k8s.io/kubernetes/cmd/catalog-apiserver/app/options"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	"k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/genericapiserver"
	catalogetcd "k8s.io/kubernetes/pkg/registry/catalog/etcd"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/storage/storagebackend"

	_ "k8s.io/kubernetes/pkg/apis/catalog/install"
)

// Run runs the specified APIServer.  This should never exit.
func Run(s *options.APIServer) error {
	genericapiserver.DefaultAndValidateRunOptions(s.ServerRunOptions)

	c := genericapiserver.NewConfig(s.ServerRunOptions)

	/*c.ProxyDialer = func(network, addr string) (net.Conn, error) { return nil, nil }
	c.ProxyTLSClientConfig = &tls.Config{}
	c.APIPrefix = "/api"
	c.APIGroupPrefix = "/apis"*/
	c.Serializer = api.Codecs

	m, err := genericapiserver.New(c)
	if err != nil {
		return err
	}

	// Create Storage
	config := storagebackend.Config{
		Prefix:     genericapiserver.DefaultEtcdPathPrefix,
		ServerList: s.ServerRunOptions.StorageConfig.ServerList,
	}
	storageFactory := genericapiserver.NewDefaultStorageFactory(config, "application/json", api.Codecs, genericapiserver.NewDefaultResourceEncodingConfig(), genericapiserver.NewResourceConfig())
	storage, err := storageFactory.New(unversioned.GroupResource{Group: catalog.GroupName, Resource: "catalog"})
	if err != nil {
		return err
	}

	restOptions := generic.RESTOptions{
		Storage:                 storage,
		Decorator:               m.StorageDecorator(),
		DeleteCollectionWorkers: s.DeleteCollectionWorkers,
	}

	catalogStorage := catalogetcd.NewREST(restOptions)
	restStorageMap := map[string]rest.Storage{
		"catalogs": catalogStorage,
	}

	// Create API Group
	catalogGroupMeta := registered.GroupOrDie(catalog.GroupName)
	catalogGroupMeta.GroupVersion = unversioned.GroupVersion{Group: "catalog", Version: "v1alpha1"}
	apiGroupInfo := &genericapiserver.APIGroupInfo{
		GroupMeta: *catalogGroupMeta,
		VersionedResourcesStorageMap: map[string]map[string]rest.Storage{
			"v1alpha1": restStorageMap,
		},
		Scheme:               api.Scheme,
		NegotiatedSerializer: api.Codecs,
	}

	err = m.InstallAPIGroup(apiGroupInfo)
	if err != nil {
		return err
	}

	m.Run(s.ServerRunOptions)
	return nil
}
