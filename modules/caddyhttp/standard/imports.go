package standard

import (
	// standard Caddy HTTP app modules
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/caddyauth"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/encode"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/encode/brotli"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/encode/gzip"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/encode/zstd"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/fileserver"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/headers"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/httpcache"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/requestbody"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/reverseproxy"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/reverseproxy/fastcgi"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/rewrite"
	_ "github.com/Unotechsoftware/caddy/v2/modules/caddyhttp/templates"
)
