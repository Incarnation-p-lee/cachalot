CD                      :=cd
GO                      :=go
RM                      :=rm -f
CAT                     :=cat

coverage_file           :=coverage.txt

go_tmp                  :=$(shell find . -name "*.go")
go_dirs                 :=$(sort $(dir $(go_tmp)))     # will remove duplicated.
go_coverages_files      :=$(addsuffix $(coverage_file), $(go_dirs))

.PHONY: test clean

test:$(go_coverages_files)
	@echo "Test      completed"

$(go_coverages_files):%/coverage.txt:%/
	@echo "Test      $<"
	@$(CD) $< \
		&& $(GO) test -covermode=atomic -coverprofile=$(coverage_file) . \
		&& $(CD) - > /dev/null
	@$(CAT) $@ >> coverage.txt && $(RM) $@

