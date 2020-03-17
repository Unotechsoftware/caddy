// Copyright 2015 Light Code Labs, LLC
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

package caddyhttp

import (
	// plug in the server
	_ "github.com/Unotechsoftware/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/Unotechsoftware/caddy/caddyhttp/basicauth"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/bind"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/browse"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/errors"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/expvar"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/extensions"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/fastcgi"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/gzip"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/header"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/index"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/internalsrv"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/limits"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/log"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/markdown"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/mime"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/pprof"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/proxy"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/push"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/redirect"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/requestid"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/rewrite"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/root"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/status"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/templates"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/timeouts"
	_ "github.com/Unotechsoftware/caddy/caddyhttp/websocket"
	_ "github.com/Unotechsoftware/caddy/onevent"
)
