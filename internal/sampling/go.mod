module internal/sampling

go 1.18

replace github.com/Incarnation-p-lee/cachalot/pkg/snapshot => ../../pkg/snapshot

replace internal/options => ../../internal/options

replace github.com/Incarnation-p-lee/cachalot/pkg/assert => ../../pkg/assert

require (
	github.com/Incarnation-p-lee/cachalot/pkg/assert v0.0.0-00010101000000-000000000000
	github.com/Incarnation-p-lee/cachalot/pkg/snapshot v0.0.0-00010101000000-000000000000
	internal/options v0.0.0-00010101000000-000000000000
	internal/utils v0.0.0-00010101000000-000000000000
)

replace internal/utils => ../../internal/utils
