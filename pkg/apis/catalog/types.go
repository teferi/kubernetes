package catalog

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

// +genclient=true

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

	Catalog     api.LocalObjectReference
	Description string

	Reference api.ObjectReference
	Data      map[string]string
}

type CatalogEntryList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogEntry
}

// +genclient=true

type CatalogEntryClaim struct {
	unversioned.TypeMeta
	api.ObjectMeta

	Spec   CatalogEntryClaimSpec
	Status CatalogEntryClaimStatus
}

type CatalogEntryClaimList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogEntryClaim
}

type CatalogEntryClaimSpec struct {
	Catalog string
	Entry   string
}

type CatalogEntryClaimStatus struct {
	State            CatalogEntryClaimState
	ProvisionedItems []api.ObjectReference
}

type CatalogEntryClaimState string

const (
	CatalogEntryClaimStateNew         CatalogEntryClaimState = "New"
	CatalogEntryClaimStateAdmitted    CatalogEntryClaimState = "Admitted"
	CatalogEntryClaimStateRejected    CatalogEntryClaimState = "Rejected"
	CatalogEntryClaimStateProvisioned CatalogEntryClaimState = "Provisioned"
)

type CatalogEntryProvider struct {
	unversioned.TypeMeta
	api.ObjectMeta

	// servicebroker.appdirect.mycompany.com
	URL         string
	Credentials *api.LocalObjectReference
	// ...

	/*
		- get catalog entries
		- create service instance(name, data[][])
		- bind instance
		- unbind instance
		- delete service instance
	*/
}

type CatalogEntryProviderList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogEntryProvider
}
