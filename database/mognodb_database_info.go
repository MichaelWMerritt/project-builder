package database

type MongoDBDatabaseInfo struct {

	Database			string							`json:"database"`
	ServerAddresses		[]MongoDBDatabaseServerAddress	`json:"serverAddresses"`
	ConnectionsPerHost	int								`json:"connectionsPerHost"`
	WriteConcern		string							`json:"writeConcern"`

}
