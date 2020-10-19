terraform {
  required_providers {
    restapi = {
      source  = "terraform.embracecloud.nl/mastercard/restapi"
      version = "1.14.0"
    }
    squidex = {
      source  = "terraform.embracecloud.nl/embracecloud/squidex"
      version = "0.3.0"
    }
  }
}

variable "url" {
}
variable "app_name" {
}
variable "client_id" {
}
variable "client_secret" {
}

provider "squidex" {
  url           = var.url
  app_name      = var.app_name
  client_id     = var.client_id
  client_secret = var.client_secret
}

module "acme" {
  source = "./acme"
  languages = {
    "nl" = false
    "de" = false
  }
  clients = {
    "gateway"                        = "Owner"
    "content-provider-introspection" = "Reader"
  }
}

output "languages" {
  value = module.acme.all_languages
}

output "clients" {
  value = module.acme.all_clients
}
