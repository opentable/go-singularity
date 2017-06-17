skel: 
	mkdir -p v$(VERSION)/dtos
	cp client.go v$(VERSION)
	cp process_apis.sh v$(VERSION)
	cp swagger-scripts/swagger.sh v$(VERSION)

swagger-install:
	go get github.com/opentable/swaggering/cmd/swagger-client-maker
