.PHONY: build

.build:
	go build -o builds/torque-hass-proxy-${GOOS}-${GOARCH}${GOARM}

build-linux-arm:
	GOOS=linux GOARCH=arm GOARM=7 $(MAKE) .build

build-linux-arm6:
	GOOS=linux GOARCH=arm GOARM=6 $(MAKE) .build

build-android-arm64:
	GOOS=android GOARCH=arm64 $(MAKE) .build

build-mac-amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) .build

build-all:
	$(MAKE) build-linux-arm
	$(MAKE) build-linux-arm6
	$(MAKE) build-android-arm64
	$(MAKE) build-mac-amd64
