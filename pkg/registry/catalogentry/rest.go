package catalogentry

import (
	"errors"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/golang/glog"
)

type catalogEntryREST struct {
	catalogEntryCache cache.Store
}

func NewREST(catalogEntryCache cache.Store) *catalogEntryREST {
	return &catalogEntryREST{catalogEntryCache: catalogEntryCache}
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
	entries := r.catalogEntryCache.List()
	items := make([]servicecatalog.CatalogEntry, len(entries))
	for i, entry := range entries {
		items[i] = *(entry.(*servicecatalog.CatalogEntry))
	}
	return &servicecatalog.CatalogEntryList{Items: items}, nil
}

func (r *catalogEntryREST) Get(ctx api.Context, name string) (runtime.Object, error) {
	glog.Errorf("SETH getting the things")
	entry, ok, err := r.catalogEntryCache.GetByKey(name)
	if err != nil {
		glog.Errorf("SETH error getting entry %s: %v\n", name, err)
		return nil, err
	}
	if !ok {
		glog.Errorf("SETH entry  not found\n")
		return nil, errors.New("entry not found")
	}
	return entry.(*servicecatalog.CatalogEntry), nil
}
