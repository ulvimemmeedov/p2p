module main

go 1.17

replace github.com/ulvimemmeedov/p2p/client => ./client/
replace github.com/ulvimemmeedov/p2p/tools => ./tools/
replace github.com/ulvimemmeedov/p2p/contracts => ./contracts/
require (
	github.com/ulvimemmeedov/p2p/client v0.0.0-00010101000000-000000000000 // indirect
	github.com/ulvimemmeedov/p2p/tools v0.0.0-00010101000000-000000000000
	github.com/ulvimemmeedov/p2p/contracts v0.0.0-00010101000000-000000000000

	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)
