package example

type Something interface {
	DoSomething(something string) string
	DoSomethingElse(somethingElse int) int
}

type Store interface {
	Get(id string) (string, error)
	Set(id string, value string) error
}
