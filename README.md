# crosswordgame-go-config

Config for deploying [crosswordgame-go](https://github.com/mcoot/crosswordgame-go).

Deploys in docker-compose with Envoy in front.

The docker-compose.yaml and envoy.yaml are templated out based on one of the configs 
in `configs` – with or without TLS. You will want to change the dns/email.

If using TLS, you should first renew the cert with:

```shell
export CONFIG_FILE="./configs/config-https.yaml"
make renew-cert
```

And then templating will be run and the app started with:

```shell
export CONFIG_FILE="./configs/config-https.yaml" # or http if not using TLS
make run
```

It can later be stopped with

```shell
make stop
```