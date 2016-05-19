package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/catalogentry"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/runtime"
)

// REST implements a RESTStorage for catalogentries against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against Catalogs.
func NewREST(opts generic.RESTOptions) *REST {
	prefix := "/catalogentries"

	newListFunc := func() runtime.Object { return &catalogapi.CatalogEntryList{} }
	storageInterface := opts.Decorator(
		opts.Storage, 100, &catalogapi.CatalogEntry{}, prefix, catalogentry.Strategy, newListFunc)

	store := &registry.Store{
		NewFunc: func() runtime.Object { return &catalogapi.CatalogEntry{} },

		// NewListFunc returns an object capable of storing results of an etcd list.
		NewListFunc: newListFunc,

		// Produces a path that etcd understands, to the root of the resource
		// by combining the namespace in the context with the given prefix
		KeyRootFunc: func(ctx api.Context) string {
			return prefix
		},

		// Produces a path that etcd understands, to the resource by combining
		// the namespace in the context with the given prefix
		KeyFunc: func(ctx api.Context, name string) (string, error) {
			return registry.NoNamespaceKeyFunc(ctx, prefix, name)
		},

		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*catalogapi.CatalogEntry).Name, nil
		},

		// Used to match objects based on labels/fields for list and watch
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return catalogentry.MatchCatalog(label, field)
		},

		QualifiedResource:       catalogapi.Resource("catalogentries"),
		DeleteCollectionWorkers: opts.DeleteCollectionWorkers,

		CreateStrategy: catalogentry.Strategy,
		UpdateStrategy: catalogentry.Strategy,
		DeleteStrategy: catalogentry.Strategy,

		Storage: storageInterface,
	}
	return &REST{store}
}
