## 限流演示

### 一、安装redis服务器

```sh
docker run -d --name redis-container -p 6379:6379 --restart unless-stopped redis redis-server --requirepass redisPassword
```

这里使用docker安装