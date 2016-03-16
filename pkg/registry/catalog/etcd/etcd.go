package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/catalog"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/runtime"
)

// REST implements a RESTStorage for catalogs against etcd
type REST struct {
	*etcdgeneric.Etcd
}

// NewREST returns a RESTStorage object that will work against Catalogs.
func NewREST(opts generic.RESTOptions) (*REST, *StatusREST) {
	prefix := "/catalogs"

	newListFunc := func() runtime.Object { return &catalogapi.CatalogList{} }

	store := &etcdgeneric.Etcd{
		NewFunc: func() runtime.Object { return &catalogapi.Catalog{} },

		// NewListFunc returns an object capable of storing results of an etcd list.
		NewListFunc: newListFunc,
		// Produces a path that etcd understands, to the root of the resource
		// by combining the namespace in the context with the given prefix
		KeyRootFunc: func(ctx api.Context) string {
			return etcdgeneric.NamespaceKeyRootFunc(ctx, prefix)
		},
		// Produces a path that etcd understands, to the resource by combining
		// the namespace in the context with the given prefix
		KeyFunc: func(ctx api.Context, name string) (string, error) {
			return etcdgeneric.NamespaceKeyFunc(ctx, prefix, name)
		},
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*catalogapi.Catalog).Name, nil
		},
		// Used to match objects based on labels/fields for list and watch
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return catalog.MatchCatalog(label, field)
		},
		QualifiedResource:       catalogapi.Resource("catalogs"),
		DeleteCollectionWorkers: opts.DeleteCollectionWorkers,

		CreateStrategy: catalog.Strategy,

		UpdateStrategy: catalog.Strategy,

		Storage: opts.Storage,
	}
	statusStore := *store
	statusStore.UpdateStrategy = catalog.StatusStrategy
	return &REST{store}, &StatusREST{store: &statusStore}
}

// StatusREST implements the REST endpoint for changing the status of a catalog
type StatusREST struct {
	store *etcdgeneric.Etcd
}

func (r *StatusREST) New() runtime.Object {
	return &catalogapi.Catalog{}
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx api.Context, obj runtime.Object) (runtime.Object, bool, error) {
	return r.store.Update(ctx, obj)
}
