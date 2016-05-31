package catalogentry

import (
	"errors"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/golang/glog"
)

type catalogEntryREST struct {
	catalogEntryCache map[string]map[string]servicecatalog.CatalogEntry
}

func NewREST(catalogEntryCache map[string]map[string]servicecatalog.CatalogEntry) *catalogEntryREST {
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
	list := &servicecatalog.CatalogEntryList{}
	namespace, ok := api.NamespaceFrom(ctx)
	if !ok {
		glog.Errorf("SETH namespace not found\n")
		return nil, errors.New("namespace not found")
	}
	entries, ok := r.catalogEntryCache[namespace]
	if !ok {
		return list, nil
	}
	items := make([]servicecatalog.CatalogEntry, len(entries))
	i := 0
	for _, value := range entries {
		items[i] = value
		i++
	}
	list.Items = items
	return list, nil
}
