// Package config
package config

import (
	"context"
	"log"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	App           AppConfig           `koanf:"app,omitempty"`
	OpenAPI       OpenAPIConfig       `koanf:"openapi,omitempty"`
	Audit         AuditConfig         `koanf:"audit,omitempty"`
	FeatureFlags  FeatureFlagsConfig  `koanf:"feature_flags,omitempty"`
	Backup        BackupConfig        `koanf:"backup,omitempty"`
	Auth          AuthConfig          `koanf:"auth,omitempty"`
	Session       SessionConfig       `koanf:"session,omitempty"`
	Jwt           JwtConfig           `koanf:"jwt,omitempty"`
	Tenancy       TenancyConfig       `koanf:"tenancy,omitempty"`
	Subscription  SubscriptionConfig  `koanf:"subscription,omitempty"`
	Database      DatabaseConfig      `koanf:"database,omitempty"`
	HTTP          HTTPConfig          `koanf:"http,omitempty"`
	Logging       LoggingConfig       `koanf:"logging,omitempty"`
	Messaging     MessagingConfig     `koanf:"messaging"`
	Observability ObservabilityConfig `koanf:"observability,omitempty"`
	Storage       StorageConfig       `koanf:"storage,omitempty"`
	Integrations  IntegrationsConfig  `koanf:"integrations,omitempty"`
	Jobs          JobsConfig          `koanf:"jobs,omitempty"`
	Modules       ModulesConfig       `koanf:"modules,omitempty"`
	Redis         RedisConfig         `koanf:"redis,omitempty"`
	Queues        QueuesConfig        `koanf:"queues,omitempty"`
	Security      SecurityConfig      `koanf:"security,omitempty"`
	Watermill     WatermillConfig     `koanf:"watermill,omitempty"`
}

func LoadConfig() *Config {
	logger := glog.New()
	ctx := context.TODO()
	k := koanf.New(".")
	for _, path := range []string{
		"configs/base/app.yaml",
		"configs/base/auth.yaml",
		"configs/base/database.yaml",
		"configs/base/http.yaml",
		"configs/base/logging.yaml",
		"configs/base/messaging.yaml",
		"configs/base/observability.yaml",
		"configs/base/storage.yaml",
		"configs/base/integrations.yaml",
		"configs/base/jobs.yaml",
		"configs/base/modules.yaml",
		"configs/base/redis.yaml",
		"configs/base/queues.yaml",
		"configs/base/security.yaml",
		"configs/base/watermill.yaml",
		"configs/base/subscription.yaml",
	} {
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			log.Fatalf("load %s: %s", path, err)
		}
	}
	err := k.Load(file.Provider("configs/environments/development.yaml"), yaml.Parser())
	if err != nil {
		log.Fatalf("load %s: %s", "configs/environments/development.yaml", err)
	}
	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("unmarshal config: %s", err)
	}
	if err := gvalid.New().Data(&cfg).Run(context.Background()); err != nil {
		logger.Error(ctx, "errors count ", len(err.Strings()))
	}

	return &cfg
}

func MustValidateConfig(cfg *Config) {
	if err := gvalid.New().Data(cfg).Run(context.Background()); err != nil {
		log.Fatalf("validate config: %s", err)
	}
}
