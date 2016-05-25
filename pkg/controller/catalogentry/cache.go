package catalogentry

import (
	"k8s.io/kubernetes/pkg/apis/servicecatalog"
	"k8s.io/kubernetes/pkg/client/cache"
	//"k8s.io/kubernetes/pkg/labels"
)

// StoreToCatalogPostingLister gives a store List and Exists methods. The store must contain only
// CatalogPostings.
type StoreToCatalogPostingLister struct {
	cache.Store
}

// Exists checks if the given CatalogPosting exists in the store.
func (s *StoreToCatalogPostingLister) Exists(cp *servicecatalog.CatalogPosting) (bool, error) {
	_, exists, err := s.Store.Get(cp)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// List lists all CatalogPostings in the store.
func (s *StoreToCatalogPostingLister) List() (cpList servicecatalog.CatalogPostingList, err error) {
	for _, cp := range s.Store.List() {
		cpList.Items = append(cpList.Items, *(cp.(*servicecatalog.CatalogPosting)))
	}
	return cpList, nil
}

func (s *StoreToCatalogPostingLister) GetCatalogPostings(catalog string) (cpList []servicecatalog.CatalogPosting, err error) {
	//var selector labels.Selector
	var cp servicecatalog.CatalogPosting

	for _, m := range s.Store.List() {
		cp = *m.(*servicecatalog.CatalogPosting)
		if cp.Catalog != catalog {
			continue
		}
		cpList = append(cpList, cp)
	}
	return
}
