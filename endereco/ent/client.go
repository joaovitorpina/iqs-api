// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"endereco/ent/migrate"

	"endereco/ent/cep"
	"endereco/ent/cidade"
	"endereco/ent/endereco"
	"endereco/ent/estado"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cep is the client for interacting with the Cep builders.
	Cep *CepClient
	// Cidade is the client for interacting with the Cidade builders.
	Cidade *CidadeClient
	// Endereco is the client for interacting with the Endereco builders.
	Endereco *EnderecoClient
	// Estado is the client for interacting with the Estado builders.
	Estado *EstadoClient
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
	c.Cep = NewCepClient(c.config)
	c.Cidade = NewCidadeClient(c.config)
	c.Endereco = NewEnderecoClient(c.config)
	c.Estado = NewEstadoClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		Cep:      NewCepClient(cfg),
		Cidade:   NewCidadeClient(cfg),
		Endereco: NewEnderecoClient(cfg),
		Estado:   NewEstadoClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Cep:      NewCepClient(cfg),
		Cidade:   NewCidadeClient(cfg),
		Endereco: NewEnderecoClient(cfg),
		Estado:   NewEstadoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cep.
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
	c.Cep.Use(hooks...)
	c.Cidade.Use(hooks...)
	c.Endereco.Use(hooks...)
	c.Estado.Use(hooks...)
}

// CepClient is a client for the Cep schema.
type CepClient struct {
	config
}

