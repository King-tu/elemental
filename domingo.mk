# ------------------------------------------------
# Copyright (C) 2016 Aporeto Inc.
#
# File  : Makefile
#
# Author: alex@aporeto.com, antoine@aporeto.com
# Date  : 2016-03-8
#
# ------------------------------------------------

MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash

APOMOCK_FOLDER := .apomock
APOMOCK_PACKAGES := $(shell if [ -f .apo.mock ]; then cat .apo.mock; fi)
NOVENDOR := $(shell glide novendor)

DIRS_WITH_MAKEFILES := $(sort $(dir $(wildcard */Makefile)))

# Remove directories which have Makefiles from being tested by top level
NOTEST_DIRS := $(DIRS_WITH_MAKEFILES)
NOTEST_DIRS := $(addsuffix ...,$(NOTEST_DIRS))
NOTEST_DIRS := $(addprefix ./,$(NOTEST_DIRS))

# Remove directories which are mock directories from being tested by top level
MOCK_DIRS := $(sort $(dir $(wildcard ./mock*)))
MOCK_DIRS := $(addsuffix ...,$(MOCK_DIRS))
MOCK_DIRS := $(addprefix ./,$(MOCK_DIRS))

TEST_DIRS := $(filter-out $(NOTEST_DIRS),$(NOVENDOR))
TEST_DIRS := $(filter-out $(MOCK_DIRS),$(TEST_DIRS))

## Update

apomakeupdate:
	@echo "# Running apomakeupdate in" $(PWD)
	@echo "NOT IMPLEMENTED YET"

## Dependencies

apoinit:
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apoinit && popd;)
	@echo "# Running apoinit in" $(PWD)
	go get ./...
	@if [ -f glide.lock ]; then glide install && glide update; fi


## Testing

apotest: apolint apomock
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apotest && popd;)
	@echo "# Running test in" $(PWD)
	[ -z "${TEST_DIRS}" ] || go vet ${TEST_DIRS}
	[ -z "${TEST_DIRS}" ] || go test -v -race -cover ${TEST_DIRS}

apolint:
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apolint && popd;)
	@echo "# Running lint in" $(PWD)
	golint .

apomock: apoclean_apomock apoclean_vendor
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apomock && popd;)
	@echo "# Running apomock in" $(PWD)
	rm -rf ${APOMOCK_FOLDER}
	mkdir -p ${APOMOCK_FOLDER}
	touch ${APOMOCK_FOLDER}/apomock.log
	kennebec --package="${APOMOCK_PACKAGES}" --output-dir=${APOMOCK_FOLDER} -v=4 -logtostderr=true>>${PWD}/${APOMOCK_FOLDER}/apomock.log 2>&1
	@if [ ! -d vendor ]; then mkdir vendor; fi;
	@if [ -d ${APOMOCK_FOLDER} ]; then cp -r ${APOMOCK_FOLDER}/* vendor; fi;


## Cleaning

apoclean_vendor:
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apoclean_vendor && popd;)
	@echo "# Running apoclean_vendor in" $(PWD)
	rm -rf vendor

apoclean_apomock:
	@$(foreach dir,$(DIRS_WITH_MAKEFILES),pushd $(dir) && make apoclean_apomock && popd;)
	@echo "# Running apoclean_apomock in" $(PWD)
	rm -rf ${APOMOCK_FOLDER}


## Docker Test Container
PROJECT_OWNER?=github.com/aporeto-inc
PROJECT_NAME?=my-super-project
BUILD_NUMBER?=latest

define DOCKER_FILE
FROM 926088932149.dkr.ecr.us-west-2.amazonaws.com/domingo
MAINTAINER Antoine Mercadal <antoine@aporeto.com>
ADD . /go/src/$(PROJECT_OWNER)/$(PROJECT_NAME)
WORKDIR /go/src/$(PROJECT_OWNER)/$(PROJECT_NAME)
endef
export DOCKER_FILE


create_build_container:
	echo "$$DOCKER_FILE" > .dockerfile-test
	docker build --file .dockerfile-test -t $(PROJECT_NAME)-build-image:$(BUILD_NUMBER) .

run_build_container:
	docker run --rm $(PROJECT_NAME)-build-image:$(BUILD_NUMBER)

clean_build_container:
	docker rmi $(PROJECT_NAME)-build-image:$(BUILD_NUMBER)
	rm -f .dockerfile-test