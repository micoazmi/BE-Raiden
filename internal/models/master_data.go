package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type MasterData struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('master_data_id_seq')"`
	DoctorId *int32 `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable"`
	ServiceId *int32 `json:"service_id,omitempty" column:"name:service_id;type:integer;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"master_data" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Users []*Users `json:"users,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:master_data_id;targetPrimaryKey:id;targetForeign:master_data_id"`
	ReservationMasters []*Reservations `json:"reservation_masters,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:master_data_id"`
	Doctor *Doctors `json:"doctor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Service *Services `json:"service,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:service_id"`
}
