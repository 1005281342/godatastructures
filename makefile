
check:
	@echo "\033[32m <====== 代码规范检查 =====> \033[0m"
	goimports -format-only -w -local github.com .
	gofmt -s -w .
	golangci-lint run
