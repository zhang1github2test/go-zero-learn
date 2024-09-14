package main

import (
	"fmt"
	. "go-zero-learn/loadbalancer"
)

func main() {
	// 定义服务器及其权重
	servers := []*Server{
		{Address: "Server1", Weight: 5},
		{Address: "Server2", Weight: 1},
		{Address: "Server3", Weight: 1},
	}

	// 创建加权轮询实例
	wrr := NewWeightedRoundRobin(servers)

	// 模拟多次请求，观察服务器选择情况
	for i := 0; i < 10; i++ {
		server := wrr.Next()
		fmt.Printf("Request %d sent to %s\n", i+1, server.Address)
	}
	fmt.Println("---------------------------------------")
	// 轮询算法

	// 创建轮询调度器实例
	rr := NewRoundRobin(servers)

	// 模拟多次请求，观察服务器选择情况
	for i := 0; i < 10; i++ {
		server := rr.Next()
		fmt.Printf("Request %d sent to %s\n", i+1, server.Address)
	}
}
