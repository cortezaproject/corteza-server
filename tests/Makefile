include ../Makefile.inc

# Run go test cmd with flags, eg:
# $> make test.integration TEST_FLAGS="-v"
# $> make test.integration TEST_FLAGS="-v -run SpecialTest"
TEST_FLAGS ?=

COVER_MODE    ?= count
COVER_PROFILE ?= .cover.out
COVER_FLAGS   ?= -covermode=$(COVER_MODE)  -coverprofile=$(COVER_PROFILE)

# Cover package maps for tests tasks
COVER_PKGS_system      = ../system/...
COVER_PKGS_compose     = ../compose/...
COVER_PKGS_federation  = ../federation/...
COVER_PKGS_automation  = ../automation/...
COVER_PKGS_pkg         = ../pkg/...
COVER_PKGS_all         = $(COVER_PKGS_pkg),$(COVER_PKGS_system),$(COVER_PKGS_compose),$(COVER_PKGS_federation),$(COVER_PKGS_automation)
COVER_PKGS_integration = $(COVER_PKGS_all)

TEST_SUITE_pkg         = ../pkg/...
TEST_SUITE_services    = ../compose/... ../system/... ../federation/... ../auth/... ../automation/...
TEST_SUITE_unit        = $(TEST_SUITE_pkg) $(TEST_SUITE_services)
TEST_SUITE_integration = ../tests/...
TEST_SUITE_store       = ../store/tests/...
TEST_SUITE_all         = $(TEST_SUITE_unit) $(TEST_SUITE_integration) $(TEST_SUITE_store)


.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-\\.%]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: test
test: $(GOTEST) ## run unit tests
	$(GOTEST) $(TEST_FLAGS) $(TEST_SUITE_all)

.PHONY: test.coverprofile.%
test.coverprofile.%: ## adds -coverprofile flag to test flags
	@ $(MAKE) test.cover.$* TEST_FLAGS="$(TEST_FLAGS) -coverprofile=$(COVER_PROFILE)"

.PHONY: test.cover.%
test.cover.%: ## adds -coverpkg flag
	@ $(MAKE) test.$* TEST_FLAGS="$(TEST_FLAGS) -coverpkg=$(COVER_PKGS_$*)"

.PHONY: test.unit
test.unit: $(GOTEST) ## run unit tests
	$(GOTEST) $(TEST_FLAGS) $(TEST_SUITE_unit)

.PHONY: test.unit.%
test.unit.%: $(GOTEST) ## run <dir> unit tests
	$(GOTEST) $(TEST_FLAGS) ../$*/...

.PHONY: test.integration
test.integration: $(GOTEST) ## run integration tests
	$(GOTEST) $(TEST_FLAGS) $(TEST_SUITE_integration)

.PHONY: test.integration.%
test.integration.%: $(GOTEST) ## run <dir> integration tests
	$(GOTEST) $(TEST_FLAGS) ../tests/$*/...

.PHONY: test.store
test.store: $(GOTEST) ## run store tests
	$(GOTEST) $(TEST_FLAGS) $(TEST_SUITE_store)

.PHONY: test.store.%
test.store.%: $(GOTEST) ## run <dir> store tests
	$(GOTEST) $(TEST_FLAGS) ../store/tests/$*/...
