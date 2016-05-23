package catalogposting

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

type catalogPostingStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

var Strategy = catalogPostingStrategy{api.Scheme, api.SimpleNameGenerator}

func (catalogPostingStrategy) NamespaceScoped() bool {
	return true
}

func (catalogPostingStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*catalogapi.CatalogPosting)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (catalogPostingStrategy) PrepareForUpdate(obj, old runtime.Object) {
}

func (catalogPostingStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// Canonicalize normalizes the object after validation.
func (catalogPostingStrategy) Canonicalize(obj runtime.Object) {
}

func (catalogPostingStrategy) AllowUnconditionalUpdate() bool {
	return true
}

func (catalogPostingStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (catalogPostingStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// CatalogPostingToSelectableFields returns a field set that represents the object for matching purposes.
func CatalogPostingToSelectableFields(catalog *catalogapi.CatalogPosting) fields.Set {
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
			catalogPosting, ok := obj.(*catalogapi.CatalogPosting)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a catalog posting.")
			}
			return labels.Set(catalog.ObjectMeta.Labels), CatalogPostingToSelectableFields(catalogPosting), nil
		},
	}
}
