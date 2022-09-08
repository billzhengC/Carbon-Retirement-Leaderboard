// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type TUser struct {
	ent.Schema
}

func (TUser) Fields() []ent.Field {
	return []ent.Field{field.Uint64("id").Comment("auto increment primary key"), field.String("wallet_pub").Unique().Comment("wallet public key"), field.String("wallet_type").Comment("wallet type phantom"), field.String("uname").Comment("name"), field.String("face").Comment("avatar"), field.Bool("gender").Comment("gender 0 secret 1 female 2 male"), field.String("twitter").Comment("twitter"), field.String("email").Comment("email"), field.String("about").Comment("about"), field.Time("last_login_time").Comment("last login time").Default(time.Now()).UpdateDefault(time.Now), field.Time("mtime").Comment("modify time").Default(time.Now()).UpdateDefault(time.Now), field.Time("ctime").Comment("create time").Default(time.Now())}
}
func (TUser) Edges() []ent.Edge {
	return nil
}
func (TUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_users",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_0900_ai_ci",
		},
	}
}
