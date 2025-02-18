package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Specialization struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('specialization_id_seq')"`
	Name *string `json:"name,omitempty" column:"name:name;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"specialization" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorSpecializations []*Doctors `json:"doctor_specializations,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:specialization_id"`
	Facilities []*Facilities `json:"facilities,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:specialization_id;targetPrimaryKey:id;targetForeign:specialization_id"`
}
