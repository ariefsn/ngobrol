generate.client:
	rm -rf src/clients/${name} && \
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/src/clients/json/${name}.json \
		-g typescript-axios \
		-o /local/src/clients/${name}