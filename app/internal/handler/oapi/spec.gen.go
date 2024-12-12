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

	"H4sIAAAAAAAC/+xWTW/cNhD9KwRboBd5tWnSFtDNTZHAgBcIkho9GAuDlmYldiUOMxw5NQz992Io7SeV",
	"pjAa9JLTrsjHN/PeDD+edImdRweOgy6edCgb6Ez8e1mWEMLvuAX3h+VmBTLoCT0QW4gQEyF3LBj5riCU",
	"ZD1bdLrQ3IAaEWpEZJofPehCBybraj1kuous3xNsdKG/yw/J5FMm+Qr0MGSa4GNvCSpd3J6GjRzrIdOv",
	"eyJw/M7UM5n21MrPBqkzrIv4naRzFkYwQvwWsW7hGmvrbt5fp9wElSUo+e45QU4WS7Q5n50tt86MXiUW",
	"9r2tvmTizc3Vb6k+WZgdyCV6BM5WMs5kR+LG1afpZPqvixovpkGBLKZ1+/EL23kkjrKipB2TN9zoQteW",
	"m/5+UWKX17ihkMt0/vCTHiT/AGVPlh8/iDA479PLXiimLhbqezAEdEizYfbi2Yet9QnYoYNzqMS0boPR",
	"ecutzKzQWUZSl++udKYfgMLo0nKxXLwQdvTgjLe60C8Xy8Vy0haTzU3PTV7HjspL07b3ptzKRA2c+v56",
	"AqgNYafGPtSRn4xArqoj0BvCbg8hCB5dGB16ufwl5d41nmJU3aTIy94ZMv1q+XO6wG6mDFQrW0HZoByy",
	"AmfuW6j0cXV0cft05PHtelhnmk0dpO3i2FrgJ2a01m3zp0h9t4XHIe4CQ6YDBgqR0UoasU2yXe/s8fq4",
	"t5l6yKazbGbPSDYew4zhbywFnvR9stzsFF+WJfaOE++vrdvK4Tjv+4/LpfyU6BhcjGa8b20Z1+d/BnSH",
	"I/dLO/jsEIqNmW7SnlqFuzr9EOT4lbz/38IKaazmrOXX/95sQZ65/bGHwL9i9fifGX18jQzDeGZ+q+mu",
	"pi3W2PPnq3nlHkxrK8OgCDYEoRmvaWVcpc7u7aS6wpzY/SoNgttR94sZ3WP8XfDp4fFsuRPN5/W+n+Kc",
	"vXNOpU2go4vqa54U6bttprHEwmeYMj5BZu+qGljhJ6f6AKTilXluw1vg1Vc9I/9R6k7NSh46ifbkEXFi",
	"wf6pEueGbP8tEdfD3wEAAP//sIahfEULAAA=",
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
