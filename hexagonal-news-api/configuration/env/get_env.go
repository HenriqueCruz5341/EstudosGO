package env

import "os"

func GetNewsAPIKey() string {
	return os.Getenv("NEWS_API_KEY")
}

func GetNewsAPIURL() string {
	return os.Getenv("NEWS_API_URL")
}
