package models

import "qlova.tech/sum"

type HostType struct {
	APACHE,
	NGINX,
	SMB,
	ROMM,
	MEGATHREAD,
	CUSTOM sum.Int[HostType]
}

var HostTypes = sum.Int[HostType]{}.Sum()
