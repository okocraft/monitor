package oapi_codegen

//go:generate go tool oapi-codegen --config=api.yml ../../../schema/openapi/monitor-api.yml
//go:generate go tool oapi-codegen --config=embeded-spec.yml ../../../schema/openapi/monitor-api.yml
