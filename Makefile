workdir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


.PHONY: format
format: ## 格式化 API 文件
	@echo "格式化 API 文件"
	@go tool goctl api format --dir=${workdir}/app/api

.PHONY: api
api: ## 自动生成 API 代码
	@echo "自动生成 API 代码"
	@/bin/sh -c "echo 'Go version:' && go version && echo 'goctl version:' && go tool goctl --version && \
				tool goctl api go --api=${workdir}/
