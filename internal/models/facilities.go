package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Facilities struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('facilities_id_seq')"`
	Name string `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	LocationId *int32 `json:"location_id,omitempty" column:"name:location_id;type:integer;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"facilities" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctors []*Doctors `json:"doctors,omitempty" join:"joinType:manyToMany;through:schedule;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
	DoctorFacilities []*Doctors `json:"doctor_facilities,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	ServiceFacilities []*Services `json:"service_facilities,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	ScheduleFacilities []*Schedule `json:"schedule_facilities,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	Specializations []*Specialization `json:"specializations,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
}
