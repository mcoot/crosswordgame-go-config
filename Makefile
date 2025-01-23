CONFIG_FILE ?= ./configs/config-http.yaml

.phony: template
make template:
	CONFIG_FILE=$(CONFIG_FILE) go run ./cmd/template/main.go