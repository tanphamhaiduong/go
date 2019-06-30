# GraphQL Go Services boiler plate
## Why we need boilerplate
For every project need CRUD, we need to build crud for many entities in your project and it cause lot of time for it. So this boilerplate help you to generate all CRUD for project in 5 minutes by define models

## How to use
### Step 1: Change DIR to project delta
`cd delta`

### Step 2: Install dependencies
`make install`

### Step 3: Change DIR to database
`cd database`

### Step 4: Migrate database
`goose mysql <username>:<password>@/delta up`

### Step 5: Define entity
You can define entity as database schema in this folder
`pkg/generator/internal/modules`  
For Example: New Entity Name is `<new entity>`

### Step 6: Generate code
Back to DIR delta  
Run command `make gen`  
This will generate code and test for `<new entity>` modules

### Step 7: Change DIR to
`internal/modules/<new entity>`

### Step 8: Generate Mock for testing
Run command `sudo go generate`  
Run command `go test ./...` for check test pass or not

### Step 9: Add to main file for init
Go to `cmd/main.go`
```
// <New Entity>
<new entity>Repository := <new entity>.NewRepository(db)
<new entity>Handler := <new entity>.NewHandler(<new entity>Repository)
<new entity>Resolver := <new entity>.NewResolver(<new entity>Handler)

resolvers := modules.NewResolver(
  companyResolver,
  featureResolver,
  permissionResolver,
  roleResolver,
  rolepermissionResolver,
  userResolver,
  <new entity>
)
```


## Dependencies
[Squirrel](https://github.com/Masterminds/squirrel)  
[GraphQL](https://github.com/graphql-go/graphql)  
[GraphQL Handler](https://github.com/graphql-go/handler)  
[Strcase](https://github.com/iancoleman/strcase)  
[Inflection](https://github.com/jinzhu/inflection)  
[Mapstructure](https://github.com/mitchellh/mapstructure)  
[Validation V9](https://gopkg.in/go-playground/validator.v9)  
[Mysql Driver](https://github.com/go-sql-driver/mysql)  
[Mux](https://github.com/gorilla/mux)  
[Logrus](https://github.com/sirupsen/logrus)
## Contributor
@tanphamhaiduong
