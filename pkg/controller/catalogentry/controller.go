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
	catalogEntryCache map[string][]servicecatalog.CatalogEntry
}

func NewController(kubeClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc, catalogEntryCache map[string][]servicecatalog.CatalogEntry) *Controller {
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
	entry := servicecatalog.CatalogEntry{
		ObjectMeta: api.ObjectMeta{
			Name: posting.Name,
		},
		Catalog:     posting.Catalog,
		Description: posting.Description,
	}
	if _, ok := c.catalogEntryCache[posting.Catalog]; !ok {
		c.catalogEntryCache[posting.Catalog] = []servicecatalog.CatalogEntry{}
	}
	c.catalogEntryCache[posting.Catalog] = append(c.catalogEntryCache[posting.Catalog], entry)
	glog.Errorf("SETH saw added posting")
}

func (c *Controller) postingUpdated(oldObj, newObj interface{}) {
	glog.Errorf("SETH saw updated posting")
}

func (c *Controller) postingDeleted(obj interface{}) {
	glog.Errorf("SETH saw deleted posting")
}
