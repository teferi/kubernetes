package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/catalogposting"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/runtime"
)

// REST implements a RESTStorage for catalog postings against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against CatalogPostings.
func NewREST(opts generic.RESTOptions) *REST {
	prefix := "/catalogpostings"

	newListFunc := func() runtime.Object { return &catalogapi.CatalogPostingList{} }
	storageInterface := opts.Decorator(
		opts.Storage, 100, &catalogapi.CatalogPosting{}, prefix, catalogposting.Strategy, newListFunc)

	store := &registry.Store{
		NewFunc: func() runtime.Object { return &catalogapi.CatalogPosting{} },

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
			return registry.NamespaceKeyFunc(ctx, prefix, name)
		},

		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*catalogapi.CatalogPosting).Name, nil
		},

		// Used to match objects based on labels/fields for list and watch
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return catalogposting.MatchCatalogPosting(label, field)
		},

		QualifiedResource:       catalogapi.Resource("catalogpostings"),
		DeleteCollectionWorkers: opts.DeleteCollectionWorkers,

		CreateStrategy: catalogposting.Strategy,
		UpdateStrategy: catalogposting.Strategy,
		DeleteStrategy: catalogposting.Strategy,

		Storage: storageInterface,
	}
	return &REST{store}
}
