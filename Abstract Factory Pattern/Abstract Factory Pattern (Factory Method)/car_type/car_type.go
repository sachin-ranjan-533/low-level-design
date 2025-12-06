package cartype

import (
	carexterior "abstract-factory/car_exterior"
	carinterior "abstract-factory/car_interior"
)

type CarType interface {
	GetExterior() carexterior.CarExterior
	GetInterior() carinterior.CarInterior
}
