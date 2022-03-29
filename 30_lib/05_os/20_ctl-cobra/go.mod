module main

go 1.17

require github.com/rebirthmonkey/lib/wkctl/cmd v0.0.1

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/rebirthmonkey/lib/wkctl/cmd => ./cmd
