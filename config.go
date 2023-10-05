package neo4j

type config struct {
	Address               string
	Username              string
	Password              string
	MaxConnectionPoolSize int
	Tls                   bool
}
