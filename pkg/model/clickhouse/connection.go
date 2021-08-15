// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clickhouse

import (
	"context"
	databasesql "database/sql"
	"fmt"
	"time"

	log "github.com/DBKernel/clickhouse-operator/pkg/announcer"

	_ "github.com/mailru/go-clickhouse"
)

type CHConnection struct {
	params *CHConnectionParams
	conn   *databasesql.DB
}

func NewConnection(params *CHConnectionParams) *CHConnection {
	// Do not establish connection immediately, do it in a lazy manner
	return &CHConnection{
		params: params,
	}
}

// connectContext
func (c *CHConnection) connectContext(ctx context.Context) {
	log.V(2).Info("Establishing connection: %s", c.params.GetDSNWithHiddenCredentials())
	dbConnection, err := databasesql.Open("clickhouse", c.params.GetDSN())
	if err != nil {
		log.V(1).A().Error("FAILED Open(%s). Err: %v", c.params.GetDSNWithHiddenCredentials(), err)
		return
	}

	// Ping should be deadlined
	var parentCtx context.Context
	if ctx == nil {
		parentCtx = context.Background()
	} else {
		parentCtx = ctx
	}
	contxt, cancel := context.WithDeadline(parentCtx, time.Now().Add(defaultTimeout))
	defer cancel()

	if err := dbConnection.PingContext(contxt); err != nil {
		log.V(1).A().Error("FAILED Ping(%s). Err: %v", c.params.GetDSNWithHiddenCredentials(), err)
		_ = dbConnection.Close()
		return
	}

	c.conn = dbConnection
}

// ensureConnectedContext
func (c *CHConnection) ensureConnectedContext(ctx context.Context) bool {
	if c.conn != nil {
		log.V(2).F().Info("Already connected: %s", c.params.GetDSNWithHiddenCredentials())
		return true
	}

	c.connectContext(ctx)

	return c.conn != nil
}

// QueryContext runs given sql query on behalf of specified context
func (c *CHConnection) QueryContext(ctx context.Context, sql string) (*Query, error) {
	if len(sql) == 0 {
		return nil, nil
	}

	var parentCtx context.Context
	if ctx == nil {
		parentCtx = context.Background()
	} else {
		parentCtx = ctx
	}
	contxt, cancel := context.WithDeadline(parentCtx, time.Now().Add(c.params.timeout))

	if !c.ensureConnectedContext(contxt) {
		cancel()
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.params.GetDSNWithHiddenCredentials(), sql)
		log.V(1).A().Error(s)
		return nil, fmt.Errorf(s)
	}

	rows, err := c.conn.QueryContext(contxt, sql)
	if err != nil {
		cancel()
		s := fmt.Sprintf("FAILED Query(%s) %v for SQL: %s", c.params.GetDSNWithHiddenCredentials(), err, sql)
		log.V(1).A().Error(s)
		return nil, err
	}

	log.V(2).Info("clickhouse.QueryContext():'%s'", sql)

	return NewQuery(contxt, cancel, rows), nil
}

// Query runs given sql query
func (c *CHConnection) Query(sql string) (*Query, error) {
	return c.QueryContext(nil, sql)
}

// ExecContext runs given sql query
func (c *CHConnection) ExecContext(ctx context.Context, sql string) error {
	if len(sql) == 0 {
		return nil
	}

	var parentCtx context.Context
	if ctx == nil {
		parentCtx = context.Background()
	} else {
		parentCtx = ctx
	}
	contxt, cancel := context.WithDeadline(parentCtx, time.Now().Add(defaultTimeout))
	defer cancel()

	if !c.ensureConnectedContext(contxt) {
		cancel()
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.params.GetDSNWithHiddenCredentials(), sql)
		log.V(1).A().Error(s)
		return fmt.Errorf(s)
	}

	_, err := c.conn.ExecContext(contxt, sql)

	if err != nil {
		cancel()
		log.V(1).A().Error("FAILED Exec(%s) %v for SQL: %s", c.params.GetDSNWithHiddenCredentials(), err, sql)
		return err
	}

	log.V(2).F().Info("\n%s", sql)

	return nil
}
