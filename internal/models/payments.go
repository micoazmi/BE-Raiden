package models

import (
	"time"
	"github.com/sev-2/raiden/pkg/db"
)

type Payments struct {
	db.ModelBase
	Id int32 `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('payments_id_seq')"`
	ReservationId *int32 `json:"reservation_id,omitempty" column:"name:reservation_id;type:integer;nullable"`
	Amount float64 `json:"amount,omitempty" column:"name:amount;type:numeric;nullable:false"`
	PaymentMethod *string `json:"payment_method,omitempty" column:"name:payment_method;type:varchar;nullable"`
	Status *string `json:"status,omitempty" column:"name:status;type:varchar;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payments" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Reservation *Reservations `json:"reservation,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
}
