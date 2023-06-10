# Specify the name of your binary
BINARY_NAME := himitsud

# Specify the directories where your source code files reside
SRC_DIR := ./src
BUILD_DIR := ./build

# Specify the compiler and flags
GO := go
GOFLAGS :=
LDFLAGS :=

# Define the build target
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Define the install target
install:
	@echo "Installing $(BINARY_NAME)..."
	@$(GO) install $(GOFLAGS) $(SRC_DIR)

# Define the clean target
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Define the default target
default: build

# Define phony targets (targets that are not file names)
.PHONY: build install clean default
