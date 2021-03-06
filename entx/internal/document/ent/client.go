// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/bearchit/gox/entx/internal/document/ent/migrate"

	"github.com/bearchit/gox/entx/internal/document/ent/collection"
	"github.com/bearchit/gox/entx/internal/document/ent/document"
	"github.com/bearchit/gox/entx/internal/document/ent/revision"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Collection is the client for interacting with the Collection builders.
	Collection *CollectionClient
	// Document is the client for interacting with the Document builders.
	Document *DocumentClient
	// Revision is the client for interacting with the Revision builders.
	Revision *RevisionClient
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
	c.Collection = NewCollectionClient(c.config)
	c.Document = NewDocumentClient(c.config)
	c.Revision = NewRevisionClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		Collection: NewCollectionClient(cfg),
		Document:   NewDocumentClient(cfg),
		Revision:   NewRevisionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		Collection: NewCollectionClient(cfg),
		Document:   NewDocumentClient(cfg),
		Revision:   NewRevisionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Collection.
//		Query().
//		Count(ctx)
//
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
	c.Collection.Use(hooks...)
	c.Document.Use(hooks...)
	c.Revision.Use(hooks...)
}

// CollectionClient is a client for the Collection schema.
type CollectionClient struct {
	config
}

// NewCollectionClient returns a client for the Collection from the given config.
func NewCollectionClient(c config) *CollectionClient {
	return &CollectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `collection.Hooks(f(g(h())))`.
func (c *CollectionClient) Use(hooks ...Hook) {
	c.hooks.Collection = append(c.hooks.Collection, hooks...)
}

// Create returns a create builder for Collection.
func (c *CollectionClient) Create() *CollectionCreate {
	mutation := newCollectionMutation(c.config, OpCreate)
	return &CollectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Collection entities.
func (c *CollectionClient) CreateBulk(builders ...*CollectionCreate) *CollectionCreateBulk {
	return &CollectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Collection.
func (c *CollectionClient) Update() *CollectionUpdate {
	mutation := newCollectionMutation(c.config, OpUpdate)
	return &CollectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CollectionClient) UpdateOne(co *Collection) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollection(co))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CollectionClient) UpdateOneID(id int) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollectionID(id))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Collection.
func (c *CollectionClient) Delete() *CollectionDelete {
	mutation := newCollectionMutation(c.config, OpDelete)
	return &CollectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CollectionClient) DeleteOne(co *Collection) *CollectionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CollectionClient) DeleteOneID(id int) *CollectionDeleteOne {
	builder := c.Delete().Where(collection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CollectionDeleteOne{builder}
}

// Query returns a query builder for Collection.
func (c *CollectionClient) Query() *CollectionQuery {
	return &CollectionQuery{
		config: c.config,
	}
}

// Get returns a Collection entity by its id.
func (c *CollectionClient) Get(ctx context.Context, id int) (*Collection, error) {
	return c.Query().Where(collection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CollectionClient) GetX(ctx context.Context, id int) *Collection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CollectionClient) Hooks() []Hook {
	return c.hooks.Collection
}

// DocumentClient is a client for the Document schema.
type DocumentClient struct {
	config
}

// NewDocumentClient returns a client for the Document from the given config.
func NewDocumentClient(c config) *DocumentClient {
	return &DocumentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `document.Hooks(f(g(h())))`.
func (c *DocumentClient) Use(hooks ...Hook) {
	c.hooks.Document = append(c.hooks.Document, hooks...)
}

// Create returns a create builder for Document.
func (c *DocumentClient) Create() *DocumentCreate {
	mutation := newDocumentMutation(c.config, OpCreate)
	return &DocumentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Document entities.
func (c *DocumentClient) CreateBulk(builders ...*DocumentCreate) *DocumentCreateBulk {
	return &DocumentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Document.
func (c *DocumentClient) Update() *DocumentUpdate {
	mutation := newDocumentMutation(c.config, OpUpdate)
	return &DocumentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DocumentClient) UpdateOne(d *Document) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocument(d))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DocumentClient) UpdateOneID(id int) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocumentID(id))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Document.
func (c *DocumentClient) Delete() *DocumentDelete {
	mutation := newDocumentMutation(c.config, OpDelete)
	return &DocumentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DocumentClient) DeleteOne(d *Document) *DocumentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DocumentClient) DeleteOneID(id int) *DocumentDeleteOne {
	builder := c.Delete().Where(document.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DocumentDeleteOne{builder}
}

// Query returns a query builder for Document.
func (c *DocumentClient) Query() *DocumentQuery {
	return &DocumentQuery{
		config: c.config,
	}
}

// Get returns a Document entity by its id.
func (c *DocumentClient) Get(ctx context.Context, id int) (*Document, error) {
	return c.Query().Where(document.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DocumentClient) GetX(ctx context.Context, id int) *Document {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DocumentClient) Hooks() []Hook {
	return c.hooks.Document
}

// RevisionClient is a client for the Revision schema.
type RevisionClient struct {
	config
}

// NewRevisionClient returns a client for the Revision from the given config.
func NewRevisionClient(c config) *RevisionClient {
	return &RevisionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `revision.Hooks(f(g(h())))`.
func (c *RevisionClient) Use(hooks ...Hook) {
	c.hooks.Revision = append(c.hooks.Revision, hooks...)
}

// Create returns a create builder for Revision.
func (c *RevisionClient) Create() *RevisionCreate {
	mutation := newRevisionMutation(c.config, OpCreate)
	return &RevisionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Revision entities.
func (c *RevisionClient) CreateBulk(builders ...*RevisionCreate) *RevisionCreateBulk {
	return &RevisionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Revision.
func (c *RevisionClient) Update() *RevisionUpdate {
	mutation := newRevisionMutation(c.config, OpUpdate)
	return &RevisionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RevisionClient) UpdateOne(r *Revision) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevision(r))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RevisionClient) UpdateOneID(id int) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevisionID(id))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Revision.
func (c *RevisionClient) Delete() *RevisionDelete {
	mutation := newRevisionMutation(c.config, OpDelete)
	return &RevisionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RevisionClient) DeleteOne(r *Revision) *RevisionDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RevisionClient) DeleteOneID(id int) *RevisionDeleteOne {
	builder := c.Delete().Where(revision.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RevisionDeleteOne{builder}
}

// Query returns a query builder for Revision.
func (c *RevisionClient) Query() *RevisionQuery {
	return &RevisionQuery{
		config: c.config,
	}
}

// Get returns a Revision entity by its id.
func (c *RevisionClient) Get(ctx context.Context, id int) (*Revision, error) {
	return c.Query().Where(revision.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RevisionClient) GetX(ctx context.Context, id int) *Revision {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RevisionClient) Hooks() []Hook {
	return c.hooks.Revision
}
