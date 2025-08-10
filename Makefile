workdir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


.PHONY: format
format: ## 格式化 API 文件
	@echo "格式化 API 文件"
	@goctl api format --dir=${workdir}/api

.PHONY: api
api: ## 自动生成 API 代码
	@echo "自动生成 API 代码"
	@/bin/sh -c "\
		echo 'Go version:' && go version && echo 'goctl version:' && goctl --version && \
		goctl api go --home ${workdir}/response --api ${workdir}/api/app.api --style=go_zero --dir ${workdir}/atlas/ && \
		goctl rpc protoc ${workdir}/pb/auth.proto --style=go_zero --proto_path=${workdir}/pb/ \
			--go_out=${workdir}/auth --go-grpc_out=${workdir}/auth --zrpc_out=${workdir}/auth && \
		goctl rpc protoc ${workdir}/pb/user.proto --style=go_zero --proto_path=${workdir}/pb/ \
			--go_out=${workdir}/user --go-grpc_out=${workdir}/user --zrpc_out=${workdir}/user \
		"
	@echo "生成swagger文件并运行"
	@/bin/sh -c "goctl api swagger --api ./api/app.api --dir atlas/internal/handler/swagger --filename app"

.PHONY: swagger
swagger:  ## 生成swagger文件并更新swagger配置文件
	@echo "生成swagger文件并运行"
	@/bin/sh -c "goctl api swagger --api ./api/app.api --dir atlas/internal/handler/swagger --filename app"


.PHONY: atlas
atlas:  ## 运行网关（*）
	@echo "starting atlas..."
	@go run ${workdir}/atlas/atlas.go -f ${workdir}/atlas/etc/atlas-api.yaml