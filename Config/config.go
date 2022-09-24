package Config

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	ConnectionString string
}

type KafkaConfiguration struct {
	URL   string
	Topic string
}

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Kafka    KafkaConfiguration
}
