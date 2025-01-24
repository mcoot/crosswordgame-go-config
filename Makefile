CONFIG_FILE ?= ./configs/config-http.yaml

.PHONY: dirs
dirs:
	mkdir -p ./out/letsencrypt/conf && mkdir -p ./out/letsencrypt/www

.PHONY: template
template: dirs
	CONFIG_FILE=$(CONFIG_FILE) go run ./cmd/template/main.go

.PHONY: renew-cert
renew-cert: dirs template
	cd ./out && docker compose -f docker-compose-certbot.yml up

.PHONY: run
run: dirs template
	cd ./out && docker compose up -d