GO                      :=$(if $(V), go, @go)
RM                      :=$(if $(V), rm -rf, @rm -rf)

go_tmp                  :=$(shell find . -type d -regex '\./[a-z_]+')
go_dirs                 :=$(addsuffix /,$(go_tmp))
go_coverages            :=$(addsuffix coverage,$(go_dirs))

.PHONY: test clean

test:$(go_coverages)
	@echo "Test      packages"

$(go_coverages):%/coverage:%/
	$(GO) test -covermode=atomic -coverprofile=$@ ./$<

clean:
	$(RM) $(go_coverages)

