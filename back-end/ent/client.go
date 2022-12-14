// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/keepcalmx/go-pigeon/ent/migrate"

	"github.com/keepcalmx/go-pigeon/ent/group"
	"github.com/keepcalmx/go-pigeon/ent/groupmsg"
	"github.com/keepcalmx/go-pigeon/ent/privatemsg"
	"github.com/keepcalmx/go-pigeon/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// GroupMsg is the client for interacting with the GroupMsg builders.
	GroupMsg *GroupMsgClient
	// PrivateMsg is the client for interacting with the PrivateMsg builders.
	PrivateMsg *PrivateMsgClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Group = NewGroupClient(c.config)
	c.GroupMsg = NewGroupMsgClient(c.config)
	c.PrivateMsg = NewPrivateMsgClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Group:      NewGroupClient(cfg),
		GroupMsg:   NewGroupMsgClient(cfg),
		PrivateMsg: NewPrivateMsgClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Group:      NewGroupClient(cfg),
		GroupMsg:   NewGroupMsgClient(cfg),
		PrivateMsg: NewPrivateMsgClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Group.Use(hooks...)
	c.GroupMsg.Use(hooks...)
	c.PrivateMsg.Use(hooks...)
	c.User.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// GroupMsgClient is a client for the GroupMsg schema.
type GroupMsgClient struct {
	config
}

// NewGroupMsgClient returns a client for the GroupMsg from the given config.
func NewGroupMsgClient(c config) *GroupMsgClient {
	return &GroupMsgClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `groupmsg.Hooks(f(g(h())))`.
func (c *GroupMsgClient) Use(hooks ...Hook) {
	c.hooks.GroupMsg = append(c.hooks.GroupMsg, hooks...)
}

// Create returns a builder for creating a GroupMsg entity.
func (c *GroupMsgClient) Create() *GroupMsgCreate {
	mutation := newGroupMsgMutation(c.config, OpCreate)
	return &GroupMsgCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GroupMsg entities.
func (c *GroupMsgClient) CreateBulk(builders ...*GroupMsgCreate) *GroupMsgCreateBulk {
	return &GroupMsgCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GroupMsg.
func (c *GroupMsgClient) Update() *GroupMsgUpdate {
	mutation := newGroupMsgMutation(c.config, OpUpdate)
	return &GroupMsgUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupMsgClient) UpdateOne(gm *GroupMsg) *GroupMsgUpdateOne {
	mutation := newGroupMsgMutation(c.config, OpUpdateOne, withGroupMsg(gm))
	return &GroupMsgUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupMsgClient) UpdateOneID(id int) *GroupMsgUpdateOne {
	mutation := newGroupMsgMutation(c.config, OpUpdateOne, withGroupMsgID(id))
	return &GroupMsgUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GroupMsg.
func (c *GroupMsgClient) Delete() *GroupMsgDelete {
	mutation := newGroupMsgMutation(c.config, OpDelete)
	return &GroupMsgDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupMsgClient) DeleteOne(gm *GroupMsg) *GroupMsgDeleteOne {
	return c.DeleteOneID(gm.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupMsgClient) DeleteOneID(id int) *GroupMsgDeleteOne {
	builder := c.Delete().Where(groupmsg.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupMsgDeleteOne{builder}
}

// Query returns a query builder for GroupMsg.
func (c *GroupMsgClient) Query() *GroupMsgQuery {
	return &GroupMsgQuery{
		config: c.config,
	}
}

// Get returns a GroupMsg entity by its id.
func (c *GroupMsgClient) Get(ctx context.Context, id int) (*GroupMsg, error) {
	return c.Query().Where(groupmsg.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupMsgClient) GetX(ctx context.Context, id int) *GroupMsg {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GroupMsgClient) Hooks() []Hook {
	return c.hooks.GroupMsg
}

// PrivateMsgClient is a client for the PrivateMsg schema.
type PrivateMsgClient struct {
	config
}

// NewPrivateMsgClient returns a client for the PrivateMsg from the given config.
func NewPrivateMsgClient(c config) *PrivateMsgClient {
	return &PrivateMsgClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `privatemsg.Hooks(f(g(h())))`.
func (c *PrivateMsgClient) Use(hooks ...Hook) {
	c.hooks.PrivateMsg = append(c.hooks.PrivateMsg, hooks...)
}

// Create returns a builder for creating a PrivateMsg entity.
func (c *PrivateMsgClient) Create() *PrivateMsgCreate {
	mutation := newPrivateMsgMutation(c.config, OpCreate)
	return &PrivateMsgCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PrivateMsg entities.
func (c *PrivateMsgClient) CreateBulk(builders ...*PrivateMsgCreate) *PrivateMsgCreateBulk {
	return &PrivateMsgCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PrivateMsg.
func (c *PrivateMsgClient) Update() *PrivateMsgUpdate {
	mutation := newPrivateMsgMutation(c.config, OpUpdate)
	return &PrivateMsgUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrivateMsgClient) UpdateOne(pm *PrivateMsg) *PrivateMsgUpdateOne {
	mutation := newPrivateMsgMutation(c.config, OpUpdateOne, withPrivateMsg(pm))
	return &PrivateMsgUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrivateMsgClient) UpdateOneID(id int) *PrivateMsgUpdateOne {
	mutation := newPrivateMsgMutation(c.config, OpUpdateOne, withPrivateMsgID(id))
	return &PrivateMsgUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PrivateMsg.
func (c *PrivateMsgClient) Delete() *PrivateMsgDelete {
	mutation := newPrivateMsgMutation(c.config, OpDelete)
	return &PrivateMsgDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PrivateMsgClient) DeleteOne(pm *PrivateMsg) *PrivateMsgDeleteOne {
	return c.DeleteOneID(pm.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PrivateMsgClient) DeleteOneID(id int) *PrivateMsgDeleteOne {
	builder := c.Delete().Where(privatemsg.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrivateMsgDeleteOne{builder}
}

// Query returns a query builder for PrivateMsg.
func (c *PrivateMsgClient) Query() *PrivateMsgQuery {
	return &PrivateMsgQuery{
		config: c.config,
	}
}

// Get returns a PrivateMsg entity by its id.
func (c *PrivateMsgClient) Get(ctx context.Context, id int) (*PrivateMsg, error) {
	return c.Query().Where(privatemsg.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrivateMsgClient) GetX(ctx context.Context, id int) *PrivateMsg {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PrivateMsgClient) Hooks() []Hook {
	return c.hooks.PrivateMsg
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.GroupsTable, user.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
