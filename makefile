GO                      :=$(if $(V), go, @go)
RM                      :=$(if $(V), rm -rf, @rm -rf)
CAT                     :=$(if $(V), cat, @cat)
ECHO                    :=$(if $(V), echo, @echo)

coverage_file           :=coverage.txt

go_tmp                  :=$(shell find . -name "*.go" | grep -v main.go)
go_dirs                 :=$(sort $(dir $(go_tmp)))     # will remove duplicated.
go_coverages_files      :=$(addsuffix $(coverage_file), $(go_dirs))

cmd_dir                 :=cmd/cachalot
cmd_main                :=$(cmd_dir)/cachalot.go

TARGET                  :=$(cmd_dir)/cachalot

.PHONY: $(TARGET) test clean

$(TARGET):$(cmd_main)
	$(ECHO) "Build     $<"
	$(GO) build $<

test:$(go_coverages_files)
	$(ECHO) "Test      completed"

$(go_coverages_files):%/coverage.txt:%
	$(ECHO) "Test      $<"
	$(GO) test -covermode=atomic -coverprofile=$@ $<
	$(CAT) $@ >> coverage.txt
	$(RM) $@

clean:
	$(RM) $(TARGET)

