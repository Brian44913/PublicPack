module main

go 1.18

require github.com/Brian44913/PublicPackage/hardware v0.0.0

replace github.com/Brian44913/PublicPackage/hardware => ./hardware

require github.com/Brian44913/PublicPackage/code v0.0.0

replace github.com/Brian44913/PublicPackage/code => ./code

require github.com/Brian44913/PublicPackage/other v0.0.0

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/safchain/ethtool v0.3.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/Brian44913/PublicPackage/other => ./other
