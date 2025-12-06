package store_manager

import (
	"car-rental/location"
	"car-rental/store"
)

type StoreManager struct {
	stores   []store.Store
	location *location.Location
}

func NewStoreManager(location *location.Location) *StoreManager {
	return &StoreManager{
		location: location,
	}
}

func (psm *StoreManager) AddStore(store store.Store) {
	psm.stores = append(psm.stores, store)
}

func (psm *StoreManager) GetStores() []store.Store {
	return psm.stores
}
