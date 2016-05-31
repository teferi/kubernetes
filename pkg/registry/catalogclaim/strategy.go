package catalogclaim

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	catalogapi "k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

type catalogClaimStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

var Strategy = catalogClaimStrategy{api.Scheme, api.SimpleNameGenerator}

func (catalogClaimStrategy) NamespaceScoped() bool {
	return true
}

func (catalogClaimStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*catalogapi.CatalogClaim)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (catalogClaimStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogClaimStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// Canonicalize normalizes the object after validation.
func (catalogClaimStrategy) Canonicalize(obj runtime.Object) {
}

func (catalogClaimStrategy) AllowUnconditionalUpdate() bool {
	return true
}

func (catalogClaimStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (catalogClaimStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// catalogClaimToSelectableFields returns a field set that represents the object for matching purposes.
func CatalogClaimToSelectableFields(catalog *catalogapi.CatalogClaim) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(catalog.ObjectMeta, true)
	specificFieldsSet := fields.Set{}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}

// MatchCatalog is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchCatalogClaim(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			catalogClaim, ok := obj.(*catalogapi.CatalogClaim)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a catalog posting.")
			}
			return labels.Set(catalogClaim.Labels), CatalogClaimToSelectableFields(catalogClaim), nil
		},
	}
}
