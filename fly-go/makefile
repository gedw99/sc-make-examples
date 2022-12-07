# Simple golang server with a gio web app

# https://fly.io/blog/hugo-s-there-flying-with-hugo-and-caddy/

# use caddy so we can run a Webserver to server it.
# can embed the WASM filesinside 

#https://github.com/joeblew99/gio-make-examples

MAKE_FSPATH=./../../gio-make/assets/make
include $(MAKE_FSPATH)/help.mk
include $(MAKE_FSPATH)/env.mk
include ./.env
include $(MAKE_FSPATH)/git.mk
include $(MAKE_FSPATH)/go.mk
include $(MAKE_FSPATH)/gio.mk
include $(MAKE_FSPATH)/caddy.mk
include $(MAKE_FSPATH)/fly.mk

#GIT_NAME=gio-make-examples
#GIT_ORG=joeblew99
#GIT_SERVER=github.com/$(GIT_ORG)/$(GIT_NAME)
#GIT_BRANCH=main

# !! toggle to work with the 2 different git remotes available.
#GIT_WORK_FSPATH=$(PWD)
#GIT_WORK_FSPATH=$(PWD)/$(GIT_REPO_ORIGIN_NAME)
#GIT_WORK_FSPATH=$(PWD)/$(GIT_REPO_UPSTREAM_NAME)

GIT_REPO_FSPATH=$(PWD)

GO_SRC_NAME=serve
GO_SRC_FSPATH=$(GIT_REPO_FSPATH)/cmd/$(GO_SRC_NAME)

GIO_SRC_NAME=gui
GIO_SRC_FSPATH=$(GIT_REPO_FSPATH)/cmd/$(GIO_SRC_NAME)

CADDY_SRC_FSPATH=$(PWD)

FLY_SRC_FSPATH=$(PWD)

all: this-gui-build this-server-build 

this-gui-build:
	$(MAKE) gio-build

	# copy web into server
	@echo -- copying web build $(GIO_BUILD_WEB_PATH) into $(GO_BUILD_FSPATH)

this-server-build:
	$(MAKE) go-build

this-server-run:
	$(MAKE) go-run



