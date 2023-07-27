envfile ?= .env

-include $(envfile)

define export_envfile
    export $(shell sed 's/=.*//' $(1))
endef

ifneq ("$(wildcard $(envfile))", "")
    $(eval $(call export_envfile,$(envfile)))
endif

.PHONY: init
init:
	@cp .env.dist .env


.PHONY: server-start
server-start:
	@go run cmd/server/main.go