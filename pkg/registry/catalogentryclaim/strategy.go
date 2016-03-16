package catalogentryclaim

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

type catalogEntryClaimStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

var Strategy = catalogEntryClaimStrategy{api.Scheme, api.SimpleNameGenerator}

func (catalogEntryClaimStrategy) NamespaceScoped() bool {
	return true
}

func (catalogEntryClaimStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*catalogapi.CatalogEntryClaim)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (catalogEntryClaimStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogEntryClaimStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// Canonicalize normalizes the object after validation.
func (catalogEntryClaimStrategy) Canonicalize(obj runtime.Object) {
}

func (catalogEntryClaimStrategy) AllowUnconditionalUpdate() bool {
	return true
}

func (catalogEntryClaimStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (catalogEntryClaimStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// CatalogEntryClaimToSelectableFields returns a field set that represents the object for matching purposes.
func CatalogEntryClaimToSelectableFields(catalog *catalogapi.CatalogEntryClaim) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(catalog.ObjectMeta, true)
	specificFieldsSet := fields.Set{}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}

// MatchCatalogEntryClaim is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchCatalogEntryClaim(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			catalog, ok := obj.(*catalogapi.CatalogEntryClaim)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a catalog entry claim.")
			}
			return labels.Set(catalog.ObjectMeta.Labels), CatalogEntryClaimToSelectableFields(catalog), nil
		},
	}
}

type catalogEntryClaimStatusStrategy struct {
	catalogEntryClaimStrategy
}

var StatusStrategy = catalogEntryClaimStatusStrategy{Strategy}

func (catalogEntryClaimStatusStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogEntryClaimStatusStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
