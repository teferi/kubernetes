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

package node

import (
	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/runtime"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/watch"
)

type CatalogController struct {
	kubeClient                  clientset.Interface
	catalogStore                cache.StoreToCatalogLister
	catalogController           *framework.Controller
	catalogEntryStore           cache.StoreToCatalogEntryLister
	catalogEntryController      *framework.Controller
	catalogEntryClaimStore      cache.StoreToCatalogEntryClaimLister
	catalogEntryClaimController *framework.Controller
}

// NewCalatogController returns a new catalog controller to sync catalog entries and claims.
func NewCatalogController(kubeClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc) *CatalogController {
	cc := &CatalogController{
		kubeClient: kubeClient,
	}

	cc.catalogStore.Store, cc.catalogController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return cc.kubeClient.Catalog().Catalogs(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return cc.kubeClient.Catalog().Catalogs(api.NamespaceAll).Watch(options)
			},
		},
		&catalog.Catalog{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    cc.catalogAdd,
			UpdateFunc: func(_, obj interface{}) { cc.catalogUpdate(obj) },
			DeleteFunc: cc.catalogDelete,
		},
	)

	cc.catalogEntryStore.Store, cc.catalogEntryController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return cc.kubeClient.Catalog().CatalogEntries(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return cc.kubeClient.Catalog().CatalogEntries(api.NamespaceAll).Watch(options)
			},
		},
		&catalog.CatalogEntry{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    cc.catalogEntryAdd,
			UpdateFunc: func(_, obj interface{}) { cc.catalogEntryUpdate(obj) },
			DeleteFunc: cc.catalogEntryDelete,
		},
	)

	cc.catalogEntryClaimStore.Store, cc.catalogEntryClaimController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return cc.kubeClient.Catalog().CatalogEntryClaims(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return cc.kubeClient.Catalog().CatalogEntryClaims(api.NamespaceAll).Watch(options)
			},
		},
		&catalog.CatalogEntryClaim{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    cc.catalogEntryClaimAdd,
			UpdateFunc: func(_, obj interface{}) { cc.catalogEntryClaimUpdate(obj) },
			DeleteFunc: cc.catalogEntryClaimDelete,
		},
	)

	return cc
}

// Run begins watching catalog resources.
func (cc *CatalogController) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	go cc.catalogController.Run(stopCh)
	go cc.catalogEntryController.Run(stopCh)
	go cc.catalogEntryClaimController.Run(stopCh)
	<-stopCh
	glog.Infof("Shutting down catalog controller")
}

func (cc *CatalogController) catalogAdd(obj interface{}) {
	c := obj.(*catalog.Catalog)
	glog.V(4).Infof("Catalog %s was added", c.Name)
}

func (cc *CatalogController) catalogUpdate(obj interface{}) {
	c := obj.(*catalog.Catalog)
	glog.V(4).Infof("Catalog %s was updated", c.Name)
}

func (cc *CatalogController) catalogDelete(obj interface{}) {
	c := obj.(*catalog.Catalog)
	glog.V(4).Infof("Catalog %s was deleted", c.Name)
}

func (cc *CatalogController) catalogEntryAdd(obj interface{}) {
	c := obj.(*catalog.CatalogEntry)
	glog.V(4).Infof("CatalogEntry %s was added", c.Name)
}

func (cc *CatalogController) catalogEntryUpdate(obj interface{}) {
	c := obj.(*catalog.CatalogEntry)
	glog.V(4).Infof("CatalogEntry %s was updated", c.Name)
}

func (cc *CatalogController) catalogEntryDelete(obj interface{}) {
	c := obj.(*catalog.CatalogEntry)
	glog.V(4).Infof("CatalogEntry %s was deleted", c.Name)
}

func (cc *CatalogController) catalogEntryClaimAdd(obj interface{}) {
	c := obj.(*catalog.CatalogEntryClaim)
	glog.V(4).Infof("CatalogEntryClaim %s was added", c.Name)
}

func (cc *CatalogController) catalogEntryClaimUpdate(obj interface{}) {
	c := obj.(*catalog.CatalogEntryClaim)
	glog.V(4).Infof("CatalogEntryClaim %s was updated", c.Name)
}

func (cc *CatalogController) catalogEntryClaimDelete(obj interface{}) {
	c := obj.(*catalog.CatalogEntryClaim)
	glog.V(4).Infof("CatalogEntryClaim %s was deleted", c.Name)
}
