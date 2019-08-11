package utils

import (
	"go/build"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

var (
	statFile = os.Stat
	regIDExt = regexp.MustCompile(`(^[a-zA-Z]{1,})(ID)$`)
)

// ToCamelCase ...
func ToCamelCase(value string) string {
	return strcase.ToCamel(value)
}

// ToCamelCaseExceptID ...
func ToCamelCaseExceptID(value string) string {
	if value == "ID" {
		return "Id"
	}
	valueCamelCase := strcase.ToCamel(value)
	if regIDExt.MatchString(valueCamelCase) {
		valueCamelCase = regIDExt.ReplaceAllString(valueCamelCase, `${1}Id`)
	}
	return valueCamelCase
}

// ToLowerCamelCase ...
func ToLowerCamelCase(value string) string {
	if value == "ID" {
		return "id"
	}
	if value == "URL" {
		return "url"
	}
	valueLowerCamelCase := strcase.ToLowerCamel(value)
	if regIDExt.MatchString(valueLowerCamelCase) {
		valueLowerCamelCase = regIDExt.ReplaceAllString(valueLowerCamelCase, `${1}Id`)
	}
	return valueLowerCamelCase
}

// ToLowerCamelCaseExceptID ...
func ToLowerCamelCaseExceptID(value string) string {
	if value == "ID" {
		return "Id"
	}
	valueLowerCamelCase := strcase.ToLowerCamel(value)
	if regIDExt.MatchString(valueLowerCamelCase) {
		valueLowerCamelCase = regIDExt.ReplaceAllString(valueLowerCamelCase, `${1}Id`)
	}
	return valueLowerCamelCase
}

// ToSnakeCase ...
func ToSnakeCase(value string) string {
	return strcase.ToSnake(value)
}

// ToLowerCase ...
func ToLowerCase(value string) string {
	return strings.ToLower(value)
}

// ToPlural ...
func ToPlural(value string) string {
	return inflection.Plural(value)
}

// ProjectDir ...
func ProjectDir() string {
	return path.Join(build.Default.GOPATH, "src", "github.com", "tanphamhaiduong", "go", "delta")
}

// FileExists ...
func FileExists(path string) bool {
	if _, err := statFile(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Add ...
func Add(a int, b int) int {
	return a + b
}
