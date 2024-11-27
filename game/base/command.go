package base

type Command interface {
	Execute() error
}
