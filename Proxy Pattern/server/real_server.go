package server

type RealServer struct{}

func (rs *RealServer) ServeRequest(user string) string {
	return "Request served to " + user
}
