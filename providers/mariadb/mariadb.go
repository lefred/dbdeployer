package mariadb

import (
	"fmt"
	"strings"

	"github.com/ProxySQL/dbdeployer/providers"
)

const ProviderName = "mariadb"

type MariaDBProvider struct{}

func NewMariaDBProvider() *MariaDBProvider {
	return &MariaDBProvider{}
}

func (p *MariaDBProvider) Name() string { return ProviderName }

func (p *MariaDBProvider) ValidateVersion(version string) error {
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return fmt.Errorf("invalid MariaDB version format: %q (expected X.Y or X.Y.Z)", version)
	}
	return nil
}

func (p *MariaDBProvider) DefaultPorts() providers.PortRange {
	return providers.PortRange{
		BasePort:         3306,
		PortsPerInstance: 1, // main
	}
}

func (p *MariaDBProvider) FindBinary(version string) (string, error) {
	return "", fmt.Errorf("MariaDBProvider.FindBinary: use sandbox package directly (not yet migrated)")
}

func (p *MariaDBProvider) CreateSandbox(config providers.SandboxConfig) (*providers.SandboxInfo, error) {
	return nil, fmt.Errorf("MariaDBProvider.CreateSandbox: use sandbox package directly (not yet migrated)")
}

func (p *MariaDBProvider) StartSandbox(dir string) error {
	return fmt.Errorf("MariaDBProvider.StartSandbox: use sandbox package directly (not yet migrated)")
}

func (p *MariaDBProvider) StopSandbox(dir string) error {
	return fmt.Errorf("MariaDBProvider.StopSandbox: use sandbox package directly (not yet migrated)")
}

func (p *MariaDBProvider) SupportedTopologies() []string {
	return []string{"single", "multiple", "replication"}
}

func (p *MariaDBProvider) CreateReplica(primary providers.SandboxInfo, config providers.SandboxConfig) (*providers.SandboxInfo, error) {
	return nil, providers.ErrNotSupported
}

func Register(reg *providers.Registry) error {
	return reg.Register(NewMariaDBProvider())
}
