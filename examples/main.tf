terraform {
  required_providers {
    squidex = {
      source = "terraform.embracecloud.nl/embracecloud/squidex"
    }
  }
}

provider "squidex" {
  client_id      = var.client_id
  client_secret  = var.client_secret
  token_endpoint = "${var.host}/identity-server/connect/token"
  url            = "${var.host}/api"
}

resource "squidex_app" "test" {
  name = "squidex-provider-test"
  description = "description1"
}

resource "squidex_client" "test" {
  app_name = squidex_app.test.name
  name = "squidex-provider-test"
  role = "squidex"
}

resource "squidex_languages" "test" {
  app_name = squidex_app.test.name
  language {
    language = "en-NL"
    is_master = true
  }
}
