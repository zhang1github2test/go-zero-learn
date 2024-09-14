package loadbalancer

// Server 表示一个服务器实例
type Server struct {
	Address string
	Weight  int
	// 当前权重
	CurrentWeight int
	// 初始权重（用于重置）
	EffectiveWeight int
}

// WeightedRoundRobin 调度器结构
type WeightedRoundRobin struct {
	servers []*Server
}

// NewWeightedRoundRobin 创建一个新的加权轮询实例
func NewWeightedRoundRobin(servers []*Server) *WeightedRoundRobin {
	for _, server := range servers {
		server.CurrentWeight = 0
		server.EffectiveWeight = server.Weight
	}
	return &WeightedRoundRobin{
		servers: servers,
	}
}

// Next 选择下一个服务器
func (wrr *WeightedRoundRobin) Next() *Server {
	totalWeight := 0
	var selectedServer *Server

	for _, server := range wrr.servers {
		// 累加每个服务器的权重
		totalWeight += server.EffectiveWeight

		// 选择当前权重最大的服务器
		server.CurrentWeight += server.EffectiveWeight
		if selectedServer == nil || server.CurrentWeight > selectedServer.CurrentWeight {
			selectedServer = server
		}
	}

	if selectedServer != nil {
		// 减去总权重
		selectedServer.CurrentWeight -= totalWeight
	}

	return selectedServer
}
