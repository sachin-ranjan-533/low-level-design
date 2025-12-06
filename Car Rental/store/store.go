package store

import (
	"car-rental/location"
	"car-rental/reservation"
	"car-rental/vehicle_inventory"
)

type Store struct {
	Id                int
	Name              string
	Location          *location.Location
	Reservations      []*reservation.Reservation
	VehicleInventorys []vehicle_inventory.VehicleInventory
}

func NewStore(id int, name string, Location *location.Location) *Store {
	return &Store{
		Id:       id,
		Name:     name,
		Location: Location,
	}
}

func (s *Store) AddReservation(reservation *reservation.Reservation) {
	s.Reservations = append(s.Reservations, reservation)
}

func (s *Store) InitializeInventory(VehicleInventory vehicle_inventory.VehicleInventory) {
	s.VehicleInventorys = append(s.VehicleInventorys, VehicleInventory)
}

func (s *Store) GetVehicleInventory(vehicleType string) vehicle_inventory.VehicleInventory {
	for _, inventory := range s.VehicleInventorys {
		if inventory.GetVehicleType() == vehicleType {
			return inventory
		}
	}
	return nil
}
