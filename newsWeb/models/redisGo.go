package models

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func init(){
	conn,err:=redis.Dial("tcp",":6379")
	if err!=nil{
		fmt.Println("redis链接失败")
		return
	}

	//操作数据
	resp,err:=conn.Do("mget","userName","pCount")

	//回复助手函数 （类型转换）
	//result,err:=redis.String(resp,err)
	result,err:=redis.Values(resp,err)
	if err!=nil{
		fmt.Println("Values err=",err)
		return
	}


	//scan函数
	var userName string
	var pCount int
	//if _, err:=redis.Scan(result,&userName,&pCount){
	//
	//}
	if _, err := redis.Scan(result,&userName,&pCount);
	err != nil {
		fmt.Println("Scan err=",err)
		return
	}

	fmt.Println("userName",userName,"pCount",pCount)
}