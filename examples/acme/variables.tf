variable "languages" {
  type        = map
  description = "iso2Code = isMaster"
  default = {
    "en" = true
    "nl" = false
  }
}

variable "clients" {
  type        = map
  description = "name = role"
  default = {
    "gateway"                        = "Owner"
    "content-provider-introspection" = "Reader"
  }
}
