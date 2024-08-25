package config

type Config struct {
	Server *Server
	MySQL  map[string]*MySQL
	Redis  *Redis
	Log    *Log
	Es     *Es
	Image  *Image
	Oss    *Oss
	Jwt    *Jwt
	Email  *Email
}

type Server struct {
	Host       string
	Port       int
	Level      string
	UploadMode string `mapstructure:"upload_mode"`
}

type MySQL struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Redis struct {
	Host     string
	Port     int
	Db       int
	Password string
}

type Log struct {
	Level      string
	Dir        string
	Filename   string
	Format     string
	ShowLine   bool `mapstructure:"show_line"`
	MaxBackups int  `mapstructure:"max_backups"`
	MaxSize    int  `mapstructure:"max_size"`
	MaxAge     int  `mapstructure:"max_age"`
	Compress   bool
}

type Es struct {
	Host     string
	Port     int
	Username string
	Password string
	Sniffer  bool
	Index    string
}

type Image struct {
	AvatarDir  string `mapstructure:"avatar_dir"`
	ProductDir string `mapstructure:"product_dir"`
}

type Oss struct {
	Endpoint        string `mapstructure:"endpoint"`
	Bucket          string
	AccessKeyID     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
}

type Jwt struct {
	Secret          string
	Issuer          string
	AccessTokenTTl  int `mapstructure:"access_token_ttl"`
	RefreshTokenTTl int `mapstructure:"refresh_token_ttl"`
}

type Email struct {
	Host     string
	Port     int
	Username string
	Alias    string
	Password string
}
