ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

clean:
	rm -rf ./lib/gorules/target
	rm ./lib/gorules.dll 
	rm ./lib/gorules/Cargo.lock ./lib/gorules.so go-rust

library:
	$(MAKE) -C lib/gorules build

build-windows:
	go env -w CGO_ENABLED="1"
	go env -w GOOS="windows"
	go env -w GOARCH="amd64" 
	go env -w CC="zig cc -target x86_64-windows-gnu"
	go env -w CXX="zig c++ -target x86_64-windows-gnu"
	go build -ldflags="-r ./lib -linkmode external" -o arm64/bootstrap

build-linux:
	go env -w CGO_ENABLED="1"  
	go env -w GOOS="linux"
	go env -w GOARCH="amd64"
	go env -w CC="zig cc -target x86_64-linux-musl"
	go env -w CXX="zig c++ -target x86_64-linux-musl"
	go build -ldflags="-r './lib' -linkmode external" -o bootstrap

all: library build

build: 
	go build -ldflags="-r ./lib" -o go-rust.exe

run: build
	./go-rust
