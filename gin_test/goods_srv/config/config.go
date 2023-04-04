package config
type MysqlConfig struct{
	Host string `mapstructure:"host" json:"host"`
	Port int  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"db" json:"db"`
	User string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
//type RedisConfig struct {
//	Host   string `mapstructure:"host" json:"host"`
//	Port   int    `mapstructure:"port" json:"port"`
//	Expire int    `mapstructure:"expire" json:"expire"`
//}
type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Host        string         `mapstructure:"host" json:"host"`
	//Tags        []string       `mapstructure:"tags" json:"tags"`
	Port        int           `mapstructure:"port" json:"port"`
	//UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	//JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	//AliSmsInfo  AliSmsConfig  `mapstructure:"sms" json:"sms"`
	//RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`

}
