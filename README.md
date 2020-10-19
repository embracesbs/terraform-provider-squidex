# Azure DevOps

Connected with Azure Pipelines, produces Artifact - an Universal Package.

# Terraform Provider Squidex

Run the following command to build the provider
```shell
go build -o terraform-provider-squidex
```

Alternatively, you can run:
```shell
make build
```

## Test sample configuration

First, build and install locally the provider.

```shell
make local-install
```

Then, create a file `terraform.tfvars` in `./example` directory.  
Configure values for your test Squidex application.  
Run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init && terraform apply
```