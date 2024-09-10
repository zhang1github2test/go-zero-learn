#### 1、使用goctl工具快速生成一个http服务

```sh
goctl api new user
```

#### 2、修改user/user.api文件

```api
	Name string `json:"name,optional"`
	Id   string `json:"id,optional"`
	Age  int    `json:"age,optional"`
```

#### 3、使用user.api生成项目的基本代码

```sh
goctl api go --api .\user\user.api --dir user
```

#### 4、编写业务逻辑代码

internal\logic\userlogic.go中的代码

```go
func (l *UserLogic) User(req *types.UserReq) (resp *types.UserReqResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserReqResp{
		UserReq: *req,
		Status:  "ok",
	}
	return
}
```

这里我们简单的把请求的对象放入到响应体中，并且给了一个状态。

#### 5、测试http服务

```sh
curl -X GET -H "Content-Type: application/json"  -d'{"id": "iidsd","name":"zhangshenglu"}' 192.168.10.11:8888/user/_query
```

响应如下：

```txt
{"name":"zhangshenglu","id":"iidsd","age":0,"status":"ok"}
```

