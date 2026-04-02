.PHONY: all build-linux build-mac build-windows clean

APP_NAME := gohttpserver
GO_FILES := assets.go httpstaticserver.go ipa.go main.go oauth2-proxy.go openid-login.go res.go utils.go utils_test.go zip.go zip_test.go
compilerFlag := -gcflags=-trimpath=$(shell pwd) -asmflags=-trimpath=$(shell pwd)

all: build-linux build-mac build-windows

build-linux: $(GO_FILES)
	GOOS=linux GOARCH=amd64 go build $(compilerFlag) -o build/linux/$(APP_NAME) $(GO_FILES)

build-mac: $(GO_FILES)
	GOOS=darwin GOARCH=arm64 go build $(compilerFlag) -o build/mac/$(APP_NAME) $(GO_FILES)
	GOOS=darwin GOARCH=amd64 go build $(compilerFlag) -o build/mac-x86/$(APP_NAME) $(GO_FILES)


build-windows: $(GO_FILES)
	GOOS=windows GOARCH=amd64 go build $(compilerFlag) -o build/windows/$(APP_NAME).exe $(GO_FILES)

clean:
	rm -Rf build
