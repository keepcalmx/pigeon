package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.String("email").Optional(),
		field.String("nickname").Optional(),
		field.String("avatar").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional(),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).Ref("users"),
	}
}
