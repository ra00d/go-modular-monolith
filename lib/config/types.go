package config

type AppConfig struct {
	Name     string `koanf:"name"     v:"required|length:1,64"`
	Env      string `koanf:"env"      v:"required|in:development,staging,production"`
	Version  string `koanf:"version"  v:"required|length:1,32"`
	Timezone string `koanf:"timezone" v:"required|length:1,64"`
	Debug    bool   `koanf:"debug"`
}

type AuditConfig struct {
	Enabled       bool `koanf:"enabled"`
	RetentionDays int  `koanf:"retention_days" v:"min:0|max:3650"`
}

type FeatureFlagsConfig struct {
	Inventory     bool `koanf:"inventory"`
	CRM           bool `koanf:"crm"`
	Accounting    bool `koanf:"accounting"`
	Manufacturing bool `koanf:"manufacturing"`
	HR            bool `koanf:"hr"`
}

type BackupConfig struct {
	Enabled  bool `koanf:"enabled"`
	Database struct {
		Cron          string `koanf:"cron" v:"required|length:1,64"`
		RetentionDays int    `koanf:"retention_days" v:"min:0|max:3650"`
	} `koanf:"database"`
	Storage struct {
		Cron          string `koanf:"cron" v:"required|length:1,64"`
		RetentionDays int    `koanf:"retention_days" v:"min:0|max:3650"`
	} `koanf:"storage"`
}

type SessionConfig struct {
	Provider   string `koanf:"provider"    v:"required|in:redis,memory,cookie"`
	CookieName string `koanf:"cookie_name" v:"required|length:1,128"`
	MaxAge     string `koanf:"max_age"     v:"required|length:1,32"`
	Secure     bool   `koanf:"secure"`
	HTTPOnly   bool   `koanf:"http_only"`
	SameSite   string `koanf:"same_site"   v:"required|in:lax,strict,none"`
}
type JwtConfig struct {
	SecretKey string `koanf:"secret_key"`
	Issure    string `koanf:"issure"`
	Ttl       int    `koanf:"ttl"`
}
type PasswordPolicyConfig struct {
	MinLength          int  `koanf:"min_length"           v:"min:6|max:128"`
	RequireUppercase   bool `koanf:"require_uppercase"`
	RequireLowercase   bool `koanf:"require_lowercase"`
	RequireNumber      bool `koanf:"require_number"`
	RequireSpecialChar bool `koanf:"require_special_char"`
}

type AuthConfig struct {
	LoginPath      string               `koanf:"login_path"      v:"required|length:1,128"`
	LogoutPath     string               `koanf:"logout_path"     v:"required|length:1,128"`
	SessionTimeout string               `koanf:"session_timeout" v:"required|length:1,32"`
	Password       PasswordPolicyConfig `koanf:"password"`
}

type TenancyConfig struct {
	Enabled bool   `koanf:"enabled"`
	Mode    string `koanf:"mode"    v:"in:single,shared,dedicated"`
}

type SubscriptionConfig struct {
	// DefaultPlanCode is the fallback plan for tenants without a live
	// subscription (lazy assignment at entitlement-resolution time).
	DefaultPlanCode string `koanf:"default_plan_code" v:"required|length:1,64"`
	// EntitlementCacheTTLSeconds bounds the Redis entitlement cache.
	// ponytail: TTL-only invalidation for catalog edits via version bump;
	// event-driven purge when Watermill lands.
	EntitlementCacheTTLSeconds int `koanf:"entitlement_cache_ttl_seconds" v:"min:1|max:86400"`
	// TrialReminderDays is how long before trial end the trial_ending event fires.
	TrialReminderDays int `koanf:"trial_reminder_days" v:"min:0|max:30"`
}

