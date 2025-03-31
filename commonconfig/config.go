package commonconfig

type HttpServerConfig struct {
	HTTPPort       int `json:"http_port" toml:"http_port" yaml:"http_port"`
	ReadTimeout    int `json:"read_timeout" toml:"read_timeout" yaml:"read_timeout"`
	WriteTimeout   int `json:"write_timeout" toml:"write_timeout" yaml:"write_timeout"`
	MaxHeaderBytes int `json:"max_header_bytes" toml:"max_header_bytes" yaml:"max_header_bytes"`
}

type EtcdConfig struct {
	Endpoints []string `json:"endpoints" toml:"endpoints" yaml:"endpoints"`
	Name      string   `json:"name" toml:"name" yaml:"name"`
}

type RegisterDiscoveryConfig struct {
	EtcdConfig `json:"etcd" toml:"etcd" yaml:"etcd"`
}

type ServerConfig struct {
	HttpServerConfig        `json:",inline" toml:",inline" yaml:",inline"`
	RegisterDiscoveryConfig `json:"register" toml:"register" yaml:"register"`
	ServiceEnv              map[string]string `json:"sevice_env" toml:"sevice_env" yaml:"sevice_env"`
	GRPCPort                int               `json:"grpc_port" toml:"grpc_port" yaml:"grpc_port"`
	HTTPSPort               int               `json:"https_port" toml:"https_port" yaml:"https_port"`
	ProjectName             string            `json:"project_name" toml:"project_name" yaml:"project_name"`
	PromePort               int               `json:"prome_port" toml:"prome_port" yaml:"prome_port"`
	PProfPort               int               `json:"pprof_port" toml:"pprof_port" yaml:"pprof_port"`
	Debug                   bool              `json:"debug" toml:"debug" yaml:"debug"`
}

type MongoClientConfig struct {
	URL         string `json:"url" toml:"url" yaml:"url"`
	MaxPoolSize int    `json:"max_pool_size" toml:"max_pool_size" yaml:"max_pool_size"`
	Timeout     int    `json:"timeout" toml:"timeout" yaml:"timeout"`
}

type RedisConfig struct {
	Address            string         `json:"address" toml:"address" yaml:"address"`
	Password           string         `json:"password" toml:"password" yaml:"password"`
	DefaultDB          int            `json:"default_db" toml:"default_db" yaml:"default_db"`
	DBIndexes          map[string]int `json:"db_indexes" toml:"db_indexes" yaml:"db_indexes"`
	MinIdleConnections int            `json:"min_idel_conns" toml:"min_idel_conns" yaml:"min_idel_conns"`
	ReadTimeout        int            `json:"read_timeout" toml:"read_timeout" yaml:"read_timeout"`
	DialTimout         int            `json:"dial_timeout" toml:"dial_timeout" yaml:"dial_timeout"`
}

type FinderConfig struct {
	Type      string `json:"type" toml:"type" yaml:"type"`
	Timeout   int    `json:"timeout" toml:"timeout" yaml:"timeout"`
	Endpoint  string `json:"endpoint" toml:"endpoint" yaml:"endpoint"`
	Region    string `json:"region" toml:"region" yaml:"region"`
	AccessKey string `json:"access_key" toml:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" toml:"secret_key" yaml:"secret_key"`
}

type DownloaderConfig struct {
	FinderConfig `json:",inline" toml:",inline" yaml:",inline"`
	SourcePath   string            `json:"source_path" toml:"source_path" yaml:"source_path"`
	LocalPath    string            `json:"local_path" toml:"local_path" yaml:"local_path"`
	Interval     int               `json:"interval" toml:"interval" yaml:"interval"`
	Extra        map[string]string `json:"extra" toml:"extra" yaml:"extra"`
}

type GRPCClientConfig struct {
	Url            string `json:"url" toml:"url" yaml:"url"`
	DialTimeout    int    `json:"dail_timeout" toml:"dail_timeout" yaml:"dail_timeout"`
	RequestTimeout int    `json:"request_timeout" toml:"request_timeout" yaml:"request_timeout"`
	InitConn       int    `json:"init_conn" toml:"init_conn" yaml:"init_conn"`
	MaxConn        int    `json:"max_conn" toml:"max_conn" yaml:"max_conn"`
	HealthCheck    bool   `json:"health_check" toml:"health_check" yaml:"health_check"`
}
