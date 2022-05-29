GO                      :=$(if $(V), go, @go)
RM                      :=$(if $(V), rm -rf, @rm -rf)
CAT                     :=$(if $(V), cat, @cat)
CD                      :=$(if $(V), cd, @cd)
ECHO                    :=$(if $(V), echo, @echo)

package_prefix          :=github.com/Incarnation-p-lee/cachalot/
coverage_file           :=coverage.txt

go_tmp                  :=$(shell find . -name "*.go" | grep -v cachalot.go)
go_dirs                 :=$(sort $(dir $(go_tmp)))     # will remove duplicated.
go_coverages_files      :=$(addsuffix $(coverage_file), $(go_dirs))

cmd_dir                 :=cmd/cachalot
cmd_main                :=$(cmd_dir)/cachalot.go
cmd                     :=$(subst .go, , $(cmd_main))

TARGET                  :=$(cmd_dir)/cachalot

.PHONY: cmd test clean

$(cmd):$(cmd_main)
	$(ECHO) "Build     $(TARGET)"
	$(CD) $(cmd_dir) && go build $(notdir $(cmd_main)) && cd -> /dev/null

test:$(go_coverages_files)
	$(ECHO) "Test      completed"

$(go_coverages_files):%/coverage.txt:%
	$(ECHO) "Test      $<"
	$(CD) $< && go test -covermode=atomic -coverprofile=$(notdir $@) \
		$(if $(filter pkg/%, $<), $(package_prefix), )$< \
		&& cd -> /dev/null
	$(CAT) $@ >> coverage.txt
	$(RM) $@

clean:
	$(RM) $(cmd) $(go_coverages_files) $(coverage_file)

