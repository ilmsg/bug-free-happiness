package storage

type Storage interface {
	Set(string, string) error
	Get(string) (string, error)
}
