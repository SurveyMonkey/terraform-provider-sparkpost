# terraform-provider-sparkpost

A terraform provider for Sparkpost

## Local Development

Run the following command to build the provider

```shell
make build
```

## Test Example Configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd example/template
terraform init
SPARKPOST_API_KEY=<api_key> terraform apply
```
