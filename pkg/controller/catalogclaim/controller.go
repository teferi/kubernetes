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
	kubeClient      clientset.Interface
	claimStore      StoreToCatalogClaimLister
	claimController *framework.Controller
}

func NewController(kubeClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc) *Controller {
	c := &Controller{
		kubeClient: kubeClient,
	}

	c.claimStore.Store, c.claimController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return c.kubeClient.Servicecatalog().CatalogClaims(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return c.kubeClient.Servicecatalog().CatalogClaims(api.NamespaceAll).Watch(options)
			},
		},
		&servicecatalog.CatalogClaim{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    c.claimAdded,
			UpdateFunc: c.claimUpdated,
			DeleteFunc: c.claimDeleted,
		},
	)

	return c
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	go c.claimController.Run(stopCh)
	<-stopCh
	glog.Infof("Shutting down catalog claim controller")
}

func (c *Controller) claimAdded(obj interface{}) {
	claim, ok := obj.(*servicecatalog.CatalogClaim)
	if !ok {
		glog.Errorf("expected type")
		return
	}
	entry, err := c.kubeClient.Servicecatalog().CatalogEntries().Get(claim.Spec.Entry)
	if err != nil {
		glog.Errorf("error getting entry %s: %v", claim.Spec.Entry, err)
		return
	}
	posting, err := c.kubeClient.Servicecatalog().CatalogPostings(entry.SourceNamespace).Get(entry.Name)
	if err != nil {
		glog.Errorf("error getting posting for entry %s: %v", entry.Name, err)
		return
	}
	_ = posting
	glog.Errorf("SETH saw added claim")
}

func (c *Controller) claimUpdated(oldObj, newObj interface{}) {

	glog.Errorf("SETH saw updated claim")
}

func (c *Controller) claimDeleted(obj interface{}) {

	glog.Errorf("SETH saw deleted claim")
}
