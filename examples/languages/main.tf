
data "squidex_languages" "all" {}

# Returns all languages
output "all_languages" {
  value = data.squidex_languages.all.languages
}
