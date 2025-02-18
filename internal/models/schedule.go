package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Schedule struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('schedule_id_seq')"`
	DoctorId int32 `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable:false"`
	FacilityId int32 `json:"facility_id,omitempty" column:"name:facility_id;type:integer;nullable:false"`
	Date time.Time `json:"date,omitempty" column:"name:date;type:date;nullable:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedule" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor *Doctors `json:"doctor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Facility *Facilities `json:"facility,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
}
