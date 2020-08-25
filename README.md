# Azure DevOps

Connected with Azure Pipelines, produces Artifact.

# Terraform Provider Squidex

Run the following command to build the provider

```shell
go build -o terraform-provider-squidex
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init && terraform apply
```