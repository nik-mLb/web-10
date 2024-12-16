package usecase

type Provider interface {
	SelectQuery(msg string) (bool, error)
	InsertQuery(msg string) error
}