# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/proto

# Путь до сгенеренных .pb.go файлов
GEN_PROTO_PATH := $(CURDIR)/api/gen

GOOGLE_API_DIR=${PROTO_PATH}/google/api

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@febffdd
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# # генерация .go файлов с помощью protoc
.protoc-generate:
	mkdir -p $(GEN_PROTO_PATH)
	$(PROTOC) --proto_path=$(CURDIR) -I$(GOOGLE_API_DIR) \
	--go_out=$(GEN_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(GEN_PROTO_PATH) --go-grpc_opt paths=source_relative \
	--grpc-gateway_out=$(GEN_PROTO_PATH) --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=$(GEN_PROTO_PATH) \
	$(PROTO_PATH)/auth_service.proto \
	$(PROTO_PATH)/auth_messages.proto

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .protoc-generate .tidy

# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/client
	go build -o $(LOCAL_BIN) ./cmd/server 
	
# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	generate \
	build