package echo

type EchoConfig struct {
	Port string
}

func NewEchoConfig(port string) EchoConfig {
	return EchoConfig{
		Port: port,
	}
}
