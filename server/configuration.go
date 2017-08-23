package server

const (
	API_ENDPOINT   = "/api/v1"
	SERVER_ADDRESS = ":8080"
)

type Configuration struct {
	ApiEndpoint string
	ServerAddress string
}