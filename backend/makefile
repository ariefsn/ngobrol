generate.gql:
	go run github.com/99designs/gqlgen generate

generate.client:
	rm -rf clients/${name} && \
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/clients/json/${name}.json \
		-g go \
		-o /local/clients/${name} \
	  -p generateInterfaces=true \
		--package-name ${name} \
		--git-user-id ariefsn \
		--git-repo-id terkirim/openapi/${name}

migrate.install:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest && go install -tags 'mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	
# && go install github.com/golang-migrate/migrate/database/mongodb@latest

migrate.create:
	migrate create -ext json -dir ./migrations/mongo -seq ${name}
