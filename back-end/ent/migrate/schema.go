// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// GroupMsgsColumns holds the columns for the "group_msgs" table.
	GroupMsgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// GroupMsgsTable holds the schema information for the "group_msgs" table.
	GroupMsgsTable = &schema.Table{
		Name:       "group_msgs",
		Columns:    GroupMsgsColumns,
		PrimaryKey: []*schema.Column{GroupMsgsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "groupmsg_from_to",
				Unique:  false,
				Columns: []*schema.Column{GroupMsgsColumns[1], GroupMsgsColumns[2]},
			},
		},
	}
	// PrivateMsgsColumns holds the columns for the "private_msgs" table.
	PrivateMsgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "read", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PrivateMsgsTable holds the schema information for the "private_msgs" table.
	PrivateMsgsTable = &schema.Table{
		Name:       "private_msgs",
		Columns:    PrivateMsgsColumns,
		PrimaryKey: []*schema.Column{PrivateMsgsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "privatemsg_from_to",
				Unique:  false,
				Columns: []*schema.Column{PrivateMsgsColumns[1], PrivateMsgsColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		GroupMsgsTable,
		PrivateMsgsTable,
		UsersTable,
		GroupUsersTable,
	}
)

func init() {
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}
