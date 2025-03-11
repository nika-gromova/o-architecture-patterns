package models

type Rule struct {
	Name              string
	Owner             *Owner
	BaseLink          *Link
	DefaultRedirectTo *Link
	Redirections      []*Redirection
}

type Owner struct {
	UUID string
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
