.PHONY: build run clean

build:
	go build -o bin/ducker ./cmd/ducker

run: build
	sudo ./bin/ducker run $(ROOTFS) $(CMD)

clean:
	rm -rf bin
