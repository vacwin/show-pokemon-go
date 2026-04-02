BINARY_NAME := pokemon-colorscripts-go

PREFIX      := $(HOME)/.local
BIN_DIR     := $(PREFIX)/bin
DATA_DIR    := $(PREFIX)/share/$(BINARY_NAME)

BUILD_DIR   := build

LDFLAGS     := -ldflags "-X main.PROGRAM_DIR=$(DATA_DIR)"

GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)

.PHONY: build install uninstall clean

build:
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "build passed: $(BUILD_DIR)/$(BINARY_NAME)"

install: build
	@mkdir -p $(BIN_DIR)
	@mkdir -p $(DATA_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/$(BINARY_NAME)
	cp -r colorscripts $(DATA_DIR)/
	cp pokemon.json $(DATA_DIR)/
	@echo ""
	@echo "done:"
	@echo "  binary: $(BIN_DIR)/$(BINARY_NAME)"
	@echo "  data:   $(DATA_DIR)/"
	@echo ""
	@case ":$$PATH:" in \
		*":$(BIN_DIR):"*) ;; \
		*) echo "warning: $(BIN_DIR) is not in PATH. Add to your .bashrc / .zshrc:"; \
		   echo "   export PATH=\"$(BIN_DIR):\$$PATH\""; \
	esac

uninstall:
	rm -f $(BIN_DIR)/$(BINARY_NAME)
	rm -rf $(DATA_DIR)
	@echo "deleted."

clean:
	rm -rf $(BUILD_DIR)
	go clean
