package servicecatalog

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

const GroupName = "servicecatalog"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func AddToScheme(scheme *runtime.Scheme) {
	// Add the API to Scheme.
	addKnownTypes(scheme)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Catalog{},
		&CatalogList{},
		&CatalogEntry{},
		&CatalogEntryList{},
		&CatalogPosting{},
		&CatalogPostingList{},
		&CatalogClaim{},
		&CatalogClaimList{},
		&api.ListOptions{},
		&api.DeleteOptions{},
	)
}

func (obj *Catalog) GetObjectKind() unversioned.ObjectKind            { return &obj.TypeMeta }
func (obj *CatalogList) GetObjectKind() unversioned.ObjectKind        { return &obj.TypeMeta }
func (obj *CatalogEntry) GetObjectKind() unversioned.ObjectKind       { return &obj.TypeMeta }
func (obj *CatalogEntryList) GetObjectKind() unversioned.ObjectKind   { return &obj.TypeMeta }
func (obj *CatalogPosting) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *CatalogPostingList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
func (obj *CatalogClaim) GetObjectKind() unversioned.ObjectKind       { return &obj.TypeMeta }
func (obj *CatalogClaimList) GetObjectKind() unversioned.ObjectKind   { return &obj.TypeMeta }