type DatabaseConfig struct {
	Driver          string `koanf:"driver"            v:"required|in:postgres,mysql,sqlite"`
	Host            string `koanf:"host"              v:"required|length:1,255"`
	Port            int    `koanf:"port"              v:"min:1|max:65535"`
	Database        string `koanf:"database"          v:"required|length:1,128"`
	Username        string `koanf:"username"          v:"required|length:1,128"`
	Password        string `koanf:"password"          v:"required|length:1,256"`
	SSLMode         string `koanf:"ssl_mode"          v:"in:disable,require,verify-ca,verify-full"`
	MaxOpenConns    int    `koanf:"max_open_conns"    v:"min:1|max:1000"`
	MaxIdleConns    int    `koanf:"max_idle_conns"    v:"min:0|max:1000"`
	ConnMaxLifetime string `koanf:"conn_max_lifetime" v:"required|length:1,32"`
}
type OpenAPIConfig struct {
	Endpoint string `koanf:"endpoint" v:"required|length:1,128"`
	Path     string `koanf:"path"     v:"required|length:1,128"`
	Auth     bool   `koanf:"auth"`
}
type HTTPConfig struct {
	Host         string `koanf:"host"          v:"required|length:1,255"`
	Port         int    `koanf:"port"          v:"min:1|max:65535"`
	Domain       string `koanf:"domain"        v:"required|length:1,255"`
	ReadTimeout  string `koanf:"read_timeout"  v:"required|length:1,32"`
	WriteTimeout string `koanf:"write_timeout" v:"required|length:1,32"`
	IdleTimeout  string `koanf:"idle_timeout"  v:"required|length:1,32"`
}

type LoggingConfig struct {
	Level  string `koanf:"level"  v:"required|in:debug,info,warn,error"`
	Format string `koanf:"format" v:"required|in:json,text"`
	File   struct {
		Enabled bool   `koanf:"enabled"`
		Path    string `koanf:"path" v:"required|length:1,255"`
	} `koanf:"file"`
}

type MessagingConfig struct {
	Email struct {
		Enabled  bool   `koanf:"enabled"`
		Provider string `koanf:"provider" v:"in:smtp,ses,mailgun,postmark"`
		SMTP     struct {
			Host      string `koanf:"host" v:"required|length:1,255"`
			Port      int    `koanf:"port" v:"min:1|max:65535"`
			Username  string `koanf:"username" v:"required|length:1,255"`
			Password  string `koanf:"password" v:"required|length:1,256"`
			FromName  string `koanf:"from_name" v:"required|length:1,255"`
			FromEmail string `koanf:"from_email" v:"required|email"`
		} `koanf:"smtp"`
	} `koanf:"email,omitempty"`
	SMS struct {
		Enabled  bool   `koanf:"enabled"`
		Provider string `koanf:"provider" v:"in:twilio"`
		Twilio   struct {
			AccountSID string `koanf:"account_sid" v:"length:1,128"`
			AuthToken  string `koanf:"auth_token" v:"length:1,128"`
			FromNumber string `koanf:"from_number" v:"length:1,32"`
		} `koanf:"twilio,omitempty"`
	} `koanf:"sms,omitempty"`
}

type ObservabilityConfig struct {
	Metrics struct {
		Enabled bool   `koanf:"enabled"`
		Path    string `koanf:"path" v:"required|length:1,128"`
	} `koanf:"metrics,omitempty"`
	Tracing struct {
		Enabled  bool   `koanf:"enabled"`
		Exporter string `koanf:"exporter" v:"required|in:otlp,jaeger,zipkin,stdout"`
		Endpoint string `koanf:"endpoint" v:"required|length:1,255"`
	} `koanf:"tracing,omitempty"`
	Profiling struct {
		Enabled bool `koanf:"enabled"`
	} `koanf:"profiling,omitempty"`
}

type StorageConfig struct {
	Default      string             `koanf:"default"         v:"required|in:local,s3"`
	MaxFileSize  int64              `koanf:"max_file_size"   v:"min:0|max:5368709120"`
	KeyPrefix    string             `koanf:"key_prefix"`
	AllowedMime  []string           `koanf:"allowed_mime"`
	AllowedExt   []string           `koanf:"allowed_ext"`
	Overwrite    bool               `koanf:"overwrite"`
	DetectMime   bool               `koanf:"detect_mime"`
	Local        *LocalStorageCfg   `koanf:"local,omitempty" v:"required-if:default,local"`
	S3           *S3StorageCfg      `koanf:"s3,omitempty"    v:"required-if:default,s3"`
}

type LocalStorageCfg struct {
	Path              string `koanf:"path" v:"required|length:1,255"`
	CreateIfMissing   bool   `koanf:"create_if_missing"`
}

type S3StorageCfg struct {
	Enabled       bool   `koanf:"enabled"`
	Endpoint      string `koanf:"endpoint"     v:"required-if:enabled,true|length:0,255"`
	Bucket        string `koanf:"bucket"       v:"required-if:enabled,true|length:1,255"`
	Region        string `koanf:"region"       v:"required-if:enabled,true|length:1,128"`
	AccessKey     string `koanf:"access_key"   v:"required-if:enabled,true|length:1,128"`
	SecretKey     string `koanf:"secret_key"   v:"required-if:enabled,true|length:1,256"`
	SessionToken  string `koanf:"session_token" v:"length:0,512"`
	PathStyle     bool   `koanf:"path_style"`
	DisableTLS    bool   `koanf:"disable_tls"`
}

