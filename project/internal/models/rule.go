package models

type Rule struct {
	Name              string
	Owner             *User
	BaseLink          *Link
	DefaultRedirectTo *Link
	Redirections      []*Redirection
}

type User struct {
	UUID string
}

type Link struct {
	URL string
}

func (l *Link) Equals(another *Link) bool {
	return l.URL == another.URL
}

type Redirection struct {
	Formula *Formula
	Target  *Link
}

type Formula struct {
	Expression string
}
