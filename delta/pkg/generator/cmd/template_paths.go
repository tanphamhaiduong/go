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
		{Source: "pkg/generator/internal/tmpl/modules/type/z_repository.tmpl", Destination: "internal/modules/{{module}}/z_repository.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/z_repository_test.tmpl", Destination: "internal/modules/{{module}}/z_repository_test.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/repository.tmpl", Destination: "internal/modules/{{module}}/repository.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/type/repository_test.tmpl", Destination: "internal/modules/{{module}}/repository_test.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/type/z_handlers.tmpl", Destination: "internal/modules/{{module}}/z_handlers.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/z_handlers_test.tmpl", Destination: "internal/modules/{{module}}/z_handlers_test.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/handlers.tmpl", Destination: "internal/modules/{{module}}/handlers.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/type/handlers_test.tmpl", Destination: "internal/modules/{{module}}/handlers_test.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/type/z_graphql.tmpl", Destination: "internal/modules/{{module}}/z_graphql.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/z_graphql_test.tmpl", Destination: "internal/modules/{{module}}/z_graphql_test.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type/graphql.tmpl", Destination: "internal/modules/{{module}}/graphql.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/type/graphql_test.tmpl", Destination: "internal/modules/{{module}}/graphql_test.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/modules/z_init.tmpl", Destination: "internal/modules/z_init.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/z_type.tmpl", Destination: "internal/modules/z_{{module}}.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/modules/type.tmpl", Destination: "internal/modules/{{module}}.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/arguments/z_type.tmpl", Destination: "internal/arguments/z_{{module}}.go", IsOverWrite: true},
		{Source: "pkg/generator/internal/tmpl/arguments/type.tmpl", Destination: "internal/arguments/{{module}}.go", IsOverWrite: false},
		{Source: "pkg/generator/internal/tmpl/models/z_type.tmpl", Destination: "internal/models/z_{{module}}.go", IsOverWrite: true},
	}
)
