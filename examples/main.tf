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
  source = "./languages"
}

output "languages" {
  value = module.acme.all_languages
}
