package catalogentry

import (
	"encoding/json"
	"fmt"

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
	kubeClient       clientset.Interface
	catalogClient    clientset.Interface
	claimStore       StoreToCatalogClaimLister
	claimController  *framework.Controller
	secretStore      cache.Store
	secretController *framework.Controller
	//clientPool      dynamic.ClientPool
}

func NewController(kubeClient clientset.Interface, catalogClient clientset.Interface, resyncPeriod controller.ResyncPeriodFunc) *Controller {
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

	c.secretStore, c.secretController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return c.kubeClient.Core().Secrets(api.NamespaceAll).List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return c.kubeClient.Core().Secrets(api.NamespaceAll).Watch(options)
			},
		},
		&api.Secret{},
		resyncPeriod(),
		framework.ResourceEventHandlerFuncs{
			AddFunc:    c.secretAdded,
			UpdateFunc: c.secretUpdated,
			DeleteFunc: c.secretDeleted,
		},
	)

	return c
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	go c.claimController.Run(stopCh)
	go c.secretController.Run(stopCh)
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
			if secret.Annotations == nil {
				secret.Annotations = make(map[string]string)
			}
			var claimers []api.ObjectReference
			claimerstr, ok := secret.Annotations["claimers"]
			if ok {
				err := json.Unmarshal([]byte(claimerstr), claimers)
				if err != nil {
					glog.Errorf("unmarshal annotation failed: %v", err)
					return
				}
			} else {
				claimers = []api.ObjectReference{}
			}
			newclaimer := api.ObjectReference{Namespace: claim.Namespace, Name: newSecret.Name}
			claimers = append(claimers, newclaimer)
			claimersjson, err := json.Marshal(claimers)
			if err != nil {
				glog.Errorf("marshal annotation failed: %v", err)
				return
			}
			secret.Annotations["claimers"] = string(claimersjson)
			_, err = c.kubeClient.Core().Secrets(posting.Namespace).Update(secret)
			if err != nil {
				glog.Errorf("source secret update failed: %v", err)
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

func (c *Controller) secretAdded(obj interface{}) {
	glog.Errorf("SETH saw added secret")
}

func (c *Controller) secretUpdated(oldObj, newObj interface{}) {
	secret, ok := newObj.(*api.Secret)
	if !ok {
		glog.Errorf("expected type secret")
		return
	}
	claimers, ok := secret.Annotations["claimers"]
	if !ok {
		glog.Errorf("SETH DEBUG no claimers annotation found")
		return
	}
	var claimerSecrets []api.ObjectReference
	if err := json.Unmarshal([]byte(claimers), &claimerSecrets); err != nil {
		glog.Errorf("failed to unmarshal annotation %s", claimers)
		return
	}

	for _, cs := range claimerSecrets {
		dss, err := c.kubeClient.Core().Secrets(cs.Namespace).Get(cs.Name)
		if err != nil {
			glog.Errorf("could not find claimer secret %s/%s", cs.Namespace, cs.Name)
			continue
		}
		dss.Data = secret.Data
		_, err = c.kubeClient.Core().Secrets(cs.Namespace).Update(dss)
		if err != nil {
			glog.Errorf("could not update claimer secret %s/%s", cs.Namespace, cs.Name)
			continue
		}
	}

	glog.Errorf("SETH saw updated secret")
}

func (c *Controller) secretDeleted(obj interface{}) {
	glog.Errorf("SETH saw deleted secret")
}
