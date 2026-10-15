package config

import (
	"fmt"
	"net/url"
)

func (c DatabaseConfig) DSN() (string, error) {
	switch c.Driver {
	case "postgres":
		return buildPostgresDSN(c), nil

	case "mysql":
		return buildMySQLDSN(c), nil

	case "sqlite":
		return c.Database, nil

	default:
		return "", fmt.Errorf("unsupported database driver: %s", c.Driver)
	}
}

func (c DatabaseConfig) DSNWithSchema(schema string) (string, error) {
	switch c.Driver {
	case "postgres":
		return buildPostgresDSNWithSchema(c, schema), nil
	default:
		return "", fmt.Errorf("unsupported database driver: %s", c.Driver)
	}
}

func buildPostgresDSN(c DatabaseConfig) string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}

	q := u.Query()
	q.Set("sslmode", c.SSLMode)
	u.RawQuery = q.Encode()

	return u.String()
}

func buildPostgresDSNWithSchema(c DatabaseConfig, schema string) string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}

	q := u.Query()
	q.Set("sslmode", c.SSLMode)
	q.Set("search_path", schema)
	u.RawQuery = q.Encode()

	return u.String()
}

func buildMySQLDSN(c DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		url.QueryEscape(c.Username),
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database,
	)
}
