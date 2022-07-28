BASE := golang
TAG := 1.18-alpine3.15

.PHONY: gen
gen:
	go run ./gen -os=linux -arch=amd64 -base=$(BASE) -tag=$(TAG) > Dockerfile.linux_amd64
	go run ./gen -os=linux -arch=arm64 -base=$(BASE) -tag=$(TAG) > Dockerfile.linux_arm64

IMAGENAME := golang-zig:go1.18-alpine3.15-zig
ZIGVER := 0.10.0-dev.2851+f639cb33a

build:
	sudo docker build \
		-f Dockerfile.linux_amd64 \
		-t $(IMAGENAME) . \
		--build-arg ZIGVER=$(ZIGVER)

publish:
	sudo docker push \
		$(IMAGENAME)

test-run:
	sudo docker run --rm -ti $(IMAGENAME) sh

clean:
#	sudo docker image rm $(IMAGENAME) || echo 1
	rm -rf Dockerfile*