package config

type MongoSupport struct {
	Host       string `env:"MONGO_SUPPORT_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_SUPPORT_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_SUPPORT_USERNAME" envDefault:""`
	Password   string `env:"MONGO_SUPPORT_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_SUPPORT_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_SUPPORT_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_SUPPORT_QUERY" envDefault:""`
}

type MongoContact struct {
	Collection string `env:"MONGO_CONTACT_COLLECTION" envDefault:"accounts"`
}

type MongoFeedback struct {
	Collection string `env:"MONGO_FEEDBACK_COLLECTION" envDefault:"feedbacks"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}
type App struct {
	DB struct {
		Support  MongoSupport
		Contact  MongoContact
		Feedback MongoFeedback
	}
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Session     Session
	RSA         RSA
	Redis       Redis
	TokenSrv    TokenSrv
}
