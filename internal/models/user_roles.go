package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type UserRoles struct {
	db.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	UserId uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	Role interface{} `json:"role,omitempty" column:"name:role;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"user_roles" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
