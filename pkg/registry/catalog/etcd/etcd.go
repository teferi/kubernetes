package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/catalog"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/runtime"
)

// REST implements a RESTStorage for catalogs against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against Catalogs.
func NewREST(opts generic.RESTOptions) *REST {
	prefix := "/catalogs"

	newListFunc := func() runtime.Object { return &catalogapi.CatalogList{} }
	storageInterface := opts.Decorator(
		opts.Storage, 100, &catalogapi.Catalog{}, prefix, catalog.Strategy, newListFunc)

	store := &registry.Store{
		NewFunc: func() runtime.Object { return &catalogapi.Catalog{} },

		// NewListFunc returns an object capable of storing results of an etcd list.
		NewListFunc: newListFunc,

		// Produces a path that etcd understands, to the root of the resource
		// by combining the namespace in the context with the given prefix
		KeyRootFunc: func(ctx api.Context) string {
			return registry.NamespaceKeyRootFunc(ctx, prefix)
		},

		// Produces a path that etcd understands, to the resource by combining
		// the namespace in the context with the given prefix
		KeyFunc: func(ctx api.Context, name string) (string, error) {
			return registry.NamespaceKeyFunc(ctx, prefix, name)
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
		DeleteStrategy: catalog.Strategy,

		Storage: storageInterface,
	}
	return &REST{store}
}
