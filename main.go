package main

import (
	"weserver/routes"
	"weserver/sqls"
)

func main() {
	sqls.RunSql()
	routes.Route()
	//arr:=make([]int,5,6)
	////arr1:=[]int{}
	//for i:=0;i<=7;i++{
	//	arr=append(arr,i)
	//}
	//print(cap(arr))
}
