package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
)

type Catalog struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

type CatalogList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Catalog `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogEntry struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Catalog     v1.LocalObjectReference `json:"catalog" protobuf:"bytes,2,opt,name=catalog"`
	Description string                  `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`

	Reference v1.ObjectReference `json:"reference" protobuf:"bytes,4,opt,name=reference"`
	Data      map[string]string  `json:"data,omitempty" protobuf:"bytes,5,rep,name=data"`
}

type CatalogEntryList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogEntry `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogEntryClaim struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   CatalogEntryClaimSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status CatalogEntryClaimStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type CatalogEntryClaimList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogEntryClaim `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogEntryClaimSpec struct {
	Catalog string `json:"catalog" protobuf:"bytes,1,opt,name=catalog"`
	Entry   string `json:"entry" protobuf:"bytes,2,opt,name=entry"`
}

type CatalogEntryClaimStatus struct {
	State            CatalogEntryClaimState `json:"state" protobuf:"bytes,1,opt,name=state,casttype=CatalogEntryClaimState"`
	ProvisionedItems []v1.ObjectReference   `json:"provisionedItems,omitempty" protobuf:"bytes,2,rep,name=provisionedItems"`
}

type CatalogEntryClaimState string

const (
	CatalogEntryClaimStateNew         CatalogEntryClaimState = "New"
	CatalogEntryClaimStateAdmitted    CatalogEntryClaimState = "Admitted"
	CatalogEntryClaimStateRejected    CatalogEntryClaimState = "Rejected"
	CatalogEntryClaimStateProvisioned CatalogEntryClaimState = "Provisioned"
)
