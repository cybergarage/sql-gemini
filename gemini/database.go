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

const (
	// MySQL is the MySQL database type
	MySQL = "mysql"
	// PostgreSQL is the PostgreSQL database type
	PostgreSQL = "postgres"
)

// Database is the configuration for the database
type Database struct {
	Host  string
	Type  string
	Image string
	Port  int
}

// NewDatabase creates a new database configuration
func NewDatabase() Database {
	return Database{
		Host:  "",
		Type:  "",
		Image: "",
		Port:  0,
	}
}
