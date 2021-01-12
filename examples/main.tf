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
  name = "squidex-provider-test5"
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
    language = "nl-NL"
    is_master = true
  }
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

resource "squidex_contributor" "test" {
  app_name          = squidex_app.test.name
  contributor_email = "michiel@qvision.nl"
  role              = squidex_role.test.name
  invite            = false
}

resource "squidex_contributor" "test_michiel_owner" {
  app_name          = squidex_app.test.name
  contributor_email = "michiel.vanklinken@embracecloud.nl"
  role              = "Owner"
  invite            = false
}

# TODO: discuss strategy, do we allow for destroy and create on schema resources?
resource "squidex_schema" "test" {
  app_name  = squidex_app.test.name
  name      = "blog1"
  published = true
  singleton = false
  properties {
    label                = "bloglabel"
    content_sidebar_url  = "someuri_1"
    contents_sidebar_url = "someuri_2"
    hints                = "This is a hint"
    tags                 = [ "tag1", "tag2" ]
  }
  fields_in_list = [ 
    "author",
    "title",
    "meta.created",
    "meta.lastModified",
    "meta.version" 
    ]
  fields_in_references = [ "author", "title" ]
  preview_urls = {
    "somepagename" = "https://fqdn/page"
  }
  category = "somecategory2"
  scripts {
    query  = "value"
    update = "value"
    create = "value"
    delete = "value"
    change = "value"
  }

  fields {
    name         = "author"
    partitioning = "invariant"
		hidden       = false # 
		locked       = false # should not be used
		disabled     = false # should not be used
		
    properties {
      field_type  = "Array"
      label       = "author-label"
      hints       = "hints"
      placeholder = "placeholder"
      required    = false
      half_width  = false
      editor_url  = "editor_url"
      tags        = ["tag1", "tag2"]
    }
    // Only properties.field_type = "Array" effect nested block (ignored for other types)
    nested {
      name         = "nestedfield"    
		  hidden       = false
		  locked       = false # should not be used
		  disabled     = false # should not be used
    
      properties {
        field_type  = "String"
        label       = "label"
        hints       = "hints"
        placeholder = "placeholder"
        required    = false
        half_width  = false
        editor_url  = "editor_url"
        tags        = ["tag1", "tag2"]
      }
    }
  }

  fields {
    name         = "title"
    partitioning = "invariant"
    properties {
      field_type = "String"
    }
  }

  fields {
    name         = "title-updated"
    partitioning = "invariant"
    properties {
      field_type = "String"
    }
  }

}