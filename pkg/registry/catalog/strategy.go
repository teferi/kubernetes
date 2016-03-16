package catalog

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/catalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

type catalogStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

var Strategy = catalogStrategy{api.Scheme, api.SimpleNameGenerator}

func (catalogStrategy) NamespaceScoped() bool {
	return true
}

func (catalogStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*catalogapi.Catalog)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (catalogStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// Canonicalize normalizes the object after validation.
func (catalogStrategy) Canonicalize(obj runtime.Object) {
}

func (catalogStrategy) AllowUnconditionalUpdate() bool {
	return true
}

func (catalogStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (catalogStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// CatalogToSelectableFields returns a field set that represents the object for matching purposes.
func CatalogToSelectableFields(catalog *catalogapi.Catalog) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(catalog.ObjectMeta, true)
	specificFieldsSet := fields.Set{}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}

// MatchCatalog is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchCatalog(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			catalog, ok := obj.(*catalogapi.Catalog)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a catalog.")
			}
			return labels.Set(catalog.ObjectMeta.Labels), CatalogToSelectableFields(catalog), nil
		},
	}
}

type catalogStatusStrategy struct {
	catalogStrategy
}

var StatusStrategy = catalogStatusStrategy{Strategy}

func (catalogStatusStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogStatusStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
