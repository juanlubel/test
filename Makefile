

dev:
	docker-compose -f ./compose/docker-compose.yml up -d
	air -c .air.toml

registry:
	docker build -t registry.spu-labs.dev/saludos .
	docker push registry.spu-labs.dev/saludos

compose:
	scp ./docker-compose.yml labs:/home/ubuntu/saludos
	ssh labs 'docker-compose -f /home/ubuntu/saludos/docker-compose.yml pull'
	ssh labs 'docker-compose -f /home/ubuntu/saludos/docker-compose.yml up -d'

templ:
	templ generate

deploy: registry compose
