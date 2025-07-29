package config

type DefaultProvider struct{}

func (d *DefaultProvider) setValues(c *Config) error {
	c.Host = "localhost:8080"
	c.DatabaseDSN = ""
	c.AccurualSystemAddress = "localhost:8080"
	return nil
}
