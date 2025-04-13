PROTO_FILES := $(shell find $(PROTO_PATH) -type f -name '*.proto')
PROTO_FILES := $(patsubst $(PROTO_PATH)/%,%,$(PROTO_FILES))
PROTO_MODULES := $(foreach file,$(PROTO_FILES),$(shell echo $(file)=$(PROTO_OUT_MODULE)/$(dir $(file)) | sed 's|/./||g; s|/$$||')) # :(

.PHONY: generate.go
generate.go:
	mkdir -p $(PROTO_OUT)
	rm -rf $(PROTO_OUT)/*
	protoc --proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUT) \
		--go-grpc_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(foreach module,$(PROTO_MODULES),--go_opt=M$(module)) \
		$(foreach module,$(PROTO_MODULES),--go-grpc_opt=M$(module)) \
		$(PROTO_FILES)


generate.py:
	mkdir -p $(PROTO_OUT)
	protoc --plugin=protoc-gen-mypy -I=. --mypy_out=$(PROTO_OUT) --python_out=$(PROTO_OUT) $(PROTO_FILES)
