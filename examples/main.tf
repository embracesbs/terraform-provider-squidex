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

# TODO: discuss strategy, do we allow for destroy and create on app resources?
resource "squidex_app" "test" {
  name = "squidex-provider-test2"
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

# TODO: discuss strategy, do we allow for destroy and create on schema resources?
resource "squidex_schema" "test" {
  app_name  = squidex_app.test.name
  name      = "squidex-provider-test-updated"
  published = true
  singleton = false
}

resource "squidex_role" "test" {
  app_name    = squidex_app.test.name
  name        = "squidex-provider-test"
  permissions = [
    "*",
    "contents.*",
    "schemas.read",
  ]
  properties  = {}
}

resource "squidex_role" "test3" {
  app_name    = squidex_app.test.name
  name        = "squidex-provider-test3"
  permissions = [
    "contents.*"
  ]
  properties  = {}
}

resource "squidex_contributor" "test" {
  app_name          = squidex_app.test.name
  contributor_email = "michiel.vanklinken@embracecloud.nl"
  role              = squidex_role.test.name
  invite            = false
}
