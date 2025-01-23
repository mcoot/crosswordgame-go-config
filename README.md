# crosswordgame-go-config

Config for deploying [crosswordgame-go](https://github.com/mcoot/crosswordgame-go).

Deploys in docker-compose with Envoy in front.

The docker-compose.yaml and envoy.yaml are templated out based on one of the configs 
in `configs` – with or without TLS.

The templating can be done with:

```shell
CONFIG_FILE="./configs/config-http.yaml" make template
```

And then the app can be started with:

```shell
cd out
docker compose start
```