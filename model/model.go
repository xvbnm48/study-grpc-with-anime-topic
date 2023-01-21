package model

type Anime struct {
	Name        string
	Description string
	Genre       string
	Studio      string
	Year        string
}

// Configuration is a struct for database configuration
type Configuration struct {
	DBHost     string
	DBName     string
	DBOptions  string
	DBPassword string
	DBPort     string
	DBUser     string
}
