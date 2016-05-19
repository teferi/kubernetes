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

	Catalog     string `json:"catalog,omitempty" protobuf:"bytes,4,opt,name=catalog"`
	Description string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
}

type CatalogEntryList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CatalogEntry `json:"items" protobuf:"bytes,2,rep,name=items"`
}
