# go-singularity

Singularity API client, originally for use with Sous. Perhaps useful to others.

## Components:

- `jq-scripts/` and `process_api.sh` are helpers to clean up the Sigularity JSON;
Singularity's generated JSON files are... not completely valid Swagger,
so they need a little massage before they're used.

## Development:

The vast majority of this code is generated using the [swagger-client-maker.](http://github.com/opentable/swaggering)
To install the tool:

```bash
go get github.com/opentable/swaggering/cmd/swagger-client-maker
```
