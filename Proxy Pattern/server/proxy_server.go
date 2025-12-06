package server

type ProxyServer struct {
	realServer *RealServer
}

func NewProxyServer() *ProxyServer {
	return &ProxyServer{
		realServer: &RealServer{},
	}
}

func (ps *ProxyServer) ServeRequest(user string) string {
	if ps.authenticate(user) {
		return ps.realServer.ServeRequest(user)
	}
	return "Access Denied"
}

func (ps *ProxyServer) authenticate(user string) bool {
	authorizedUsers := map[string]bool{
		"admin": true,
		"user1": true,
	}
	return authorizedUsers[user]
}
