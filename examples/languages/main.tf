
data "squidex_app_languages" "all" {}

# Returns all languages
output "all_languages" {
  value = data.squidex_app_languages.all.languages
}
