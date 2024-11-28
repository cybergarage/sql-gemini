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

package main

import (
	"os"
	"time"

	"github.com/cybergarage/sql-gemini/gemini"
	"github.com/urfave/cli/v2"
)

const (
	ProgramName = "sql-gemini"
)

func main() {
	config := gemini.Config{
		Oracle: gemini.Database{},
	}

	app := &cli.App{
		Name:     ProgramName,
		Version:  gemini.Version,
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "oracle-host",
				Value:       "localhost",
				Usage:       "Oracle database host",
				Destination: &config.Oracle.Host,
			},
			&cli.StringFlag{
				Name:        "oracle-type",
				Usage:       "Oracle database type",
				Destination: &config.Oracle.Type,
			},
			&cli.StringFlag{
				Name:        "oracle-image",
				Usage:       "Oracle database docker image",
				Destination: &config.Oracle.Image,
			},
		},
		Action: func(cCtx *cli.Context) error {
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
