package catalogentry

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

type catalogEntryStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

var Strategy = catalogEntryStrategy{api.Scheme, api.SimpleNameGenerator}

func (catalogEntryStrategy) NamespaceScoped() bool {
	return true
}

func (catalogEntryStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*catalogapi.CatalogEntry)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (catalogEntryStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogEntryStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// Canonicalize normalizes the object after validation.
func (catalogEntryStrategy) Canonicalize(obj runtime.Object) {
}

func (catalogEntryStrategy) AllowUnconditionalUpdate() bool {
	return true
}

func (catalogEntryStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (catalogEntryStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// CatalogEntryToSelectableFields returns a field set that represents the object for matching purposes.
func CatalogEntryToSelectableFields(catalog *catalogapi.CatalogEntry) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(catalog.ObjectMeta, true)
	specificFieldsSet := fields.Set{}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}

// MatchCatalogEntry is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchCatalogEntry(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			catalog, ok := obj.(*catalogapi.CatalogEntry)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a catalog entry.")
			}
			return labels.Set(catalog.ObjectMeta.Labels), CatalogEntryToSelectableFields(catalog), nil
		},
	}
}

type catalogEntryStatusStrategy struct {
	catalogEntryStrategy
}

var StatusStrategy = catalogEntryStatusStrategy{Strategy}

func (catalogEntryStatusStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogEntryStatusStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
