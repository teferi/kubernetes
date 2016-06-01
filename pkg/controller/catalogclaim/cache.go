package catalogentry

import (
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
)

// StoreToCatalogClaimLister gives a store List and Exists methods. The store must contain only
// CatalogClaims.
type StoreToCatalogClaimLister struct {
	cache.Store
}

// Exists checks if the given CatalogClaim exists in the store.
func (s *StoreToCatalogClaimLister) Exists(cc *servicecatalog.CatalogClaim) (bool, error) {
	_, exists, err := s.Store.Get(cc)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// List lists all CatalogClaims in the store.
func (s *StoreToCatalogClaimLister) List() (ccList servicecatalog.CatalogClaimList, err error) {
	for _, cc := range s.Store.List() {
		ccList.Items = append(ccList.Items, *(cc.(*servicecatalog.CatalogClaim)))
	}
	return ccList, nil
}

/*func (s *StoreToCatalogClaimLister) GetCatalogClaims(catalog string) (ccList []servicecatalog.CatalogClaim, err error) {
	//var selector labels.Selector
	var cc servicecatalog.CatalogClaim

	for _, m := range s.Store.List() {
		cc = *m.(*servicecatalog.CatalogClaim)
		if cc.Catalog != catalog {
			continue
		}
		ccList = append(ccList, cc)
	}
	return
}*/
