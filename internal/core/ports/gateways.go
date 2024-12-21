package ports

type PubSub interface {
	Pub(message string) error
}
