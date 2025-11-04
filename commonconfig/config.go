package commonconfig

// HttpServerConfig defines configuration for HTTP server settings
type HttpServerConfig struct {
	HTTPPort       int `json:"http_port" toml:"http_port" yaml:"http_port"`
	ReadTimeout    int `json:"read_timeout" toml:"read_timeout" yaml:"read_timeout"`
	WriteTimeout   int `json:"write_timeout" toml:"write_timeout" yaml:"write_timeout"`
	MaxHeaderBytes int `json:"max_header_bytes" toml:"max_header_bytes" yaml:"max_header_bytes"`
}

// EtcdConfig defines configuration for etcd client connection
type EtcdConfig struct {
	Endpoints   []string `json:"endpoints" toml:"endpoints" yaml:"endpoints"`
	Name        string   `json:"name" toml:"name" yaml:"name"`
	Username    string   `json:"username" toml:"username" yaml:"username"`
	Password    string   `json:"password" toml:"password" yaml:"password"`
	DialTimeout int      `json:"dial_timeout" toml:"dial_timeout" yaml:"dial_timeout"`
}

// RegisterDiscoveryConfig defines configuration for service registration and discovery
type RegisterDiscoveryConfig struct {
	EtcdConfig  `json:"etcd" toml:"etcd" yaml:"etcd"`
	ServiceName string `json:"service_name" toml:"service_name" yaml:"service_name"`
	TTL         int    `json:"ttl" toml:"ttl" yaml:"ttl"`
}

// ServerConfig defines the main server configuration combining various components
type ServerConfig struct {
	HttpServerConfig        `json:",inline" toml:",inline" yaml:",inline"`
	RegisterDiscoveryConfig `json:"register" toml:"register" yaml:"register"`
	ServiceEnv              map[string]string `json:"service_env" toml:"service_env" yaml:"service_env"` // Fixed typo: sevice_env -> service_env
	GRPCPort                int               `json:"grpc_port" toml:"grpc_port" yaml:"grpc_port"`
	HTTPSPort               int               `json:"https_port" toml:"https_port" yaml:"https_port"`
	ProjectName             string            `json:"project_name" toml:"project_name" yaml:"project_name"`
	PromePort               int               `json:"prome_port" toml:"prome_port" yaml:"prome_port"`
	PProfPort               int               `json:"pprof_port" toml:"pprof_port" yaml:"pprof_port"`
	Debug                   bool              `json:"debug" toml:"debug" yaml:"debug"`
	ShutdownTimeout         int               `json:"shutdown_timeout" toml:"shutdown_timeout" yaml:"shutdown_timeout"`
	LogLevel                string            `json:"log_level" toml:"log_level" yaml:"log_level"`
	LogDir                  string            `json:"log_dir" toml:"log_dir" yaml:"log_dir"`
}

// MongoClientConfig defines configuration for MongoDB client connection
type MongoClientConfig struct {
	URL            string `json:"url" toml:"url" yaml:"url"`
	MaxPoolSize    int    `json:"max_pool_size" toml:"max_pool_size" yaml:"max_pool_size"`
	Timeout        int    `json:"timeout" toml:"timeout" yaml:"timeout"`
	MinPoolSize    int    `json:"min_pool_size" toml:"min_pool_size" yaml:"min_pool_size"`
	ConnectTimeout int    `json:"connect_timeout" toml:"connect_timeout" yaml:"connect_timeout"`
	Database       string `json:"database" toml:"database" yaml:"database"`
}

// RedisConfig defines configuration for Redis client connection
type RedisConfig struct {
	Address            string         `json:"address" toml:"address" yaml:"address"`
	Password           string         `json:"password" toml:"password" yaml:"password"`
	DefaultDB          int            `json:"default_db" toml:"default_db" yaml:"default_db"`
	DBIndexes          map[string]int `json:"db_indexes" toml:"db_indexes" yaml:"db_indexes"`
	MinIdleConnections int            `json:"min_idle_conns" toml:"min_idle_conns" yaml:"min_idle_conns"`
	ReadTimeout        int            `json:"read_timeout" toml:"read_timeout" yaml:"read_timeout"`
	DialTimeout        int            `json:"dial_timeout" toml:"dial_timeout" yaml:"dial_timeout"`
	PoolSize           int            `json:"pool_size" toml:"pool_size" yaml:"pool_size"`
	PoolTimeout        int            `json:"pool_timeout" toml:"pool_timeout" yaml:"pool_timeout"`
	IdleTimeout        int            `json:"idle_timeout" toml:"idle_timeout" yaml:"idle_timeout"`
}

// GRPCClientConfig defines configuration for gRPC client connection
type GRPCClientConfig struct {
	Url              string `json:"url" toml:"url" yaml:"url"`
	DialTimeout      int    `json:"dial_timeout" toml:"dial_timeout" yaml:"dial_timeout"` // Fixed typo: dail_timeout -> dial_timeout
	RequestTimeout   int    `json:"request_timeout" toml:"request_timeout" yaml:"request_timeout"`
	InitConn         int    `json:"init_conn" toml:"init_conn" yaml:"init_conn"`
	MaxConn          int    `json:"max_conn" toml:"max_conn" yaml:"max_conn"`
	HealthCheck      bool   `json:"health_check" toml:"health_check" yaml:"health_check"`
	KeepAliveTime    int    `json:"keep_alive_time" toml:"keep_alive_time" yaml:"keep_alive_time"`
	KeepAliveTimeout int    `json:"keep_alive_timeout" toml:"keep_alive_timeout" yaml:"keep_alive_timeout"`
	UseTLS           bool   `json:"use_tls" toml:"use_tls" yaml:"use_tls"`
}

// DatabaseConfig defines common database configuration (optional composite type)
type DatabaseConfig struct {
	MongoConfig  *MongoClientConfig `json:"mongo,omitempty" toml:"mongo,omitempty" yaml:"mongo,omitempty"`
	RedisConfig  *RedisConfig       `json:"redis,omitempty" toml:"redis,omitempty" yaml:"redis,omitempty"`
	MaxOpenConns int                `json:"max_open_conns" toml:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns int                `json:"max_idle_conns" toml:"max_idle_conns" yaml:"max_idle_conns"`
}

// ClientConfig defines composite configuration for various clients (optional)
type ClientConfig struct {
	GRPCClient *GRPCClientConfig  `json:"grpc,omitempty" toml:"grpc,omitempty" yaml:"grpc,omitempty"`
	Redis      *RedisConfig       `json:"redis,omitempty" toml:"redis,omitempty" yaml:"redis,omitempty"`
	Mongo      *MongoClientConfig `json:"mongo,omitempty" toml:"mongo,omitempty" yaml:"mongo,omitempty"`
}
