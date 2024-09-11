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
		log.Fatal("数据库连接失败", err)
	}
	log.Println(appinfo)

}
