package service

import (
	"encoding/json"
	"fmt"
	"log"

	model "code.huawink.com/beiwanglu/mq-server/model"
)

//从mq中接收信息，写入数据库
func CreateTask() {
	//mq数据通道
	ch, err := model.MQ.Channel()
	fmt.Println("接受到了createTask的请求mq")
	if err != nil {
		panic(err)
	}
	//从task_queue通道中获取消息
	/* alt+shift+a,块注释 */
	/*
	   name：队列名称;
	   durable：是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;
	   autoDelete：是否自动删除;
	   noWait：是否非阻塞,true为是,不等待RMQ返回信息;
	   args：参数,传nil即可;
	   exclusive：是否设置排他
	*/
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	/*
		设置Qos，设置预取大小prefetch，
		当prefetch=1时，表示在没收到consumer的ACK消息之前，只会为其consumer分派一个消息。
	*/
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}
	// 读出数据
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	//处于一个监听状态，一致监听我们的生产端的生产，所以我们要阻塞主进程
	go func() {
		//将通道中的数据反序列化，然后在数据库中创建
		fmt.Println("接受到了createTask的请求,go func")
		for d := range msgs {
			var t model.Task
			//Unmarshal函数解析json编码的数据并将结果存入v指向的值。
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()
}
