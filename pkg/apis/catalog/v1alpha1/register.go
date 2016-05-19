package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/runtime"
	versionedwatch "k8s.io/kubernetes/pkg/watch/versioned"
)

const GroupName = "catalog"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1alpha1"}

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
		&v1.ListOptions{},
		&v1.DeleteOptions{},
	)
	versionedwatch.AddToGroupVersion(scheme, SchemeGroupVersion)
}

func (obj *Catalog) GetObjectKind() unversioned.ObjectKind          { return &obj.TypeMeta }
func (obj *CatalogList) GetObjectKind() unversioned.ObjectKind      { return &obj.TypeMeta }
func (obj *CatalogEntry) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *CatalogEntryList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
