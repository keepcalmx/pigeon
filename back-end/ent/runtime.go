// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/keepcalmx/go-pigeon/ent/group"
	"github.com/keepcalmx/go-pigeon/ent/groupmsg"
	"github.com/keepcalmx/go-pigeon/ent/privatemsg"
	"github.com/keepcalmx/go-pigeon/ent/schema"
	"github.com/keepcalmx/go-pigeon/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[3].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(time.Time)
	groupmsgFields := schema.GroupMsg{}.Fields()
	_ = groupmsgFields
	// groupmsgDescCreatedAt is the schema descriptor for created_at field.
	groupmsgDescCreatedAt := groupmsgFields[4].Descriptor()
	// groupmsg.DefaultCreatedAt holds the default value on creation for the created_at field.
	groupmsg.DefaultCreatedAt = groupmsgDescCreatedAt.Default.(time.Time)
	privatemsgFields := schema.PrivateMsg{}.Fields()
	_ = privatemsgFields
	// privatemsgDescCreatedAt is the schema descriptor for created_at field.
	privatemsgDescCreatedAt := privatemsgFields[5].Descriptor()
	// privatemsg.DefaultCreatedAt holds the default value on creation for the created_at field.
	privatemsg.DefaultCreatedAt = privatemsgDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
}