// NewCepClient returns a client for the Cep from the given config.
func NewCepClient(c config) *CepClient {
	return &CepClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cep.Hooks(f(g(h())))`.
func (c *CepClient) Use(hooks ...Hook) {
	c.hooks.Cep = append(c.hooks.Cep, hooks...)
}

// Create returns a create builder for Cep.
func (c *CepClient) Create() *CepCreate {
	mutation := newCepMutation(c.config, OpCreate)
	return &CepCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cep entities.
func (c *CepClient) CreateBulk(builders ...*CepCreate) *CepCreateBulk {
	return &CepCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cep.
func (c *CepClient) Update() *CepUpdate {
	mutation := newCepMutation(c.config, OpUpdate)
	return &CepUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CepClient) UpdateOne(ce *Cep) *CepUpdateOne {
	mutation := newCepMutation(c.config, OpUpdateOne, withCep(ce))
	return &CepUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CepClient) UpdateOneID(id int32) *CepUpdateOne {
	mutation := newCepMutation(c.config, OpUpdateOne, withCepID(id))
	return &CepUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cep.
func (c *CepClient) Delete() *CepDelete {
	mutation := newCepMutation(c.config, OpDelete)
	return &CepDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CepClient) DeleteOne(ce *Cep) *CepDeleteOne {
	return c.DeleteOneID(ce.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CepClient) DeleteOneID(id int32) *CepDeleteOne {
	builder := c.Delete().Where(cep.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CepDeleteOne{builder}
}

// Query returns a query builder for Cep.
func (c *CepClient) Query() *CepQuery {
	return &CepQuery{
		config: c.config,
	}
}

// Get returns a Cep entity by its id.
func (c *CepClient) Get(ctx context.Context, id int32) (*Cep, error) {
	return c.Query().Where(cep.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CepClient) GetX(ctx context.Context, id int32) *Cep {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCidade queries the cidade edge of a Cep.
func (c *CepClient) QueryCidade(ce *Cep) *CidadeQuery {
	query := &CidadeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ce.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cep.Table, cep.FieldID, id),
			sqlgraph.To(cidade.Table, cidade.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cep.CidadeTable, cep.CidadeColumn),
		)
		fromV = sqlgraph.Neighbors(ce.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEnderecos queries the enderecos edge of a Cep.
func (c *CepClient) QueryEnderecos(ce *Cep) *EnderecoQuery {
	query := &EnderecoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ce.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cep.Table, cep.FieldID, id),
			sqlgraph.To(endereco.Table, endereco.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cep.EnderecosTable, cep.EnderecosColumn),
		)
		fromV = sqlgraph.Neighbors(ce.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CepClient) Hooks() []Hook {
	return c.hooks.Cep
}

// CidadeClient is a client for the Cidade schema.
type CidadeClient struct {
	config
}

// NewCidadeClient returns a client for the Cidade from the given config.
func NewCidadeClient(c config) *CidadeClient {
	return &CidadeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cidade.Hooks(f(g(h())))`.
func (c *CidadeClient) Use(hooks ...Hook) {
	c.hooks.Cidade = append(c.hooks.Cidade, hooks...)
}

// Create returns a create builder for Cidade.
func (c *CidadeClient) Create() *CidadeCreate {
	mutation := newCidadeMutation(c.config, OpCreate)
	return &CidadeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cidade entities.
func (c *CidadeClient) CreateBulk(builders ...*CidadeCreate) *CidadeCreateBulk {
	return &CidadeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cidade.
func (c *CidadeClient) Update() *CidadeUpdate {
	mutation := newCidadeMutation(c.config, OpUpdate)
	return &CidadeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CidadeClient) UpdateOne(ci *Cidade) *CidadeUpdateOne {
	mutation := newCidadeMutation(c.config, OpUpdateOne, withCidade(ci))
	return &CidadeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CidadeClient) UpdateOneID(id int) *CidadeUpdateOne {
	mutation := newCidadeMutation(c.config, OpUpdateOne, withCidadeID(id))
	return &CidadeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cidade.
func (c *CidadeClient) Delete() *CidadeDelete {
	mutation := newCidadeMutation(c.config, OpDelete)
	return &CidadeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CidadeClient) DeleteOne(ci *Cidade) *CidadeDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CidadeClient) DeleteOneID(id int) *CidadeDeleteOne {
	builder := c.Delete().Where(cidade.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CidadeDeleteOne{builder}
}

// Query returns a query builder for Cidade.
func (c *CidadeClient) Query() *CidadeQuery {
	return &CidadeQuery{
		config: c.config,
	}
}

// Get returns a Cidade entity by its id.
func (c *CidadeClient) Get(ctx context.Context, id int) (*Cidade, error) {
	return c.Query().Where(cidade.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CidadeClient) GetX(ctx context.Context, id int) *Cidade {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEstado queries the estado edge of a Cidade.
func (c *CidadeClient) QueryEstado(ci *Cidade) *EstadoQuery {
	query := &EstadoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cidade.Table, cidade.FieldID, id),
			sqlgraph.To(estado.Table, estado.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cidade.EstadoTable, cidade.EstadoColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCeps queries the ceps edge of a Cidade.
func (c *CidadeClient) QueryCeps(ci *Cidade) *CepQuery {
	query := &CepQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cidade.Table, cidade.FieldID, id),
			sqlgraph.To(cep.Table, cep.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cidade.CepsTable, cidade.CepsColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CidadeClient) Hooks() []Hook {
	return c.hooks.Cidade
}

// EnderecoClient is a client for the Endereco schema.
type EnderecoClient struct {
	config
}

// NewEnderecoClient returns a client for the Endereco from the given config.
func NewEnderecoClient(c config) *EnderecoClient {
	return &EnderecoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `endereco.Hooks(f(g(h())))`.
func (c *EnderecoClient) Use(hooks ...Hook) {
	c.hooks.Endereco = append(c.hooks.Endereco, hooks...)
}

// Create returns a create builder for Endereco.
func (c *EnderecoClient) Create() *EnderecoCreate {
	mutation := newEnderecoMutation(c.config, OpCreate)
	return &EnderecoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Endereco entities.
func (c *EnderecoClient) CreateBulk(builders ...*EnderecoCreate) *EnderecoCreateBulk {
	return &EnderecoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Endereco.
func (c *EnderecoClient) Update() *EnderecoUpdate {
	mutation := newEnderecoMutation(c.config, OpUpdate)
	return &EnderecoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EnderecoClient) UpdateOne(e *Endereco) *EnderecoUpdateOne {
	mutation := newEnderecoMutation(c.config, OpUpdateOne, withEndereco(e))
	return &EnderecoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EnderecoClient) UpdateOneID(id int) *EnderecoUpdateOne {
	mutation := newEnderecoMutation(c.config, OpUpdateOne, withEnderecoID(id))
	return &EnderecoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Endereco.
func (c *EnderecoClient) Delete() *EnderecoDelete {
	mutation := newEnderecoMutation(c.config, OpDelete)
	return &EnderecoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EnderecoClient) DeleteOne(e *Endereco) *EnderecoDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EnderecoClient) DeleteOneID(id int) *EnderecoDeleteOne {
	builder := c.Delete().Where(endereco.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EnderecoDeleteOne{builder}
}

// Query returns a query builder for Endereco.
func (c *EnderecoClient) Query() *EnderecoQuery {
	return &EnderecoQuery{
		config: c.config,
	}
}

// Get returns a Endereco entity by its id.
func (c *EnderecoClient) Get(ctx context.Context, id int) (*Endereco, error) {
	return c.Query().Where(endereco.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EnderecoClient) GetX(ctx context.Context, id int) *Endereco {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCep queries the cep edge of a Endereco.
func (c *EnderecoClient) QueryCep(e *Endereco) *CepQuery {
	query := &CepQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(endereco.Table, endereco.FieldID, id),
			sqlgraph.To(cep.Table, cep.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, endereco.CepTable, endereco.CepColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EnderecoClient) Hooks() []Hook {
	return c.hooks.Endereco
}

// EstadoClient is a client for the Estado schema.
type EstadoClient struct {
	config
}

// NewEstadoClient returns a client for the Estado from the given config.
func NewEstadoClient(c config) *EstadoClient {
	return &EstadoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `estado.Hooks(f(g(h())))`.
func (c *EstadoClient) Use(hooks ...Hook) {
	c.hooks.Estado = append(c.hooks.Estado, hooks...)
}

// Create returns a create builder for Estado.
func (c *EstadoClient) Create() *EstadoCreate {
	mutation := newEstadoMutation(c.config, OpCreate)
	return &EstadoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Estado entities.
func (c *EstadoClient) CreateBulk(builders ...*EstadoCreate) *EstadoCreateBulk {
	return &EstadoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Estado.
func (c *EstadoClient) Update() *EstadoUpdate {
	mutation := newEstadoMutation(c.config, OpUpdate)
	return &EstadoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EstadoClient) UpdateOne(e *Estado) *EstadoUpdateOne {
	mutation := newEstadoMutation(c.config, OpUpdateOne, withEstado(e))
	return &EstadoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EstadoClient) UpdateOneID(id int) *EstadoUpdateOne {
	mutation := newEstadoMutation(c.config, OpUpdateOne, withEstadoID(id))
	return &EstadoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Estado.
func (c *EstadoClient) Delete() *EstadoDelete {
	mutation := newEstadoMutation(c.config, OpDelete)
	return &EstadoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EstadoClient) DeleteOne(e *Estado) *EstadoDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EstadoClient) DeleteOneID(id int) *EstadoDeleteOne {
	builder := c.Delete().Where(estado.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EstadoDeleteOne{builder}
}

// Query returns a query builder for Estado.
func (c *EstadoClient) Query() *EstadoQuery {
	return &EstadoQuery{
		config: c.config,
	}
}

// Get returns a Estado entity by its id.
func (c *EstadoClient) Get(ctx context.Context, id int) (*Estado, error) {
	return c.Query().Where(estado.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EstadoClient) GetX(ctx context.Context, id int) *Estado {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCidades queries the cidades edge of a Estado.
func (c *EstadoClient) QueryCidades(e *Estado) *CidadeQuery {
	query := &CidadeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(estado.Table, estado.FieldID, id),
			sqlgraph.To(cidade.Table, cidade.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, estado.CidadesTable, estado.CidadesColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EstadoClient) Hooks() []Hook {
	return c.hooks.Estado
}
