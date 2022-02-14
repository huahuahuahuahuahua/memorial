package main

import (
	"code.huawink.com/beiwanglu/mq-server/conf"
	services "code.huawink.com/beiwanglu/mq-server/services"
)

func main() {
	conf.Init()
	forever := make(chan bool)
	services.CreateTask()
	<-forever
}
