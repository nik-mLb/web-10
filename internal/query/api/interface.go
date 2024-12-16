package api

type Usecase interface {
	FetchQuery(msg string) (bool, error)
	SetQuery(msg string) error
}