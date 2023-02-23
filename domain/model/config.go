package model

type Config struct {
	Server  *ServerType
	Service *ServiceType
}

type ServerType struct {
	Port    int
	Timeout int
}

type ServiceType struct {
	Endpoint string
	Name     string
}
