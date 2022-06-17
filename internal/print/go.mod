module internal/print

go 1.18

replace github.com/Incarnation-p-lee/cachalot/pkg/snapshot => ../../pkg/snapshot

replace github.com/Incarnation-p-lee/cachalot/pkg/assert => ../../pkg/assert

replace internal/options => ../../internal/options

require (
	github.com/Incarnation-p-lee/cachalot/pkg/assert v0.0.0-00010101000000-000000000000
	github.com/Incarnation-p-lee/cachalot/pkg/snapshot v0.0.0-00010101000000-000000000000
	internal/options v0.0.0-00010101000000-000000000000
)
