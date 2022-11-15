package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// GroupMsg holds the schema definition for the GroupMsg entity.
type GroupMsg struct {
	ent.Schema
}

// Fields of the GroupMsg.
func (GroupMsg) Fields() []ent.Field {
	return []ent.Field{
		field.String("from"),
		field.String("to"),
		field.String("type"),
		field.String("content"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Indexes of the GroupMsg.
func (GroupMsg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("from", "to"),
	}
}

// Edges of the GroupMsg.
func (GroupMsg) Edges() []ent.Edge {
	return nil
}
