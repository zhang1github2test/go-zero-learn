package loadbalancer

type LoadBalancer interface {
	Next() *Server
}

// RoundRobin 调度器结构
type RoundRobin struct {
	servers []*Server
	current int
}

// NewRoundRobin 创建一个新的轮询调度器
func NewRoundRobin(servers []*Server) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		current: -1,
	}
}

// Next 选择下一个服务器
func (rr *RoundRobin) Next() *Server {
	if len(rr.servers) == 0 {
		return nil
	}

	// 移动到下一个服务器
	rr.current = (rr.current + 1) % len(rr.servers)
	return rr.servers[rr.current]
}
