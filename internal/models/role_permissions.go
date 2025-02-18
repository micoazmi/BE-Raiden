package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type RolePermissions struct {
	db.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Role interface{} `json:"role,omitempty" column:"name:role;nullable:false"`
	Permission interface{} `json:"permission,omitempty" column:"name:permission;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"role_permissions" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
