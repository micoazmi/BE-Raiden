package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Services struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('services_id_seq')"`
	Name string `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	FacilityId *int32 `json:"facility_id,omitempty" column:"name:facility_id;type:integer;nullable"`
	Price float64 `json:"price,omitempty" column:"name:price;type:numeric;nullable:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"services" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Facility *Facilities `json:"facility,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
	MasterDatumServices []*MasterData `json:"master_datum_services,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:service_id"`
	Doctors []*Doctors `json:"doctors,omitempty" join:"joinType:manyToMany;through:master_data;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
}
