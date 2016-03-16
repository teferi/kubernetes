package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
)

type Catalog struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`
}

type CatalogList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`

	Items []Catalog `json:"items"`
}

type CatalogEntry struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`

	Catalog     v1.LocalObjectReference `json:"catalog"`
	Description string                  `json:"description,omitempty"`

	Reference v1.ObjectReference `json:"reference"`
	Data      map[string]string  `json:"data,omitempty"`
}

type CatalogEntryList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`

	Items []CatalogEntry `json:"items"`
}

type CatalogEntryClaim struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`

	Spec   CatalogEntryClaimSpec   `json:"spec"`
	Status CatalogEntryClaimStatus `json:"status,omitempty"`
}

type CatalogEntryClaimList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`

	Items []CatalogEntryClaim `json:"items"`
}

type CatalogEntryClaimSpec struct {
	Catalog string `json:"catalog"`
	Entry   string `json:"entry"`
}

type CatalogEntryClaimStatus struct {
	State            CatalogEntryClaimState `json:"state"`
	ProvisionedItems []v1.ObjectReference   `json:"provisionedItems,omitempty"`
}

type CatalogEntryClaimState string

const (
	CatalogEntryClaimStateNew         CatalogEntryClaimState = "New"
	CatalogEntryClaimStateAdmitted    CatalogEntryClaimState = "Admitted"
	CatalogEntryClaimStateRejected    CatalogEntryClaimState = "Rejected"
	CatalogEntryClaimStateProvisioned CatalogEntryClaimState = "Provisioned"
)

type CatalogEntryProvider struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`

	// servicebroker.appdirect.mycompany.com
	URL         string                   `json:"url,omitempty"`
	Credentials *v1.LocalObjectReference `json:"credentials,omitempty"`
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
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`

	Items []CatalogEntryProvider `json:"items"`
}
