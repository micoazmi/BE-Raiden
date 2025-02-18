package models

import (
	"github.com/google/uuid"
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Users struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('users_id_seq')"`
	Name string `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	Email string `json:"email,omitempty" column:"name:email;type:varchar;nullable:false;unique"`
	Password string `json:"password,omitempty" column:"name:password;type:text;nullable:false"`
	UserUuid *uuid.UUID `json:"user_uuid,omitempty" column:"name:user_uuid;type:uuid;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ReservationUsers []*Reservations `json:"reservation_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	MasterData []*MasterData `json:"master_data,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
}
