// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//1aD/GCPGgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
