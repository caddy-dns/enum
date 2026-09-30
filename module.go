package enum

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/enum"
)

// Provider lets Caddy read and manipulate DNS records hosted by enum.
type Provider struct{ *enum.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.enum",
		New: func() caddy.Module { return &Provider{new(enum.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()

	p.Provider.APIKey = repl.ReplaceAll(p.Provider.APIKey, "")
	p.Provider.ProjectID = repl.ReplaceAll(p.Provider.ProjectID, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	enum {
//	    api_key <api_key>
//	    project_id <project_id>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			var field *string

			switch d.Val() {
			case "api_key":
				field = &p.Provider.APIKey
			case "project_id":
				field = &p.Provider.ProjectID
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}

			if *field != "" {
				return d.Errf("%s already set", d.Val())
			}

			if !d.NextArg() {
				return d.ArgErr()
			}

			*field = d.Val()

			if d.NextArg() {
				return d.ArgErr()
			}
		}
	}

	if p.Provider.APIKey == "" {
		return d.Err("missing api_key")
	}

	if p.Provider.ProjectID == "" {
		return d.Err("missing project_id")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
