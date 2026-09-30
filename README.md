enum module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [enum](https://enum.co).

## Caddy module name

```
dns.providers.enum
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "enum",
				"api_key": "{env.ENUM_API_KEY}",
				"project_id": "{env.ENUM_PROJECT_ID}"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns enum {
		api_key {env.ENUM_API_KEY}
		project_id {env.ENUM_PROJECT_ID}
	}
}
```

```
# one site
tls {
	dns enum {
		api_key {env.ENUM_API_KEY}
		project_id {env.ENUM_PROJECT_ID}
	}
}
```

`api_key` is the API key of a service account, `project_id` the ID of the project that owns the zone. See [libdns/enum](https://github.com/libdns/enum) for caveats.
