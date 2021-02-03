# Azure DevOps

Connected with Azure Pipelines, produces Artifact - an Universal Package.

## Terraform Provider Squidex

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
# on windows: make local-install-win
```

This will create the provider binary.
Copy that binary into the location that Terraform will try to find it.
Depending on the OS, version and Terraform version this could differ. On Windows 64 with Terraformn 0.13: /terraform.d/plugins/terraform.embracecloud.nl/embracecloud/squidex/0.4.0/windows_amd64"

Then, create a file `terraform.tfvars` in `./examples` directory.  
Configure values for your test Squidex application/host.

Either connect to a remote existing squidex, or run a local docker;

```bash
docker-compose --env-file ./docker-compose.env up -d
# docker-compose down
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init -backend=false
terraform apply -auto-approve
```

## Release

Use a git tag with semver (e.g. v0.4.2). This will trigger a pipeline to release the provider as a package.
Currently the package is zip containing linux 64 & windows 64 binaries.

- Trigger & pipeline: Azure Pipelines
- Package: Azure Artifacts (Universal)