type IntegrationsConfig struct {
	Enabled bool `koanf:"enabled"`
	Payment *struct {
		Enabled bool `koanf:"enabled"`
		Stripe  struct {
			SecretKey string `koanf:"secret_key" v:"required-if:enabled,true|length:1,256"`
		} `koanf:"stripe" v:"required-if:enabled,true"`
	} `koanf:"payment,omitempty"`
	WhatsApp *struct {
		Enabled bool   `koanf:"enabled"`
		Token   string `koanf:"token" v:"required-if:enabled,true|length:1,256"`
	} `koanf:"whatsapp,omitempty"`
	Webhook *struct {
		Enabled bool `koanf:"enabled"`
	} `koanf:"webhook,omitempty"`
}

type JobsConfig struct {
	Enabled     bool `koanf:"enabled"`
	CleanupLogs struct {
		Cron string `koanf:"cron" v:"required|length:1,64"`
	} `koanf:"cleanup_logs,omitempty"`
	RecalculateInventory struct {
		Cron string `koanf:"cron" v:"required|length:1,64"`
	} `koanf:"recalculate_inventory,omitempty"`
	GenerateReports struct {
		Cron string `koanf:"cron" v:"required|length:1,64"`
	} `koanf:"generate_reports,omitempty"`
}

type ModuleToggle struct {
	Enabled bool `koanf:"enabled"`
}

type ModulesConfig struct {
	Users       ModuleToggle `koanf:"users"`
	Roles       ModuleToggle `koanf:"roles"`
	Permissions ModuleToggle `koanf:"permissions"`
	Customers   ModuleToggle `koanf:"customers"`
	Suppliers   ModuleToggle `koanf:"suppliers"`
	Inventory   ModuleToggle `koanf:"inventory"`
	Products    ModuleToggle `koanf:"products"`
	Sales       ModuleToggle `koanf:"sales"`
	Purchases   ModuleToggle `koanf:"purchases"`
	Accounting  ModuleToggle `koanf:"accounting"`
	HR          ModuleToggle `koanf:"hr"`
	Payroll     ModuleToggle `koanf:"payroll"`
	CRM         ModuleToggle `koanf:"crm"`
	Reports     ModuleToggle `koanf:"reports"`
}

type RedisConfig struct {
	Enabled  bool   `koanf:"enabled"`
	Host     string `koanf:"host"      v:"required|length:1,255"`
	Port     int    `koanf:"port"      v:"min:1|max:65535"`
	UserName string `koanf:"user_name"`
	Password string `koanf:"password"  v:"length:0,256"`
	DB       int    `koanf:"db"        v:"min:0|max:15"`
	PoolSize int    `koanf:"pool_size" v:"min:1|max:1000"`
}

type QueuesConfig struct {
	Enabled    bool   `koanf:"enabled"`
	Workers    int    `koanf:"workers"     v:"min:1|max:1000"`
	RetryCount int    `koanf:"retry_count" v:"min:0|max:100"`
	RetryDelay string `koanf:"retry_delay" v:"required|length:1,32"`
}

type SecurityConfig struct {
	CORS struct {
		Enabled      bool     `koanf:"enabled"`
		AllowOrigins []string `koanf:"allow_origins" v:"required|length:1,100"`
	} `koanf:"cors"`
	RateLimit struct {
		Enabled  bool   `koanf:"enabled"`
		Requests int    `koanf:"requests" v:"min:1|max:1000000"`
		Duration string `koanf:"duration" v:"required|length:1,32"`
	} `koanf:"rate_limit"`
}

type WatermillConfig struct {
	Transport  string `koanf:"transport"   v:"required|in:postgres,memory,kafka,rabbitmq"`
	CommandBus struct {
		Enabled bool `koanf:"enabled"`
	} `koanf:"command_bus"`
	EventBus struct {
		Enabled bool `koanf:"enabled"`
	} `koanf:"event_bus"`
	Router struct {
		CloseTimeout string `koanf:"close_timeout" v:"required|length:1,32"`
	} `koanf:"router"`
	Postgres struct {
		Schema string `koanf:"schema" v:"required|length:1,128"`
	} `koanf:"postgres"`
}
