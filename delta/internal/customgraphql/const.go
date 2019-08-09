package customgraphql

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// This is custom scala example for graphql
func serializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case mysql.NullTime:
		if value.Valid != true {
			return nil
		}
		buff, err := value.Time.MarshalText()
		if err != nil {
			return nil
		}

		return string(buff)
	case *mysql.NullTime:
		if value == nil {
			return nil
		}
		return serializeDateTime(*value)
	default:
		return nil
	}
}

func unserializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := time.Time{}
		err := t.UnmarshalText(value)
		if err != nil {
			return mysql.NullTime{
				Valid: false,
			}
		}

		return mysql.NullTime{
			Time:  t,
			Valid: true,
		}
	case string:
		return unserializeDateTime([]byte(value))
	case *string:
		if value == nil {
			return mysql.NullTime{
				Valid: false,
			}
		}
		return unserializeDateTime([]byte(*value))
	default:
		return mysql.NullTime{
			Valid: false,
		}
	}
}

var NullDateTime = graphql.NewScalar(graphql.ScalarConfig{
	Name: "NullDateTime",
	Description: "The `NullDateTime` scalar type represents a mysql.NullDateTime." +
		" The NullDateTime is serialized as an RFC 3339 quoted string",
	Serialize:  serializeDateTime,
	ParseValue: unserializeDateTime,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return unserializeDateTime(valueAST.Value)
		}
		return nil
	},
})
