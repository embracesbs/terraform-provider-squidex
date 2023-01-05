# squidexclient

The squidexclient package is scaffolded using the [openapi-generator](https://github.com/OpenAPITools/openapi-generator).

- download the openapi 3.0 json file, go to squidex gui, api, general api, hit the download OpenApi button.
- [install the generator](https://github.com/OpenAPITools/openapi-generator#1---installation).
- Installation can also be done with npm: `npm install @openapitools/openapi-generator-cli -g`
- run the generator `openapi-generator-cli generate -g go -i ./squidex-5.3-swagger.json -o ./squidexclient`.
- fix scaffolded issues related to this [GitHub issue](https://github.com/OpenAPITools/openapi-generator/pull/2897).
  - Run a find & replace `package openapi` into `package squidexclient`
  - Remove the go.mod file from the squidexclient folder
  - rename all golang constants (enums) because they are duplicated. Just prefix them with the type name they belong to.
  - I had to add 1 type manually `FilterNodeOfIJsonValue`, because it was missing and implemented as type object instead.
  - Correct all types on the dto's by removing the `OneOf...` prefix of the type name.

TODO: describe model_field_properties_dto_full.go custom replacement

- copy `model_field_properties_dto_full.go` to squidexclient folder/package
- delete `model_field_properties_dto.go` from squidexclient folder/package
