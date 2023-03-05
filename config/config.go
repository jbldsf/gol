package config

import "os"

var (
	HTTPPort, HTTPSPort int
)

func init() {
	switch os.Getenv("ENV") {
	case "production":
		HTTPPort = 5000
		HTTPSPort = 5001
	case "test":
		HTTPPort = 3000
		HTTPSPort = 3001
	}
}
