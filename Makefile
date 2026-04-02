# Имя бинарника
BINARY_NAME := pokemon-colorscripts-go

# Куда ставим данные и бинарник (per-user, без sudo)
PREFIX      := $(HOME)/.local
BIN_DIR     := $(PREFIX)/bin
DATA_DIR    := $(PREFIX)/share/$(BINARY_NAME)

# Папка для сборки
BUILD_DIR   := build

# Вшиваем путь к данным в бинарник на этапе компиляции
LDFLAGS     := -ldflags "-X main.PROGRAM_DIR=$(DATA_DIR)"

# Определяем ОС и архитектуру для кросс-компиляции
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)

.PHONY: build install uninstall clean

## build: собирает бинарник в build/
build:
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Собрано: $(BUILD_DIR)/$(BINARY_NAME)"

## install: ставит бинарник в ~/.local/bin/, данные в ~/.local/share/
install: build
	@mkdir -p $(BIN_DIR)
	@mkdir -p $(DATA_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/$(BINARY_NAME)
	cp -r colorscripts $(DATA_DIR)/
	cp pokemon.json $(DATA_DIR)/
	@echo ""
	@echo "Установлено:"
	@echo "  бинарник: $(BIN_DIR)/$(BINARY_NAME)"
	@echo "  данные:   $(DATA_DIR)/"
	@echo ""
	@# Проверяем, есть ли ~/.local/bin в PATH
	@case ":$$PATH:" in \
		*":$(BIN_DIR):"*) ;; \
		*) echo "⚠  $(BIN_DIR) не в PATH. Добавь в свой .bashrc / .zshrc:"; \
		   echo "   export PATH=\"$(BIN_DIR):\$$PATH\""; \
		   echo "" ;; \
	esac

## uninstall: удаляет бинарник и данные
uninstall:
	rm -f $(BIN_DIR)/$(BINARY_NAME)
	rm -rf $(DATA_DIR)
	@echo "Удалено."

## clean: чистит папку сборки
clean:
	rm -rf $(BUILD_DIR)
