terraform {
  required_providers {
    squidex = {
      source  = "terraform.embracecloud.nl/embracecloud/squidex"
      version = "0.3.0"
    }
  }
}

data "squidex_languages" "all" {}
data "squidex_clients" "all" {}

resource "squidex_language" "languages" {
  for_each = var.languages

  iso_2_code = each.key
  is_master  = each.value
}

resource "squidex_client" "clients" {
  for_each = var.clients

  name = each.key
  role = each.value
}

# resource "squidex_language" "nl" {
#   iso_2_code = "nl"
#   is_master  = true
# }

# resource "squidex_language" "de" {
#   iso_2_code = "de"
#   is_master  = false
# }

# resource "squidex_client" "gateway" {
#   name = "gateway"
#   role = "Owner"
# }

# resource "squidex_client" "content-provider-introspection" {
#   name = "content-provider-introspection"
#   role = "Reader"
# }

# Returns all languages
output "all_languages" {
  value = data.squidex_languages.all.languages
}

# Returns all clients
output "all_clients" {
  value = data.squidex_clients.all.clients
}
