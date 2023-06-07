module main

go 1.18

require github.com/Brian44913/PublicPackage/hardware v0.0.0

replace github.com/Brian44913/PublicPackage/hardware => ./hardware

require github.com/Brian44913/PublicPackage/code v0.0.0

replace github.com/Brian44913/PublicPackage/code => ./code

require github.com/Brian44913/PublicPackage/other v0.0.0

require (
	github.com/safchain/ethtool v0.3.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/Brian44913/PublicPackage/other => ./other
