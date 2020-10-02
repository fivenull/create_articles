package conf

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type: "UNSET",
	//Type:     "mysql",
	DBFile:   "mwxyarticles.db",
	Port:     3306,
	Host:     "192.168.10.10",
	User:     "homestead",
	Password: "secret",
	Name:     "mwxy_article",
}

// SystemConfig 系统公用配置
var SystemConfig = &system{
	Debug:  false,
	Mode:   "master",
	Listen: ":517",
}

// CORSConfig 跨域配置
var CORSConfig = &cors{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS", "DELETE"},
	AllowHeaders:     []string{"Cookie", "X-Policy", "Authorization", "Content-Length", "Content-Type", "X-Path", "X-FileName"},
	AllowCredentials: true,
	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
}
