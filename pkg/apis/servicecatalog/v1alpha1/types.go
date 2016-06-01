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

	Catalog         string `json:"catalog,omitempty" protobuf:"bytes,4,opt,name=catalog"`
	Description     string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	SourceNamespace string `json:"sourceNamespace,omitEmpty" protobuf:"bytes,5,opt,name=sourceNamespace"`
}

type CatalogEntryList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogEntry `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogPosting struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Catalog        string             `json:"catalog" protobuf:"bytes,2,opt,name=catalog"`
	Description    string             `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	LocalResources *LocalResourceSpec `json:"localResources,omitempty" protobuf:"bytes,4,opt,name=localResources"`
}

type LocalResourceSpec struct {
	Items []v1.ObjectReference `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"`
}

type CatalogPostingList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogPosting `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogClaim struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   CatalogClaimSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status CatalogClaimStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
}

type CatalogClaimList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogClaim `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CatalogClaimSpec struct {
	Catalog string `json:"catalog" protobuf:"bytes,1,opt,name=catalog"`
	Entry   string `json:"entry" protobuf:"bytes,2,opt,name=entry"`
}

type CatalogClaimStatus struct {
	State            CatalogClaimState    `json:"state" protobuf:"bytes,1,opt,name=state,casttype=CatalogClaimState"`
	CreatedResources []v1.ObjectReference `json:"createdResources,omitempty" protobuf:"bytes,2,rep,name=createdResources"`
}

type CatalogClaimState string

const (
	CatalogClaimStateNew       CatalogClaimState = "New"
	CatalogClaimStatePending   CatalogClaimState = "Pending"
	CatalogClaimStateCompleted CatalogClaimState = "Completed"
)
