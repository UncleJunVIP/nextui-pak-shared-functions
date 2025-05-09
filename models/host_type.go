package models

import "qlova.tech/sum"

type HostType struct {
	APACHE,
	NGINX,
	ROMM,
	MEGATHREAD,
	CUSTOM sum.Int[HostType]
}

var HostTypes = sum.Int[HostType]{}.Sum()
