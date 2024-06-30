package config

import "time"

var (
	RemoteAddressList = []string{"localhost:8000"}
	ZipfianConstant   = 0.9
	Latency           = 3 * time.Millisecond
)
