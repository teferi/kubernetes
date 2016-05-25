package catalogentry

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/runtime"
)

type catalogEntryREST struct {
	catalogEntryCache cache.Store
}

func NewREST(catalogEntryCache cache.Store) *catalogEntryREST {
	return &catalogEntryREST{}
}

func (r *catalogEntryREST) New() runtime.Object {
	return &servicecatalog.CatalogEntry{}
}

func (r *catalogEntryREST) NewList() runtime.Object {
	return &servicecatalog.CatalogEntryList{}
}

func (r *catalogEntryREST) List(ctx api.Context, options *api.ListOptions) (runtime.Object, error) {
	// kubectl get --namespace finance catalogentries
	// /apis/servicecatalog/v1alpha1/namespaces/finance/catalogentries
	// kubectl get --namespace finance catalogentries/oracle
	// kubectl get catalogentries/finance/oracle <-- does not work
	//r.catalogEntryCache.Get(<namemspace>)
	return &servicecatalog.CatalogEntryList{}, nil

}
