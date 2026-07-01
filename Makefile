PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg

DEMO_BIN :=  $(PWD)/demo/bin

GEN_SML_DIR := $(GEN_DIR)/mxsml
GEN_EXTENSION_DIR := $(GEN_DIR)/mxsmlextension

.PHONY: all
all: demo

.PHONY: demo
demo:
	go build -o $(DEMO_BIN)/mxsmlDemo     ./demo/mxsml
	go build -o $(DEMO_BIN)/mxsmlExtDemo  ./demo/mxsmlextension
	go build -o $(DEMO_BIN)/mxsmlGetEidDemo  ./demo/mxsmlgeteid

.PHONY: gen
gen:
	c-for-go -out $(PKG_DIR) $(GEN_SML_DIR)/MxSml.yaml
	cp $(GEN_SML_DIR)/MxSml.h $(PKG_DIR)/mxsml/
	cp $(GEN_SML_DIR)/MxSmlMcm.h $(PKG_DIR)/mxsml/
	cd $(PKG_DIR)/mxsml; \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
		rm -rf _obj cgo_helpers.go types.go; \

	sed -i -e 's|mxsml/MxSml\.h:[0-9]\+|MxSml\.h|g' \
	 $(PKG_DIR)/mxsml/mxsml.go
	sed -i -e 's|mxsml/MxSmlMcm\.h:[0-9]\+|MxSmlMcm\.h|g' \
	 $(PKG_DIR)/mxsml/mxsml.go
	sed -i -e 's|mxsml/MxSml\.h:[0-9]\+|MxSml\.h|g' \
	 $(PKG_DIR)/mxsml/const.go

.PHONY: gen-extension
gen-extension:
	cp $(GEN_EXTENSION_DIR)/MxSmlExtension.h  $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp
	sed -i 's/sml/Sml/g'  $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp
	sed -i 's/mxSmlExSystemGetMxSmlVersion/mxSmlExSystemGetMxsmlVersion/g' $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp
	sed -i 's/typedef struct \([A-Za-z0-9_]*\)\* \([A-Za-z0-9_]*\);/typedef struct {\n    struct \1* handle;\n} \2;/g' \
	 $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp

	c-for-go -out $(PKG_DIR) $(GEN_EXTENSION_DIR)/MxSmlExtension.yaml
	cp $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp $(PKG_DIR)/mxsmlextension/MxSmlExtension.h

	cd $(PKG_DIR)/mxsmlextension; \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
		rm -rf _obj cgo_helpers.go types.go; \

	sed -i -e 's|mxsmlextension/MxSmlExtension\.h\.tmp:[0-9]\+|MxSmlExtension\.h|g' \
	 $(PKG_DIR)/mxsmlextension/mxsmlextension.go
	sed -i -e 's|mxsmlextension/MxSmlExtension\.h\.tmp:[0-9]\+|MxSmlExtension\.h|g' \
	 $(PKG_DIR)/mxsmlextension/const.go

	rm $(GEN_EXTENSION_DIR)/MxSmlExtension.h.tmp

clean:
	-rm $(DEMO_BIN) -rf
