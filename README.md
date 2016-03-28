# sous-singularity
Singularity API client, for use with Sous

## Components:

- *swaggering* is a Swagger 1.2 code generator library.
- *swagger-client-maker* is the CLI tool that uses the above.
`swagger-client-maker <source services.json> <target directory>`
will create DTOs and operation methods as indicated by the `service.json` and api JSON files.
- *client* is the HTTP client libraries used by the generated API.
- *client/jq-scripts* and *client/process_api.sh* are helpers to clean up the Sigularity JSON;
Singularity's generated JSON files are... not completely valid Swagger,
so they need a little massage before they're used.
