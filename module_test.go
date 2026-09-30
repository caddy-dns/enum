package enum

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/enum"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	for _, tc := range []struct {
		config  string
		wantErr bool
	}{
		{"enum {\n api_key {env.ENUM_API_KEY}\n project_id proj-123\n}", false},
		{"enum {\n api_key key\n}", true},
		{"enum {\n project_id proj-123\n}", true},
		{"enum key", true},
		{"enum {\n api_key a\n api_key b\n project_id proj-123\n}", true},
		{"enum {\n token key\n}", true},
	} {
		p := Provider{new(enum.Provider)}

		err := p.UnmarshalCaddyfile(caddyfile.NewTestDispenser(tc.config))
		if (err != nil) != tc.wantErr {
			t.Errorf("%q: got error %v, want error %v", tc.config, err, tc.wantErr)
		}

		if err == nil && (p.APIKey != "{env.ENUM_API_KEY}" || p.ProjectID != "proj-123") {
			t.Errorf("%q: got %+v", tc.config, p.Provider)
		}
	}
}
