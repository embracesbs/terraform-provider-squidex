terraform {
  required_providers {
    squidex = {
      source = "terraform.embracecloud.nl/embracecloud/squidex"
    
    }
  }
}


provider "squidex" {
  client_id = "root"
  client_secret = "4HUt5ONG1k0OknPA"
  token_endpoint = "https://squidex-embracecloudte.features.embracecloud.io/identity-server/connect/token"
  url = "https://squidex-embracecloudte.features.embracecloud.io/api"
  
}
resource "squidex_app" "test" {
  name = "test-sven-test"
  description = "description1"
}
resource "squidex_client" "test" {
  app_name = squidex_app.test.name
  name = "testsven"
  role = "squidex"
}
resource "squidex_languages" "test" {
  app_name = squidex_app.test.name



        language {
      language = "en"
      is_master = true
    }


  
}





