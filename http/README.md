## 一、go-zero生成http服务演示

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

192.168.10.35:9030  root  doris##SZT87341

root :doris##SZT87341@tcp(192.168.10.35:9030)/business

## 一、go-zero连接mysql数据库演示

#### 1、生成dbmodel的相关数据

```sh
goctl model mysql datasource -table="*" --url "root:mysql#SZT123@tcp(120.78.161.145:3309)/mh_verify" --dir dbmodel
```

这里使用goctl model mysql 命令来生成。

**`-table="\*"`**:

- 这个选项指定要从数据库中生成模型的表。
- `*` 表示生成所有数据表的模型代码。可以用特定的表名代替 `*` 来生成单个或多个特定表的模型。

**`--url "root:mysql#SZT123@tcp(120.78.161.145:3309)/mh_verify"`**:

- `--url` 用于指定数据库的连接 URL，包括用户名、密码、主机、端口和数据库名称。
- `root`: 数据库的用户名。
- `mysql#SZT123`: 数据库密码，其中 `mysql#SZT123` 是密码。
- `tcp(120.78.161.145:3309)`: 连接的数据库主机和端口，主机是 `120.78.161.145`，端口是 `3309`。
- `/mh_verify`: 要访问的数据库名称，这里是 `mh_verify`。

**`--dir dbmodel`**:

- `--dir` 指定生成的模型代码文件的存放目录。
- `dbmodel` 表示模型代码会存放在 `dbmodel` 目录下

#### 2、测试生成的代码

```
package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-learn/http/server-start/user/dbmodel"
	"log"
)

func main() {
	sqlConn := sqlx.NewMysql("root:mysql#SZT123@tcp(120.78.161.145:3309)/mh_verify?charset=utf8mb4&parseTime=true")
	appInfoModel := dbmodel.NewAppInfoModel(sqlConn)
	appinfo, err := appInfoModel.FindOne(context.Background(), 3)
	if err != nil {
		log.Fatal("查询失败", err)
	}
	log.Println(appinfo)

}
```

使用findOne的时候，如果没有查询到数据。则会出现报错，否则就能正常返回数据。