package commonconfig

type HttpServerConfig struct {
	HTTPPort       int `toml:"http_port"`
	ReadTimeout    int `toml:"read_timeout"`
	WriteTimeout   int `toml:"write_timeout"`
	MaxHeaderBytes int `toml:"max_header_bytes"`
}

type ServerConfig struct {
	HttpServerConfig `toml:",inline"`
	ServiceEnv       map[string]string `toml:"sevice_env"`
	GRPCPort         int               `toml:"grpc_port"`
	HTTPSPort        int               `toml:"https_port"`
	ProjectName      string            `toml:"project_name"`
	PromePort        int               `toml:"prome_port"`
	PProfPort        int               `toml:"pprof_port"`
	Debug            bool              `toml:"debug"`
}

type MongoClientConfig struct {
	URL         string `toml:"url"`
	MaxPoolSize int    `toml:"max_pool_size"`
	Timeout     int    `toml:"timeout"`
}

type RedisConfig struct {
	Address            string         `toml:"address"`
	Password           string         `toml:"password"`
	DefaultDB          int            `toml:"default_db"`
	DBIndexes          map[string]int `toml:"db_indexes"`
	MinIdleConnections int            `toml:"min_idel_conns"`
	ReadTimeout        int            `toml:"read_timeout"`
	DialTimout         int            `toml:"dial_timeout"`
}
