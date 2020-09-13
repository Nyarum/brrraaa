package config

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
}

type Config struct {
	Addr     string
	Database Database
}

func NewConfig() (Config, error) {
	return Config{
		Addr: ":1973",
		Database: Database{
			Host:     "localhost",
			User:     "bra",
			Password: "password",
			Name:     "bra",
		},
	}, nil
}
