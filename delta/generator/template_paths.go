package main

// Path ...
type Path struct {
	Source      string
	Destination string
	IsOverWrite bool
}

// Paths ...
type Paths []Path

var (
	templatePaths = Paths{
		{Source: "generator/tmpl/modules/type/z_repository.tmpl", Destination: "server/modules/{{module}}/z_repository.go", IsOverWrite: true},
		{Source: "generator/tmpl/modules/type/repository.tmpl", Destination: "server/modules/{{module}}/repository.go", IsOverWrite: false},
		{Source: "generator/tmpl/modules/type/z_handlers.tmpl", Destination: "server/modules/{{module}}/z_handlers.go", IsOverWrite: true},
		{Source: "generator/tmpl/modules/type/handlers.tmpl", Destination: "server/modules/{{module}}/handlers.go", IsOverWrite: false},
		{Source: "generator/tmpl/modules/type/z_graphql.tmpl", Destination: "server/modules/{{module}}/z_graphql.go", IsOverWrite: true},
		{Source: "generator/tmpl/modules/type/graphql.tmpl", Destination: "server/modules/{{module}}/graphql.go", IsOverWrite: false},
		{Source: "generator/tmpl/modules/z_init.tmpl", Destination: "server/modules/z_init.go", IsOverWrite: true},
		{Source: "generator/tmpl/modules/z_type.tmpl", Destination: "server/modules/z_{{module}}.go", IsOverWrite: true},
		{Source: "generator/tmpl/modules/type.tmpl", Destination: "server/modules/{{module}}.go", IsOverWrite: false},
		{Source: "generator/tmpl/arguments/z_type.tmpl", Destination: "server/arguments/z_{{module}}.go", IsOverWrite: true},
		{Source: "generator/tmpl/arguments/type.tmpl", Destination: "server/arguments/{{module}}.go", IsOverWrite: false},
		{Source: "generator/tmpl/models/z_type.tmpl", Destination: "server/models/z_{{module}}.go", IsOverWrite: true},
	}
)
