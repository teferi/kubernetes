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
	"crypto/tls"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/kubernetes/cmd/catalog-apiserver/app/options"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	"k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/genericapiserver"

	_ "k8s.io/kubernetes/pkg/apis/catalog/install"
)

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewAPIServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "catalog-apiserver",
		Long: `The Catalog API server validates and configures data for the 
		resource catalog api objects. The API Server services REST operations
		and provides the frontend to the cluster's shared state through which
		all other components interact.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(s *options.APIServer) error {
	genericapiserver.DefaultAndValidateRunOptions(s.ServerRunOptions)

	config := genericapiserver.NewConfig(s.ServerRunOptions)

	config.ProxyDialer = func(network, addr string) (net.Conn, error) { return nil, nil }
	config.ProxyTLSClientConfig = &tls.Config{}
	config.APIPrefix = "/api"
	config.APIGroupPrefix = "/apis"
	config.Serializer = api.Codecs

	m, err := genericapiserver.New(config)
	if err != nil {
		return err
	}

	catalogGroupMeta := registered.GroupOrDie(catalog.GroupName)
	catalogGroupMeta.GroupVersion = unversioned.GroupVersion{Group: "catalog", Version: "v1alpha1"}
	apiGroupInfo := &genericapiserver.APIGroupInfo{
		GroupMeta:                    *catalogGroupMeta,
		VersionedResourcesStorageMap: map[string]map[string]rest.Storage{},
		OptionsExternalVersion:       &registered.GroupOrDie(api.GroupName).GroupVersion,
		Scheme:                       api.Scheme,
		ParameterCodec:               api.ParameterCodec,
		NegotiatedSerializer:         api.Codecs,
	}

	m.InstallAPIGroup(apiGroupInfo)

	m.Run(s.ServerRunOptions)
	return nil
}
