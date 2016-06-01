package catalogentry

import (
	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/runtime"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/watch"
)

type Controller struct {
	kubeClient        clientset.Interface
	postingStore      StoreToCatalogPostingLister
	postingController *framework.Controller
	catalogEntryCache cache.Store
}

func NewController(kubeClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc, catalogEntryCache cache.Store) *Controller {
	c := &Controller{
		kubeClient:        kubeClient,
		catalogEntryCache: catalogEntryCache,
	}

	c.postingStore.Store, c.postingController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return c.kubeClient.Servicecatalog().CatalogPostings(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return c.kubeClient.Servicecatalog().CatalogPostings(api.NamespaceAll).Watch(options)
			},
		},
		&servicecatalog.CatalogPosting{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    c.postingAdded,
			UpdateFunc: c.postingUpdated,
			DeleteFunc: c.postingDeleted,
		},
	)

	return c
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	go c.postingController.Run(stopCh)
	<-stopCh
	glog.Infof("Shutting down catalog entry controller")
}

func (c *Controller) postingAdded(obj interface{}) {
	posting, ok := obj.(*servicecatalog.CatalogPosting)
	if !ok {
		glog.Errorf("expected type")
		return
	}
	entry := &servicecatalog.CatalogEntry{
		ObjectMeta: api.ObjectMeta{
			Name: posting.Name,
		},
		Catalog:         posting.Catalog,
		Description:     posting.Description,
		SourceNamespace: posting.Namespace,
	}
	c.catalogEntryCache.Add(entry)
	glog.Errorf("SETH saw added posting")
}

func (c *Controller) postingUpdated(oldObj, newObj interface{}) {
	/*old, ok := oldObj.(*servicecatalog.CatalogPosting)
	if !ok {
		glog.Errorf("expected type")
		return
	}
	posting, ok := newObj.(*servicecatalog.CatalogPosting)
	if !ok {
		glog.Errorf("expected type")
		return
	}
	catalog, ok := c.catalogEntryCache[old.Catalog]
	if !ok {
		glog.Error("expected catalog")
		return
	}
	if old.Catalog != posting.Catalog {
		glog.Error("catalog field can't be updated")
		return
	}
	delete(catalog, old.Name)
	catalog[posting.Name] = servicecatalog.CatalogEntry{
		ObjectMeta: api.ObjectMeta{
			Name: posting.Name,
		},
		Catalog:         posting.Catalog,
		Description:     posting.Description,
		SourceNamespace: posting.Namespace,
	}*/
	glog.Errorf("SETH saw updated posting")
}

func (c *Controller) postingDeleted(obj interface{}) {
	posting, ok := obj.(*servicecatalog.CatalogPosting)
	if !ok {
		glog.Errorf("expected type")
		return
	}
	err := c.catalogEntryCache.Delete(posting.Name)
	if err != nil {
		glog.Errorf("store failed delete on entry %s, %v\n", posting.Name, err)
	}
	glog.Errorf("SETH saw deleted posting")
}
