# https://github.com/fly-apps/postgres-ha

MAKE_FSPATH=./../../gio-make/assets/make
include $(MAKE_FSPATH)/help.mk
include $(MAKE_FSPATH)/git.mk
include $(MAKE_FSPATH)/go.mk
include $(MAKE_FSPATH)/gio.mk
include $(MAKE_FSPATH)/caddy.mk
include $(MAKE_FSPATH)/fly.mk

FLY_APP=gedw99-postgres-ha

LIB_NAME=postgres-ha

print:
	$(MAKE) fly-print

dep:
	git clone https://github.com/fly-apps/postgres-ha
	#fly init YOURAPPNAME --import example-fly.toml

	git clone https://github.com/fergusstrange/embedded-postgres
dep-delete:
	rm -rf postgres-ha
	rm -rf embedded-postgres

dep-os:
	$(MAKE) fly-dep

fly-init:
	cd $(LIB_NAME) && flyctl init

init:
	$(FLY_BIN) auth login
	$(FLY_BIN) init $(FLY_APP) --import example-fly.toml



# https://fly.io/docs/reference/postgres/

include fly-db.mk

FLY_DB_APP_NAME=gedw99-db-test01
FLY_DB_CONSOLE_URL=https://fly.io/apps/$(FLY_DB_APP_NAME)

FLY_DB_ORG_NAME=personal
FLY_DB_REGION_NAME=fra

FLY_DB_CLUSTER_INIT_SIZE=1
FLY_DB_VM_SIZE=shared-cpu-1x
FLY_DB_VOLUME_SIZE=1

FLY_DB_PASSWORD=3eef4766d963b85863fe36535014be1fb708c810d3e06892
FLY_DB_HOST_NAME=gedw99-db-test.internal.
FLY_DB_HOST_URL=postgres://postgres:$(FLY_DB_PASSWORD)@$(FLY_DB_HOST_NAME)