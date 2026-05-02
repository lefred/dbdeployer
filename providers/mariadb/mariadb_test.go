package mariadb

import (
	"testing"

	"github.com/ProxySQL/dbdeployer/providers"
)

func TestMariaDBProviderName(t *testing.T) {
	p := NewMariaDBProvider()
	if p.Name() != "mariadb" {
		t.Errorf("expected 'mariadb', got %q", p.Name())
	}
}

func TestMariaDBProviderValidateVersion(t *testing.T) {
	p := NewMariaDBProvider()
	tests := []struct {
		version string
		wantErr bool
	}{
		{"12.2.2", false},
		{"11.4.10", false},
		{"10.5", false},
		{"invalid", true},
	}
	for _, tt := range tests {
		err := p.ValidateVersion(tt.version)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateVersion(%q) error = %v, wantErr %v", tt.version, err, tt.wantErr)
		}
	}
}

func TestMariaDBProviderRegister(t *testing.T) {
	reg := providers.NewRegistry()
	if err := Register(reg); err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	p, err := reg.Get("mariadb")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if p.Name() != "mariadb" {
		t.Errorf("expected 'mariadb', got %q", p.Name())
	}
}
