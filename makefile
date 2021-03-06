
check:
ifeq ($(strip $(BK_CI_PIPELINE_NAME)),)
	@echo "\033[32m <====== 规范校验 =====> \033[0m"
	goimports -format-only -w -local github.com .
	gofmt -s -w .
	golangci-lint run
endif