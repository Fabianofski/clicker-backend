package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GameState holds the schema definition for the GameState entity.
type GameState struct {
	ent.Schema
}

// Fields of the GameState.
func (GameState) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
	}
}

// Edges of the GameState.
func (GameState) Edges() []ent.Edge {
	return nil
}
