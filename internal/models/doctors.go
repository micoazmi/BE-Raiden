package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('doctors_id_seq')"`
	UserId *int32 `json:"user_id,omitempty" column:"name:user_id;type:integer;nullable"`
	SpecializationId *int32 `json:"specialization_id,omitempty" column:"name:specialization_id;type:integer;nullable"`
	FacilityId *int32 `json:"facility_id,omitempty" column:"name:facility_id;type:integer;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Facilities []*Facilities `json:"facilities,omitempty" join:"joinType:manyToMany;through:schedule;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	Specialization *Specialization `json:"specialization,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:specialization_id"`
	Facility *Facilities `json:"facility,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
	MasterDatumDoctors []*MasterData `json:"master_datum_doctors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	ScheduleDoctors []*Schedule `json:"schedule_doctors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	Services []*Services `json:"services,omitempty" join:"joinType:manyToMany;through:master_data;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
