package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// PrivateMsg holds the schema definition for the PrivateMsg entity.
type PrivateMsg struct {
	ent.Schema
}

// Fields of the PrivateMsg.
func (PrivateMsg) Fields() []ent.Field {
	return []ent.Field{
		field.String("from"),
		field.String("to"),
		field.String("type"),
		field.String("content"),
		field.Bool("read"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Indexes of the PrivateMsg.
func (PrivateMsg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("from", "to"),
	}
}

// Edges of the PrivateMsg.
func (PrivateMsg) Edges() []ent.Edge {
	return nil
}
