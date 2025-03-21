package main

import (
	"fmt"
	"go/format"
	"os"
	"strings"
	"unicode"

	"github.com/okocraft/monitor/internal/domain/permission/codegen/definition"
	"github.com/okocraft/monitor/internal/domain/permission/codegen/identifier"
)

func main() {
	idFile, err := os.ReadFile("permission_ids.json")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	idProvider, err := identifier.DecodeFromJson(idFile)
	if err != nil {
		panic(err)
	}

	w := strings.Builder{}

	_, err = w.WriteString("// Code generated by permission/codegen; DO NOT EDIT\n\n")
	if err != nil {
		panic(err)
	}

	_, err = w.WriteString("package permission\n")
	if err != nil {
		panic(err)
	}

	for _, perm := range definition.GetPermissions() {
		id := idProvider.GetOrCreateID(perm.Name)
		w.WriteString(fmt.Sprintf(
			"var %s = Permission{\n"+
				"ID: %d,\n"+
				"Name: \"%s\",\n"+
				"DefaultValue: %v,\n"+
				"}\n\n",
			toVarName(perm), id, perm.Name, perm.DefaultValue,
		))
	}

	formatted, err := format.Source([]byte(w.String()))
	err = os.WriteFile("permissions.gen.go", formatted, os.ModePerm)
	if err != nil {
		panic(err)
	}

	data, err := idProvider.EncodeToJson()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("permission_ids.json", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func toVarName(perm definition.DefinedPermission) string {
	upperCase := true
	varName := strings.Builder{}
	for _, c := range perm.Name {
		switch {
		case c == '.':
			upperCase = true
		case upperCase:
			varName.WriteRune(unicode.ToUpper(c))
			upperCase = false
		default:
			varName.WriteRune(c)
		}
	}
	return varName.String()
}
