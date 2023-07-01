module main

go 1.18

require github.com/Brian44913/PublicPack/hardware v0.0.0

replace github.com/Brian44913/PublicPack/hardware => ./hardware

require github.com/Brian44913/PublicPack/code v0.0.0

replace github.com/Brian44913/PublicPack/code => ./code

require github.com/Brian44913/PublicPack/other v0.0.0

replace github.com/Brian44913/PublicPack/other => ./other

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/safchain/ethtool v0.3.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/sys v0.9.0 // indirect
)
