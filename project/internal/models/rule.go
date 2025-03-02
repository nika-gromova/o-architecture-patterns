package models

type Rule struct {
	Name         string
	Owner        Owner
	BaseLink     Link
	DefaultLink  *Link
	Redirections []*Redirection
}

type Owner struct {
	Username string
}

type Link struct {
	URL string
}

type Redirection struct {
	Formula Formula
	Target  Link
}

type Formula struct {
	Expression string
}
