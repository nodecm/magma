SHELL := /bin/bash

ifndef MAGMA_MODULES
$(error MAGMA_MODULES is undefined)
endif

##########################################
## Some Useful Directory/File Variables ##
##########################################
# Allow MAGMA_ROOT and PLUGIN_DIR to be set from environment to support travis
ifndef MAGMA_ROOT
MAGMA_ROOT := /home/$(USER)/magma
endif
export MAGMA_ROOT

SWAGGER_ROOT := $(MAGMA_ROOT)/orc8r/cloud/deploy/files/static/apidocs
SWAGGER_YML := $(SWAGGER_ROOT)/swagger.yml
SWAGGER_TEMPLATE :=$(SWAGGER_ROOT)/swagger-template.yml
SWAGGER_TEMP_GEN := $(SWAGGER_ROOT)/temp
SWAGGER_COMMON := swagger-common.yml
export SWAGGER_ROOT
export SWAGGER_TEMPLATE
export SWAGGER_TEMP_GEN
export SWAGGER_COMMON

BIN_DIR := $(GOPATH)/bin

CERTS_DIR := $(MAGMA_ROOT)/.cache/test_certs
ifndef PLUGIN_DIR
PLUGIN_DIR := /home/$(USER)/go/plugins
endif
export PLUGIN_DIR

############################
## Declare Module Targets ##
############################
define add_module
	$(eval MODULE := $1)
	$(eval include $(MODULE)/cloud/go/defs.mk)

	BUILD_LIST += $(MODULE)_build
	CLEAN_LIST += $(MODULE)_clean
	CLEAN_GEN_LIST += $(MODULE)_cleangen
	COVER_LIST += $(MODULE)_cover
	DOWNLOAD_LIST += $(MODULE)_download
	FMT_LIST += $(MODULE)_fmt
	GEN_LIST += $(MODULE)_gen
	LINT_LIST += $(MODULE)_lint
	PLUGIN_LIST += $(MODULE)_plugin
	SWAGGER_LIST += $(MODULE)_swagger
	TEST_LIST += $(MODULE)_test
	TIDY_LIST += $(MODULE)_tidy
	TOOLS_LIST += $(MODULE)_tools
	VET_LIST += $(MODULE)_vet

	SERVICES += $(foreach srvs, $($(MODULE)_SERVICE_NAMES), $(srvs))
endef
_ := $(foreach module, $(MAGMA_MODULES), $(eval $(call add_module,$(module))))


###################
## General Rules ##
###################

RUN_LIST = $(foreach srv, $(SERVICES), $(srv)_run) obsidian_run
run: build $(RUN_LIST) create_admin_oper
$(RUN_LIST): %_run:
	sudo systemctl restart magma@$*

precommit: fmt test vet swagger_validate

build: plugin $(BUILD_LIST) swagger
$(BUILD_LIST): %_build:
	make -C $*/cloud/go build

clean: $(CLEAN_LIST)
$(CLEAN_LIST): %_clean:
	make -C $*/cloud/go clean

clean_gen: $(CLEAN_GEN_LIST)
	rm -f $(SWAGGER_YML);
	rm -rf $(SWAGGER_TEMP_GEN)
$(CLEAN_GEN_LIST): %_cleangen:
	make -C $*/cloud/go clean_gen

COVER_DIR := /var/tmp
export COVER_DIR
cover: $(COVER_LIST)
$(COVER_LIST): %_cover:
	make -C $*/cloud/go cover

download: $(DOWNLOAD_LIST)
$(DOWNLOAD_LIST): %_download:
	make -C $*/cloud/go download

fmt: $(FMT_LIST)
$(FMT_LIST): %_fmt:
	make -C $*/cloud/go fmt

gen: tools $(GEN_LIST)
$(GEN_LIST): %_gen:
	make -C $*/cloud/go gen

lint: $(LINT_LIST)
$(LINT_LIST): %_lint:
	make -C $*/cloud/go lint

plugin: $(PLUGIN_LIST)
$(PLUGIN_LIST): %_plugin:
	make -C $*/cloud/go plugin

test: $(TEST_LIST)
$(TEST_LIST): %_test:
	make -C $*/cloud/go test

tidy: $(TIDY_LIST)
$(TIDY_LIST): %_tidy:
	make -C $*/cloud/go tidy

tools: $(TOOLS_LIST)
$(TOOLS_LIST): %_tools:
	make -C $*/cloud/go tools

vet: $(VET_LIST)
$(VET_LIST): %_vet:
	make -C $*/cloud/go vet

swagger: swagger_directories tools $(SWAGGER_LIST)
	combine_swagger --inp=$(SWAGGER_TEMP_GEN) --common=$(SWAGGER_ROOT)/$(SWAGGER_COMMON) --out=$(SWAGGER_YML)
	rm -rf $(SWAGGER_TEMP_GEN)

swagger_directories:
	rm -rf $(SWAGGER_TEMP_GEN)
	mkdir $(SWAGGER_TEMP_GEN)

$(SWAGGER_LIST): %_swagger:
	make -C $*/cloud/go copy_swagger_files

swagger_validate: swagger
	swagger validate $(SWAGGER_YML)

status:  # Get the status of all Magma services
	sudo systemctl status magma@*

stop:  # Stops all Magma services
	sudo systemctl stop magma@*

##################
## Misc Targets ##
##################

create_admin_oper:  # Creates new admin operator and adds to certifier
	test -s $(CERTS_DIR)/admin_operator.pem || { $(BIN_DIR)/accessc add-admin -cert $(CERTS_DIR)/admin_operator admin_operator; }
	test -s $(CERTS_DIR)/admin_operator.pfx || { openssl pkcs12 -passout pass:magma -export -out $(CERTS_DIR)/admin_operator.pfx -inkey $(CERTS_DIR)/admin_operator.key.pem -in $(CERTS_DIR)/admin_operator.pem; }

restore_admin_operator:  # Adds existing admin operator to certifier
	$(BIN_DIR)/accessc add-existing -admin -cert $(CERTS_DIR)/admin_operator.pem admin_operator

NGHTTPX_INST_LIST = open clientcert obsidian
restart_proxy: # Restart/Start all nghttpx proxy instances
	$(foreach s, $(NGHTTPX_INST_LIST), sudo systemctl restart nghttpx@$(s);)

restart_services: $(RUN_LIST) # Restart/Start all magma services

restart: restart_proxy restart_services # Restart proxies & all magma services

# Check if all magma services are healthy
CHECK_LIST = $(foreach srv, $(SERVICES), $(srv)_check) obsidian_check
check: check_sleep $(CHECK_LIST)
check_sleep:
	sleep 5
$(CHECK_LIST): %_check:
	sudo systemctl is-active magma@$* || (sudo journalctl -u magma@$* -n 20 --no-pager && false)

list:  # List all commands
	@echo -e "\nAvailable commands:\n"
	@grep '^[[:lower:]].*:.*#.*' Makefile | awk -F ':|#' '{ printf "\t%-30s%s\n", $$1, $$3 }'
	@echo -e "\n"

# Run target for travis - don't create admin certs since we don't spin up postgres
travis_run: build $(RUN_LIST)
