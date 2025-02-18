package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Reservations struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('reservations_id_seq')"`
	UserId *int32 `json:"user_id,omitempty" column:"name:user_id;type:integer;nullable"`
	MasterDataId *int32 `json:"master_data_id,omitempty" column:"name:master_data_id;type:integer;nullable"`
	Date time.Time `json:"date,omitempty" column:"name:date;type:date;nullable:false"`
	Time time.Time `json:"time,omitempty" column:"name:time;type:time;nullable:false"`
	Status *string `json:"status,omitempty" column:"name:status;type:varchar;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`
	UserUuid *uuid.UUID `json:"user_uuid,omitempty" column:"name:user_uuid;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	MasterDatumMaster *MasterData `json:"master_datum_master,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:master_data_id"`
	PaymentReservations []*Payments `json:"payment_reservations,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
}
