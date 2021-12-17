package config

type (
	Database struct {
		Host        string `mapstructure:"host"`
		Port        int    `mapstructure:"port"`
		Name        string `mapstructure:"name"`
		TablePrefix string `mapstructure:"table_prefix"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
	}
)

type conf struct {
	AppName    string   `mapstructure:"app_name"`    //应用标题
	Debug      bool     `mapstructure:"debug"`       //debug
	Listen     string   `mapstructure:"listen"`      //监听ip+端口
	UploadSize int64    `mapstructure:"upload_size"` //上传文件大小
	Database   Database `mapstructure:"database"`    //数据库
}

var Conf conf
