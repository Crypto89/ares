.PHONY: build-swagger
build-swagger:
	@rm -rf api/{client,models,restapi}
	@swagger generate server -f api/openapi.yaml --exclude-main -A Ares --target api