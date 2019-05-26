package modules

import (
	"github.com/tanphamhaiduong/go/delta/generator/gotype"
)

// Field is one field of struct
type Field struct {
	Name               string
	GoType             string
	QueryType          string
	Validate           string
	GraphQLType        string
	GraphQLDescription string
	GraphQLRequired    bool
}

// Fields list of Field
type Fields []Field

// IsInt64 ...
func (f *Field) IsInt64() bool {
	return f.GoType == gotype.Int64
}

// IsNotBoolean ...
func (f *Field) IsNotBoolean() bool {
	return f.GoType != gotype.Bool
}

// IsNotIDField ...
func (f *Field) IsNotIDField() bool {
	return f.Name != "ID"
}

// Module need to generate
type Module struct {
	Name   string
	Fields Fields
}

// Modules list of Struct
type Modules []Module

var (
	// declare struct and methods for each struct
	modules = Modules{}
)

// GetModules ...
func GetModules() Modules {
	return modules
}
