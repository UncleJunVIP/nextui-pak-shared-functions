package models

import "qlova.tech/sum"

type HostType struct {
	ROMM,
	MEGATHREAD,
	APACHE sum.Int[HostType] // Apache is an internal type
}

var HostTypes = sum.Int[HostType]{}.Sum()
