package schema

import (
	"encoding/json"

	"github.com/deakodev/go-dotify/dot"
)

type Schema interface {
	Type() string
	Dotify() dot.Graph
}

type SchemaType int

const (
	SchemaNone SchemaType = iota
	SchemaBST
	SchemaRBT
)

var schemaName = map[SchemaType]string{
	SchemaNone: "None",
	SchemaBST:  "BST",
	SchemaRBT:  "RBT",
}

func (st SchemaType) String() string {
	return schemaName[st]
}

func Unmarshal(data []byte) (Schema, error) {
	var meta struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	var schema Schema
	// switch meta.Type {
	// // case "BST":
	// // 	schema = &BSTree{}
	// // case "RBT":
	// // 	schema = &RBTree{}
	// default:
	// 	return nil, fmt.Errorf("unsupported type: %s", meta.Type)
	// }

	// if err := json.Unmarshal(data, &schema); err != nil {
	// 	return nil, err
	// }

	return schema, nil
}
