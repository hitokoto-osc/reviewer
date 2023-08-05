package model

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser = UserPattern
