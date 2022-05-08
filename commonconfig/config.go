package commonconfig

type HttpServerConfig struct {
	HTTPPort       int `json:"http_port" toml:"http_port"`
	ReadTimeout    int `json:"read_timeout" toml:"read_timeout"`
	WriteTimeout   int `json:"write_timeout" toml:"write_timeout"`
	MaxHeaderBytes int `json:"max_header_bytes" toml:"max_header_bytes"`
}

type EtcdConfig struct {
	Endpoints  []string `json:"endpoints" toml:"endpoints"`
	PrefixPath string   `json:"prefix_path" toml:"prefix_path"`
}

type ServerConfig struct {
	HttpServerConfig `json:",inline" toml:",inline"`
	EtcdConfig       `json:"etcd" toml:"etcd"`
	ServiceEnv       map[string]string `json:"sevice_env" toml:"sevice_env"`
	GRPCPort         int               `json:"grpc_port" toml:"grpc_port"`
	HTTPSPort        int               `json:"https_port" toml:"https_port"`
	ProjectName      string            `json:"project_name" toml:"project_name"`
	PromePort        int               `json:"prome_port" toml:"prome_port"`
	PProfPort        int               `json:"pprof_port" toml:"pprof_port"`
	Debug            bool              `json:"debug" toml:"debug"`
}

type MongoClientConfig struct {
	URL         string `json:"url" toml:"url"`
	MaxPoolSize int    `json:"max_pool_size" toml:"max_pool_size"`
	Timeout     int    `json:"timeout" toml:"timeout"`
}

type RedisConfig struct {
	Address            string         `json:"address" toml:"address"`
	Password           string         `json:"password" toml:"password"`
	DefaultDB          int            `json:"default_db" toml:"default_db"`
	DBIndexes          map[string]int `json:"db_indexes" toml:"db_indexes"`
	MinIdleConnections int            `json:"min_idel_conns" toml:"min_idel_conns"`
	ReadTimeout        int            `json:"read_timeout" toml:"read_timeout"`
	DialTimout         int            `json:"dial_timeout" toml:"dial_timeout"`
}

type FinderConfig struct {
	Type      string `json:"type" toml:"type"`
	Timeout   int    `json:"timeout" toml:"timeout"`
	Endpoint  string `json:"endpoint" toml:"endpoint"`
	Region    string `json:"region" toml:"region"`
	AccessKey string `json:"access_key" toml:"access_key"`
	SecretKey string `json:"secret_key" toml:"secret_key"`
}

type DownloaderConfig struct {
	FinderConfig `json:",inline" toml:",inline"`
	SourcePath   string            `json:"source_path" toml:"source_path"`
	LocalPath    string            `json:"local_path" toml:"local_path"`
	Interval     int               `json:"interval" toml:"interval"`
	Extra        map[string]string `json:"extra" toml:"extra"`
}

type GRPCClientConfig struct {
	Url            string `json:"url" toml:"url"`
	DialTimeout    int    `json:"dail_timeout" toml:"dail_timeout"`
	RequestTimeout int    `json:"request_timeout" toml:"request_timeout"`
	InitConn       int    `json:"init_conn" toml:"init_conn"`
	MaxConn        int    `json:"max_conn" toml:"max_conn"`
	HealthCheck    bool   `json:"health_check" toml:"health_check"`
}

type ZkRegConfig struct {
	Hosts  []string `toml:"hosts"`
	Prefix string   `toml:"prefix"`
}

type PanguConfig struct {
	ZKHosts     []string `toml:"zkhosts"`
	ServiceName string   `toml:"service_name"`
}
