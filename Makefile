# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAMES=go-monkey go-min go-rcf

all: test build

build:
	@for NAME in $(BINARY_NAMES); do \
		cd cmd/$${NAME} && $(GOBUILD) -o ../../$${NAME} -v; \
		cd ../../; \
	done
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
fmt:
	for go_file in `find . -name \*.go`; do \
		go fmt $${go_file}; \
	done
