package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("title").NotEmpty(),
		field.Time("startAt"),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user", User.Type).Ref("schedules"),
	}
}
