terraform {
  required_providers {
    squidex = {
      source = "terraform.embracecloud.nl/embracecloud/squidex"
      version = "0.7.1"
    }
  }

  required_version = "~> 1.1.1"
}
provider "squidex" {
  client_id           = var.client_id
  client_secret       = var.client_secret
  token_endpoint      = "${var.host}/identity-server/connect/token"
  url                 = "${var.host}/api"
  schema_delete_allow = true
}

# TODO: discuss strategy, do we allow for destroy and create on app resources?
resource "squidex_app" "test" {
  name        = "squidex-provider-test9"
  description = "description1"
}

resource "squidex_app" "test2" {
  name        = "squidex-provider-test2"
  description = "description1"
}

resource "squidex_client" "test" {
  app_name = squidex_app.test.name
  name     = "squidex-provider-test"
  role     = squidex_role.embracecloud_admin.name
}

resource "squidex_languages" "test" {
  app_name = squidex_app.test.name
  language {
    language  = "en-US"
    is_master = true
  }
  language {
    language  = "nl-NL"
    is_master = false
  }
}

resource "squidex_role" "embracecloud_admin" {
  app_name = squidex_app.test.name
  name     = "embracecloud-admin"
  permissions = [
    "contents.*",
    "schemas.read",
    "roles",
    "contributors",
    "assets",
    "languages",
    "patterns",
    "workflows",
    "search",
    "backups",
    "comments",
    "clients",
    "history",
    "ping",
    "plans",
    "rules",
    "translate",
    "usage"
  ]
  properties = {}
}
resource "squidex_role" "app_admin" {
  app_name = squidex_app.test.name
  name     = "app-admin"
  permissions = [
    "contents.*",
    "schemas.read",
    "roles.read",
    "contributors",
    "assets",
    "languages.read",
    "patterns.read",
    "workflows.read",
    "search",
    "backups",
    "comments",
    "history",
    "rules",
    "translate",
    "squidex.apps.*"
  ]
  //  properties  = {"ui.api.hide" = true}
  properties = {}
}

resource "squidex_contributor" "embracecloud_contributors" {
  for_each          = toset(["admin@embracecloud.nl", "test1@embracecloud.nl", "test2@embracecloud.nl"])
  app_name          = squidex_app.test.name
  contributor_email = each.value
  role              = squidex_role.embracecloud_admin.name
}


resource "squidex_contributor" "app_admin_contributors" {
  for_each          = toset(["admin@embracecloud.nl", "test1@embracecloud.nl", "test2@embracecloud.nl"])
  app_name          = squidex_app.test.name
  contributor_email = each.value
  role              = squidex_role.app_admin.name
}

 resource "squidex_contributor" "admin_owner" {
   app_name          = squidex_app.test.name
   contributor_email = "admin@embracecloud.nl"
   role              = "Owner"
 }

# TODO: discuss strategy, do we allow for destroy and create on schema resources?
resource "squidex_schema" "test" {
  schema_field_delete_allow   = true
  schema_field_recreate_allow = true
  depends_on                  = [squidex_languages.test]
  app_name                    = squidex_app.test.name
  name                        = "blog1"
  published                   = true
  singleton                   = false
  properties {
    label                = "bloglabel"
    content_sidebar_url  = "someuri_1"
    contents_sidebar_url = "someuri_2"
    hints                = "This is a hint"
    tags                 = ["tag1", "tag2"]
  }
  fields_in_list = [
    "meta.created",
    "meta.lastModified",
    "meta.version"
  ]
  fields_in_references = []
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
    name           = "self-reference-field"
    partitioning   = "invariant"
    self_reference = true
    hidden         = false # 
    locked         = false # should not be used
    disabled       = false # should not be used
    properties {
      field_type = "References"
      editor     = "List"
    }
  }

  fields {
    name         = "reference-1"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type = "References"
      editor     = "Tags"
    }
  }

  fields {
    name         = "referenceCAPS-2"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type = "References"
      editor     = "Tags"
      default_values = {
        "nl-NL" = "string-1"
        "en-US" = "string-3"
      }
    }
  }

  fields {
    name         = "number-1"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type = "Number"
      editor     = "Input"
    }
  }

  fields {
    name         = "bool-1"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type = "Boolean"
      editor     = "Checkbox"
      default_values = {
        "nl-NL" = "True"
        "en-US" = "false"
      }
    }
  }

  fields {
    name         = "string-3"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type = "String"
      editor     = "Input"
    }
  }

  fields {
    name         = "string-2"
    partitioning = "language"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used
    properties {
      field_type    = "String"
      editor        = "Input"
      default_value = ["fake"]
      default_values = {
        "nl-NL" = "defaultvalue-nl-2"
      }
    }
  }

  fields {
    name         = "string-1"
    partitioning = "invariant"
    hidden       = false # 
    locked       = false # should not be used
    disabled     = false # should not be used

    properties {
      field_type    = "String"
      editor        = "Input"
      label         = "author-label"
      hints         = "hints2"
      placeholder   = "placeholder"
      required      = false
      half_width    = false
      editor_url    = "editor_url"
      unique        = true
      tags          = ["tag1", "tag2"]
      default_value = ["defaultvalue"]
    }
    // Only properties.field_type = "Array" effect nested block (ignored for other types)
    nested {
      name     = "nestedfield"
      hidden   = false
      locked   = false # should not be used
      disabled = false # should not be used

      properties {
        field_type  = "String"
        editor      = "Input"
        label       = "label"
        hints       = "hints"
        placeholder = "placeholder"
        required    = false
        half_width  = false
        editor_url  = "editor_url"
        unique      = true
        tags        = ["tag1", "tag2"]
      }
    }
  }

  fields {
    name         = "datetime-1"
    partitioning = "invariant"
    properties {
      field_type = "DateTime"
      editor     = "DateTime"
      min_value  = "1900-01-01T00:00:00Z"
      max_value  = "2021-02-01T23:59:59Z"
    }
  }

  fields {
    name = "bug"
    properties {
      field_type = "DateTime"
      editor     = "DateTime"
      min_value  = "1900-01-01T00:00:00Z"
      max_value  = "2021-02-01T23:59:59Z"
    }
  }
}
