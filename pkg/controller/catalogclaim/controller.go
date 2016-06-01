package catalogentry

import (
	"fmt"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/typed/dynamic"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/runtime"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/watch"
)

type Controller struct {
	kubeClient      clientset.Interface
	catalogClient   clientset.Interface
	claimStore      StoreToCatalogClaimLister
	claimController *framework.Controller
	clientPool      dynamic.ClientPool
}

func NewController(kubeClient clientset.Interface, catalogClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc, clientPool dynamic.ClientPool) *Controller {
	c := &Controller{
		kubeClient:    kubeClient,
		catalogClient: catalogClient,
	}

	c.claimStore.Store, c.claimController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return c.catalogClient.Servicecatalog().CatalogClaims(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return c.catalogClient.Servicecatalog().CatalogClaims(api.NamespaceAll).Watch(options)
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
	entry, err := c.catalogClient.Servicecatalog().CatalogEntries().Get(claim.Spec.Entry)
	if err != nil {
		glog.Errorf("error getting entry %s: %v", claim.Spec.Entry, err)
		return
	}
	posting, err := c.catalogClient.Servicecatalog().CatalogPostings(entry.SourceNamespace).Get(entry.Name)
	if err != nil {
		glog.Errorf("error getting posting for entry %s: %v", entry.Name, err)
		return
	}

	if err != nil {
		glog.Errorf("dynamic client creation failed %v", err)
		return
	}
	for _, localResource := range posting.LocalResources.Items {
		/* please no
		gv, err := unversioned.ParseGroupVersion(localResource.APIVersion)
		if err != nil {
			glog.Errorf("parse gv failed %v", err)
			return
		}
		dynamicClient, err := c.clientPool.ClientForGroupVersion(gv)
		if err != nil {
			glog.Errorf("create dynamicClient failed %v", err)
			return
		}
		apiResource := unversioned.APIResource{Name: localResource.Kind, Namespaced: true}
		obj := dynamicClient.Resource(&apiResource, posting.Namespace).Get(localResource.Name)
		data, err := json.Marshal(obj)
		obj, err = runtime.Decode(codec, data)
		*/
		switch {
		case localResource.APIVersion == "v1" && localResource.Kind == "Secret":
			secret, err := c.kubeClient.Core().Secrets(posting.Namespace).Get(localResource.Name)
			if err != nil {
				glog.Errorf("secret not found: %v", err)
				return
			}
			newSecret := api.Secret{
				ObjectMeta: api.ObjectMeta{
					Name: fmt.Sprintf("%s-%s", claim.Name, localResource.Name),
				},
				Data: secret.Data,
				Type: secret.Type,
			}
			_, err = c.kubeClient.Core().Secrets(claim.Namespace).Create(&newSecret)
			if err != nil {
				glog.Errorf("secret creation failed: %v", err)
				return
			}
		}

		// set status and created resources
	}
	glog.Errorf("SETH saw added claim")
}

func (c *Controller) claimUpdated(oldObj, newObj interface{}) {

	glog.Errorf("SETH saw updated claim")
}

func (c *Controller) claimDeleted(obj interface{}) {

	glog.Errorf("SETH saw deleted claim")
}
