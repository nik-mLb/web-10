package api

type Usecase interface {
	FetchCount() (int, error)
	SetCounter(msg int) error
}