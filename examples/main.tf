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

resource "squidex_language" "en" {
  iso_2_code = "en"
  is_master = true
}

resource "squidex_language" "de" {
  iso_2_code = "de"
  is_master = false
}

module "acme" {
  source = "./languages"
}

output "languages" {
  value = module.acme.all_languages
}
