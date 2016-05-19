package catalog

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

// +genclient=true,nonNamespaced=true

type Catalog struct {
	unversioned.TypeMeta
	api.ObjectMeta
}

type CatalogList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []Catalog
}

// +genclient=true

type CatalogEntry struct {
	unversioned.TypeMeta
	api.ObjectMeta

	Catalog     string
	Description string
}

type CatalogEntryList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogEntry
}
