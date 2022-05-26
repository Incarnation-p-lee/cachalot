module cmdflags

go 1.18

replace (
	internal/options => ../../internal/options
	pkg/assert => ../../pkg/assert
)

require (
	internal/options v0.0.0-00010101000000-000000000000
	pkg/assert v0.0.0-00010101000000-000000000000
)
