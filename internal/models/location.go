package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Location struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('location_id_seq')"`
	City *string `json:"city,omitempty" column:"name:city;type:varchar;nullable"`
	Street *string `json:"street,omitempty" column:"name:street;type:varchar;nullable"`
	PostalCode *int32 `json:"postal_code,omitempty" column:"name:postal_code;type:integer;nullable"`
	Country *string `json:"country,omitempty" column:"name:country;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"location" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
