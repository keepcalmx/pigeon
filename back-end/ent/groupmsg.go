// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/keepcalmx/go-pigeon/ent/groupmsg"
)

// GroupMsg is the model entity for the GroupMsg schema.
type GroupMsg struct {
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
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMsg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmsg.FieldID:
			values[i] = new(sql.NullInt64)
		case groupmsg.FieldFrom, groupmsg.FieldTo, groupmsg.FieldType, groupmsg.FieldContent:
			values[i] = new(sql.NullString)
		case groupmsg.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupMsg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMsg fields.
func (gm *GroupMsg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmsg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gm.ID = int(value.Int64)
		case groupmsg.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				gm.From = value.String
			}
		case groupmsg.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				gm.To = value.String
			}
		case groupmsg.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				gm.Type = value.String
			}
		case groupmsg.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				gm.Content = value.String
			}
		case groupmsg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GroupMsg.
// Note that you need to call GroupMsg.Unwrap() before calling this method if this GroupMsg
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMsg) Update() *GroupMsgUpdateOne {
	return (&GroupMsgClient{config: gm.config}).UpdateOne(gm)
}

// Unwrap unwraps the GroupMsg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMsg) Unwrap() *GroupMsg {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMsg is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMsg) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMsg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("from=")
	builder.WriteString(gm.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(gm.To)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(gm.Type)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(gm.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GroupMsgs is a parsable slice of GroupMsg.
type GroupMsgs []*GroupMsg

func (gm GroupMsgs) config(cfg config) {
	for _i := range gm {
		gm[_i].config = cfg
	}
}
