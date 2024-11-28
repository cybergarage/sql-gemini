// Copyright (C) 2024 The sql-gemini Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gemini

import (
	"fmt"

	"github.com/cybergarage/go-sqltest/sqltest"
)

// SQLClient is the client for the SQL database
type SQLClient = sqltest.Client

// NewClientFrom creates a new SQL client from the database configuration
func NewClientFrom(cfg *Database) (SQLClient, error) {
	var client SQLClient
	switch cfg.Type {
	case MySQL:
		client = sqltest.NewMySQLClient()
	case PostgreSQL:
		client = sqltest.NewPostgresClient()
	}
	if client == nil {
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
	client.SetHost(cfg.Host)
	return client, nil
}
