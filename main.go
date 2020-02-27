package main

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/common"
	_ "github.com/go-sql-driver/mysql"
)


func main()  {

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}


