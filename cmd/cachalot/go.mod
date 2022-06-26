module main

go 1.18

replace internal/options => ../../internal/options

replace internal/cmdflags => ../../internal/cmdflags

replace internal/sampling => ../../internal/sampling

require (
	github.com/Incarnation-p-lee/cachalot/pkg/assert v0.0.0-00010101000000-000000000000
	internal/cmdflags v0.0.0-00010101000000-000000000000
	internal/options v0.0.0-00010101000000-000000000000
	internal/print v0.0.0-00010101000000-000000000000
	internal/sampling v0.0.0-00010101000000-000000000000
)

require (
	github.com/Incarnation-p-lee/cachalot/pkg/snapshot v0.0.0-00010101000000-000000000000 // indirect
	internal/utils v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/Incarnation-p-lee/cachalot/pkg/snapshot => ../../pkg/snapshot

replace github.com/Incarnation-p-lee/cachalot/pkg/assert => ../../pkg/assert

replace internal/print => ../../internal/print

replace internal/utils => ../../internal/utils
