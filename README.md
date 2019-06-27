# GraphQL Go Services boiler plate
## Why we need boilerplate
For every project need CRUD, we need to build crud for many entities in your project and it cause lot of time for it. So this boilerplate help you to generate all CRUD for project in 5 minutes by define models

## How to use

### Step 1
`cd delta`
### Step 2
`make install`
### Step 3
`cd database`
### Step 4: Migrate database
`goose mysql root:root@/delta up`
### Step 5: 


## Dependencies
[Squirrel]:https://github.com/Masterminds/squirrel
[GraphQL]: github.com/graphql-go/graphql
[GraphQL Handler]:github.com/graphql-go/handler
[Strcase]: github.com/iancoleman/strcase
[Inflection]: github.com/jinzhu/inflection
[Mapstructure]: github.com/mitchellh/mapstructure
[Validation V9]: gopkg.in/go-playground/validator.v9
[Mysql Driver]: github.com/go-sql-driver/mysql
## Contributor
@tanphamhaiduong
