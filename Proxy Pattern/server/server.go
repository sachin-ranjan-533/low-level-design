package server

type Server interface {
	ServerRequest(user string) string
}
