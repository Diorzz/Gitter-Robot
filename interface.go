package main

type PostMsg interface {
	GetUrl() string
	GetHeader() string
	GetBody(qes string) string
}
