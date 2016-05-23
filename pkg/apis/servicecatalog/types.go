package servicecatalog

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

// +genclient=true,nonNamespaced=true

type CatalogEntry struct {
	unversioned.TypeMeta
	api.ObjectMeta

	Catalog         string
	Description     string
	SourceNamespace string

	// what gets created when this entry is claimed and provisioned, e.g.:
	// type: secret
	// name: dbpassword
	Output []api.ObjectReference
}

type CatalogEntryList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogEntry
}

// +genclient=true

type CatalogPosting struct {
	unversioned.TypeMeta
	api.ObjectMeta

	Catalog     string
	Description string
	/* e.g.
	{
		"targets": [
			{
				"from": "secret/oracle-password",
				"as": "dbpassword"
			}
		]
	}
	*/
	Data map[string]interface{}
}

type CatalogPostingList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogPosting
}

// +genclient=true

type CatalogClaim struct {
	unversioned.TypeMeta
	api.ObjectMeta

	Spec   CatalogClaimSpec
	Status CatalogClaimStatus
}

type CatalogClaimList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	Items []CatalogClaim
}

type CatalogClaimSpec struct {
	Catalog        string
	Entry          string
	ResourcePrefix string
}

type CatalogClaimStatus struct {
	State            CatalogClaimState
	CreatedResources []api.ObjectReference
}

type CatalogClaimState string

const (
	CatalogClaimStateNew       CatalogClaimState = "New"
	CatalogClaimStatePending   CatalogClaimState = "Pending"
	CatalogClaimStateCompleted CatalogClaimState = "Completed"
	CatalogClaimStateError     CatalogClaimState = "Error"
)
