include Makefile.inc


# Dev Support apps settings
DEV_MINIO_PORT        ?= 9000
DEV_MAILHOG_SMTP_ADDR ?= 1025
DEV_MAILHOG_HTTP_ADDR ?= 8025

GIN_ARG_LADDR ?= localhost
GIN_ARGS      ?= --laddr $(GIN_ARG_LADDR) --immediate


DOCKER                ?= docker


.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-\\.%]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: help.build
help.build: ## show build make targets
	@ $(MAKE) --directory=build help

.PHONY: help.test
help.test: ## show test make targets
	@ $(MAKE) --directory=tests help

########################################################################################################################
# Building & packaging

build%: ## run build targets
	@ $(MAKE) --directory=build build$*



#build: $(BUILD_DEST_DIR)/$(BUILD_BIN_NAME) ## build corteza-server
#
#$(BUILD_DEST_DIR)/$(BUILD_BIN_NAME):
#		GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build $(LDFLAGS) -o $@ cmd/corteza/main.go
#
#release: build $(BUILD_DEST_DIR)/$(RELEASE_NAME)
#
#$(BUILD_DEST_DIR)/$(RELEASE_NAME):
#	@ mkdir -p $(RELEASE_BASEDIR) $(RELEASE_BASEDIR)/bin
#	@ cp $(RELEASE_EXTRA_FILES) $(RELEASE_BASEDIR)/
#	@ cp -r provision $(RELEASE_BASEDIR)
#	@ rm -f $(RELEASE_BASEDIR)/provision/README.adoc $(RELEASE_BASEDIR)/provision/update.sh
#	@ cp $(BUILD_DEST_DIR)/$(BUILD_BIN_NAME) $(RELEASE_BASEDIR)/bin/$(BUILD_FLAVOUR)-server
#	tar -C $(dir $(RELEASE_BASEDIR)) -czf $(BUILD_DEST_DIR)/$(RELEASE_NAME) $(notdir $(RELEASE_BASEDIR))
#
#release-clean:
#	rm -rf $(BUILD_DEST_DIR)/$(BUILD_BIN_NAME)
#	rm -rf $(BUILD_DEST_DIR)/$(RELEASE_NAME)


#######################################################################################################################
# Quality Assurance

test%: ## run test targets
	@ $(MAKE) --directory=tests test$*


#.PHONY: test
#test: ## run all tests
#	@ $(MAKE) --directory=tests test
#
#test.coverprofile.%:  ## adds -coverprofile flag to test flags
#	@ $(MAKE) --directory=tests test.coverprofile.$*
#
#test.cover.%: ## adds -coverpkg flag
#	@ $(MAKE) --directory=tests test.cover.$*
#
#.PHONY: test.unit
#test.unit: ## run unit tests
#	@ $(MAKE) --directory=tests test.unit
#
#.PHONY: test.unit.%
#test.unit.%: ## run <dir> unit tests
#	@ $(MAKE) --directory=tests test.unit.$*
#
#.PHONY: test.integration
#test.integration: ## run integration tests
#	@ $(MAKE) --directory=tests test.integration
#
#.PHONY: test.integration.%
#test.integration.%: ## run <dir> integration tests
#	@ $(MAKE) --directory=tests test.integration.$*
#
#.PHONY: test.store
#test.store: ## run store tests
#	@ $(MAKE) --directory=tests test.store
#
#.PHONY: test.store.%
#test.store.%: ## run <dir> store tests
#	@ $(MAKE) --directory=tests test.store.$*

vet: ## run go vet
	$(GO) vet ./...

critic: $(GOCRITIC) ## run gocritic
	$(GOCRITIC) check-project .

staticcheck: $(STATICCHECK) ## run staticcheck
	$(STATICCHECK) ./pkg/... ./system/... ./compose/... ./automation/...

qa: vet critic test ## run quality assurance

mocks: $(MOCKGEN) ## create mocks
	$(MOCKGEN) -package mail -source pkg/mail/mail.go -destination pkg/mail/mail_mock_test.go


########################################################################################################################
# Development

watch: $(GIN)
	$(GIN) $(GIN_ARGS) --build cmd/corteza run -- serve

# Development helper - reruns test when files change
#
# make watch.test.unit
# make watch.test.pkg
# make watch.test.all
# make watch.test.pkg TEST_FLAGS="-v"

watch.test.%: $(FSWATCH)
	( make test.$* || exit 0 ) && ( $(FSWATCH) -o . | xargs -n1 -I{} make test.$* )

watch.test: watch.test.all

watch.codegen: $(CODEGEN)
	@ $(CODEGEN) -w -v

realize: watch # BC

mailhog.up:
	$(DOCKER) run --rm --publish $(DEV_MAILHOG_HTTP_ADDR):8025 --publish $(DEV_MAILHOG_SMTP_ADDR):1025 mailhog/mailhog

minio.up:
	# Runs temp minio server
	# No volume mounts because we do not want the data to persist
	$(DOCKER) run --rm --publish $(DEV_MINIO_PORT):9000 --env-file .env minio/minio server /data

# codegen: $(PROTOGEN)
codegen: $(CODEGEN)
	@ $(CODEGEN) -v

clean.codegen:
	rm -f $(CODEGEN)

webapp:
	@ $(MAKE) --directory=webapp
