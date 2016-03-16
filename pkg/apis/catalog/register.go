package catalog

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

const GroupName = "catalog"

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
		&CatalogEntryClaim{},
		&CatalogEntryClaimList{},
		&CatalogEntryProvider{},
		&CatalogEntryProviderList{},
		&api.ListOptions{},
	)
}

func (obj *Catalog) GetObjectKind() unversioned.ObjectKind                  { return &obj.TypeMeta }
func (obj *CatalogList) GetObjectKind() unversioned.ObjectKind              { return &obj.TypeMeta }
func (obj *CatalogEntry) GetObjectKind() unversioned.ObjectKind             { return &obj.TypeMeta }
func (obj *CatalogEntryList) GetObjectKind() unversioned.ObjectKind         { return &obj.TypeMeta }
func (obj *CatalogEntryClaim) GetObjectKind() unversioned.ObjectKind        { return &obj.TypeMeta }
func (obj *CatalogEntryClaimList) GetObjectKind() unversioned.ObjectKind    { return &obj.TypeMeta }
func (obj *CatalogEntryProvider) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *CatalogEntryProviderList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
