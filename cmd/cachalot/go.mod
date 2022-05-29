module main

go 1.18

replace internal/options => ../../internal/options

replace internal/cmdflags => ../../internal/cmdflags

replace internal/sampling => ../../internal/sampling

require (
	internal/cmdflags v0.0.0-00010101000000-000000000000
	internal/options v0.0.0-00010101000000-000000000000
	internal/sampling v0.0.0-00010101000000-000000000000
)

require github.com/Incarnation-p-lee/cachalot/pkg/snapshot v0.0.0-00010101000000-000000000000 // indirect

replace github.com/Incarnation-p-lee/cachalot/pkg/snapshot => ../../pkg/snapshot

replace github.com/Incarnation-p-lee/cachalot/pkg/assert => ../../pkg/assert
