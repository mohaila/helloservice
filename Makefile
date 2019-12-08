.PHONY: native all linux clean image

linux:
	@sh scripts/build-linux.sh

native:
	@sh scripts/build-native.sh

clean:
	@echo "Cleaning build"
	@rm -f build/linux/*	
	@rm -f build/native/*	

image: linux
	@sh scripts/build-image.sh

deploy: image
	@sh scripts/deploy-docker.sh

all: deploy	
