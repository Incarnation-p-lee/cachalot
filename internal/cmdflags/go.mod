module internal/cmdflags

go 1.18

replace internal/options => ../options

require (
	internal/options v0.0.0-00010101000000-000000000000
	pkg/assert v0.0.0-00010101000000-000000000000
)

replace pkg/assert => ../../pkg/assert
