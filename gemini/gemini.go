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

type Gemini struct {
	*Config
}

// GeminiOption is a function that configures a Gemini instance
type GeminiOption func(*Gemini) error

// WithGeminiConfig configures a Gemini instance with the specified Config
func WithGeminiConfig(cfg *Config) GeminiOption {
	return func(g *Gemini) error {
		g.Config = cfg
		if err := g.Config.Validate(); err != nil {
			return err
		}
		return nil
	}
}

// NewGemini creates a new Gemini instance
func NewGemini(opts ...GeminiOption) (*Gemini, error) {
	gemini := &Gemini{
		Config: &Config{},
	}
	for _, opt := range opts {
		if err := opt(gemini); err != nil {
			return gemini, err
		}
	}
	return gemini, nil
}

// Run runs the Gemini instance
func (g *Gemini) Run() error {
	return nil
}
