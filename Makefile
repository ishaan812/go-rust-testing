ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

clean:
	rm -rf ./lib/gorules/target
	rm ./lib/gorules.dll 
	rm ./lib/gorules/Cargo.lock ./lib/gorules.so go-rust

library:
	$(MAKE) -C lib/gorules build

build-windows:
	go env -w CGO_ENABLED="1"
	go env -w GOOS="linux"
	go env -w GOARCH="arm64" 
	go env -w CC="zig cc -target aarch64-linux-musl"
	go env -w CXX="zig c++ -target aarch64-linux-musl"
	go build -ldflags="-r ./lib -linkmode external" -o arm64/bootstrap

build-linux:
	go env -w CGO_ENABLED="1"  
	go env -w GOOS="linux"
	go env -w GOARCH="amd64"
	go env -w CC="zig cc -target x86_64-linux-musl"
	go env -w CXX="zig c++ -target x86_64-linux-musl"
	go build -ldflags="-r './lib' -linkmode external" -o amd64/bootstrap

all: library build

run: build-linux
	./bootstrap
