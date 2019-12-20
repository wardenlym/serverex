package identity

type ServerOption struct {
	ListenAddress  string
	DataSourceName string
}

type JWTOption struct {
	SecurityKey string
	Issuer      string
	Audience    string
}

type SMSOption struct {
	AliSMSTemplateCode string
	AliSMSSignName     string
}

type GoogleToken struct {
	ClientId     string
	ClientSecret string
}

type Config struct {
	ServerOption
	JWTOption
	SMSOption
	GoogleToken
}

var default_prod_config = Config{
	ServerOption: ServerOption{
		ListenAddress:  ":6000",
		DataSourceName: "root:root@tcp(127.0.0.1:3306)/identity_prod?charset=utf8&parseTime=True&loc=Local",
	},
	JWTOption: JWTOption{
		SecurityKey: "MegaIdentitySecurityKey",
		Issuer:      "MegaIdentityService",
		Audience:    "ExcaliburClient",
	},
	SMSOption: SMSOption{
		AliSMSTemplateCode: "SMS_139227612",
		AliSMSSignName:     "无尽秘境",
	},

	GoogleToken: GoogleToken{
		ClientId:     "1074639094933-dbitlpstp4464et46hp6b2ns4l6fken7.apps.googleusercontent.com",
		ClientSecret: "l-juUIFpUOfaHfZTdKs39_H9",
	},
}
