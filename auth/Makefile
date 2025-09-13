# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/proto

# Путь до сгенеренных .pb.go файлов
GEN_PROTO_PATH := $(CURDIR)/api/gen

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# генерация .go файлов с помощью protoc
.protoc-generate:
	mkdir -p $(GEN_PROTO_PATH)
	$(PROTOC) --proto_path=$(CURDIR) \
	--go_out=$(GEN_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(GEN_PROTO_PATH) --go-grpc_opt paths=source_relative \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .protoc-generate .tidy

# Билд приложения
build:
	go build -o $(LOCAL_BIN) cmd/client
	go build -o $(LOCAL_BIN) cmd/server 
	
# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	generate \
	build