GO                      :=$(if $(V), go, @go)
RM                      :=$(if $(V), rm -rf, @rm -rf)

go_tmp                  :=$(shell find . -name "*.go")
go_dirs                 :=$(sort $(dir $(go_tmp)))     # will remove duplicated.
go_coverages            :=$(addsuffix coverage, $(go_dirs))

.PHONY: test clean

test:$(go_coverages)
	@echo "Test      completed"

$(go_coverages):%/coverage:%/
	$(GO) test -covermode=atomic -coverprofile=$@ ./$<
	@cat $@ >> coverage.txt
	$(RM) $(go_coverages)

