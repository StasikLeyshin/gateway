goCmd ?= go

DIR_GRPC_GATEWAY=internal/api/grpc-gateway/pb
DIR_GRPC=internal/api/grpc/pb

ifeq ($(shell echo "check_quotes"),"check_quotes")
   WINDOWS := yes
else
   WINDOWS := no
endif

ifeq ($(WINDOWS),yes)
   mkdir = mkdir $(subst /,\,$(1)) > nul 2>&1 || (exit 0)
   echo = echo $(1)
else
   mkdir = mkdir -p $(1)
   echo = echo $(1)
endif

all: gen_grpc_gateway_proto gen_grpc_proto

gen_grpc_gateway_proto: create_directory_grpc_gateway
	protoc --proto_path=api/grpc-gateway-proto \
	--go_out=internal/api/grpc-gateway/pb \
	--go-grpc_out=internal/api/grpc-gateway/pb \
	--grpc-gateway_out=internal/api/grpc-gateway/pb \
	--openapiv2_out=doc/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=base_grpc_gateway \
	api/grpc-gateway-proto/services/*.proto

gen_grpc_proto: create_directory_grpc
	protoc --proto_path=api/grpc-proto \
	--go_out=internal/api/grpc/pb \
	--go-grpc_out=internal/api/grpc/pb \
	--grpc-gateway_out=internal/api/grpc/pb \
	api/grpc-proto/services/*.proto

create_directory_grpc_gateway:
	$(call mkdir,$(DIR_GRPC_GATEWAY))

create_directory_grpc:
	$(call mkdir,$(DIR_GRPC))