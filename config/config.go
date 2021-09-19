package config

type Config struct {
	Authorities map[string]*Authority `yaml:"authorities"`
}

type Authority struct {
	RootFile *string `yaml:"rootFile"`
	ACME *ACME `yaml:"acme"`
}

type ACME struct {
	Server *string `yaml:"server"`
	Email string `yaml:"email"`
	KeyFile string `yaml:"keyFile"`
	Solvers []*ACMESolver `yaml:"solvers"`
}

type ACMESolver struct {
	Patterns []string `yaml:"patterns"`
	DNS *ACMESolverDNS `yaml:"dns"`
}

type ACMESolverDNS struct {
	Endpoint string `yaml:"endpoint"`
	SharedKeyFile string `yaml:"sharedKeyFile"`
	CertificateThumbprint *string `yaml:"certificateThumbprint"`
}