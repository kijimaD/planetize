// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package generated

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

	"H4sIAAAAAAAC/7RWX2/jRBD/KsfCoxsnVwHCbycOqpOiU5TkXkDotBdvnAXba3bXiHKyhG2C7lQKCIne",
	"3VM5uGurJo0qIcQfVeqH2TpKvwVa20nsOGkcId7WuzPz+81vZ2f8GHSI5RAb2ZwB7TFgnR6yYLx8n9hd",
	"bMiVQ4mDKMcoMSEu7SRLzJEVL96iqAs08KY6j6amodQkTiv2Ap4C+K6DgAYgpXA3/obGpsHa0ChG8hRA",
	"0ecupkgH2sczminAJzNz8uhT1OHSP8eskKeOWGcdncT3rrT0FGBDC5XzuC8tJWHGHrrULOfUZOwBNbOS",
	"rfdpQ6OOGS+oMwVOSStJtmu1ksL/B6Ha0NhcqzY0ErkWUsgSX8b4Q4T0D2xOd4uMTWx/tg5cutelnacA",
	"x31kYtaTuOudGjNjT0mrsIzb/IEw17JgQnutU2oqawJzsxRQOzZcVDNxVxJpshnP+cyyWaX2DrIRhRzp",
	"d/i0Lih2OCY20IAIfxbBbyK8EOFTEZyJ8LUIh/Iz+FOExyL8XQQ/RaPDycvvhP9M+IfjZ6/HLwKgAPQl",
	"tByZGrhdvf32VvWdrdp2u/qetl3Tqu9+BBTQJdSCHGhAhxxtcRyXRUqQcYptY0qwnt77TcxOJJXwiQhP",
	"hT960KznGPQ4d5imqulOpUMsFVKOOyZSa1kqLsWrSDSyxVSaSdQfXh/s/R+iNBFzTV58JMjmFG/Q5+fv",
	"bUmTN6al8RDyMoGypbRYqrlYyoznqqqc9/fSYl/3968uX4rgDxG8kpv+KPpxP3q6nxM+Xd1imK/UtjV/",
	"yqXBx0fB5MjPQU1Onl/9s5c9W4rWnraA0lgiuExTDAfLERcsCriZCVgAbrZaiZJRPxT+aHI6HD//Pgez",
	"Q0Tox8x+Ef5ABE9EsBf1h9HFDyI8EOGZCM5Xg95PR8iNoKsvbnXgdNAuC12yJ3QR0iuQE6tET8iPxjX5",
	"DIT/SvjfCP8wvpjz5fWyQ4Q/uD74VfgvpKDf9qPR39O7PL+RQ/yzsCGHq7++nhwdA6Vcm1gY6sVWkTfY",
	"kEtOBoMUc5XNBNtdUgzcMKGNOP4K0Vt3GvfeALORunACFPAFoixxqlWqlapkTRxkQwcDDWxXapWqHKGQ",
	"92I1VN21rLgHGGiJuMlpHIFCuXdPBxq4m+5SxBxis7gPZ/8NUjfP87x/AwAA//+hPKagxAsAAA==",
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
