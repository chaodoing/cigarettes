BINARY_NAME=cigarettes
DATE=`date +%FT%T%z`
VERSION=1.0.2
ENV=release

build:
	@echo "build apple-darwin version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0  GOOS=darwin go build -o release/apple/darwin/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./

windows:
	@echo "build windows-i386 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc" go build -o release/windows/i386/bin/$(BINARY_NAME).exe -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/windows/i386/

	@echo "build windows-amd64 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" go build -o release/windows/amd64/bin/$(BINARY_NAME).exe -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/windows/amd64/
linux:
	@echo "build linux-amd64 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  release/linux/amd64/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/linux/amd64/
	@echo "build linux-i386 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o  release/linux/i386/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/linux/i386/
all:
	@echo "build apple-darwin version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -o release/apple/darwin/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/apple/darwin/

	@echo "build linux-amd64 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  release/linux/amd64/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/linux/amd64/

	@echo "build linux-i386 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o  release/linux/i386/bin/$(BINARY_NAME) -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/linux/i386/

	@echo "build windows-i386 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc" go build -o release/windows/i386/bin/$(BINARY_NAME).exe -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/windows/i386/

	@echo "build windows-amd64 version: ${VERSION} date: ${DATE}"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" go build -o release/windows/amd64/bin/$(BINARY_NAME).exe -ldflags '-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.NAME=$(BINARY_NAME)' ./
	# @cp -r conf release/windows/amd64/

clean:
	@rm -rf release
	@rm -rf $(BINARY_NAME)
install:
	@cp release/apple/darwin/$(BINARY_NAME) $(GOPATH)/bin/