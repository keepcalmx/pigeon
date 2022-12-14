// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/keepcalmx/go-pigeon/ent/privatemsg"
)

// PrivateMsg is the model entity for the PrivateMsg schema.
type PrivateMsg struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PrivateMsg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case privatemsg.FieldRead:
			values[i] = new(sql.NullBool)
		case privatemsg.FieldID:
			values[i] = new(sql.NullInt64)
		case privatemsg.FieldFrom, privatemsg.FieldTo, privatemsg.FieldType, privatemsg.FieldContent:
			values[i] = new(sql.NullString)
		case privatemsg.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PrivateMsg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PrivateMsg fields.
func (pm *PrivateMsg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case privatemsg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pm.ID = int(value.Int64)
		case privatemsg.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				pm.From = value.String
			}
		case privatemsg.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				pm.To = value.String
			}
		case privatemsg.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pm.Type = value.String
			}
		case privatemsg.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pm.Content = value.String
			}
		case privatemsg.FieldRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field read", values[i])
			} else if value.Valid {
				pm.Read = value.Bool
			}
		case privatemsg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pm.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PrivateMsg.
// Note that you need to call PrivateMsg.Unwrap() before calling this method if this PrivateMsg
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PrivateMsg) Update() *PrivateMsgUpdateOne {
	return (&PrivateMsgClient{config: pm.config}).UpdateOne(pm)
}

// Unwrap unwraps the PrivateMsg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *PrivateMsg) Unwrap() *PrivateMsg {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PrivateMsg is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PrivateMsg) String() string {
	var builder strings.Builder
	builder.WriteString("PrivateMsg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("from=")
	builder.WriteString(pm.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(pm.To)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pm.Type)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pm.Content)
	builder.WriteString(", ")
	builder.WriteString("read=")
	builder.WriteString(fmt.Sprintf("%v", pm.Read))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pm.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PrivateMsgs is a parsable slice of PrivateMsg.
type PrivateMsgs []*PrivateMsg

func (pm PrivateMsgs) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}
