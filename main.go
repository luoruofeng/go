package main

import "go_demo/qps"

// import "go_demo/router_demo"

func main() {
	// boardcaster_mongo.ServerApp()
	// boardcaster_mongo.ClientApp()
	// router_demo.Web()
	// middleware.SimpleApp()
	// middleware.ChiApp()
	// middleware.GinApp()
	// db.MysqlApp()
	// db.GormApp()
	qps.QpsApp()
}
