package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/wwqdrh/logger"
)

// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const MQURL = "amqp://guest:guest@127.0.0.1:5672/handbook"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！")

	return rabbitmq
}

// Destory 断开channel和connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

////////////////////
// NewRabbitMQSimple
// 简单模式Step 1.创建简单模式下的RabbitMq实例
// 不使用交换机
////////////////////

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式Step 2:简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// ConsumeSimple 使用 goroutine 消费消息
func (r *RabbitMQ) ConsumeSimple(ctx context.Context) chan []byte {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	messages := make(chan []byte)
	// 启用协和处理消息
	go func() {
		for {
			select {
			case d := <-msgs:
				// 实现我们要实现的逻辑函数
				messages <- d.Body
			case <-ctx.Done():
				return
			}
		}
	}()
	return messages
}

////////////////////
// publish/consumer
// 一条消息会被多个消费者尝试监听
// 使用扇形交换机
////////////////////
func NewFanoutPublish(queueName string, messages chan string) {
	// 连接rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.DefaultLogger.Error(err.Error() + " Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	// 创建信道
	ch, err := conn.Channel()
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to open a channel")
	}
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"simplefanout", // 交换机名字
		"fanout",       // 交换机类型，fanout发布订阅模式
		true,           // 是否持久化
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to declare an exchange")
	}
	for body := range messages {
		// 推送消息
		err = ch.Publish(
			"simplefanout", // exchange（交换机名字，跟前面声明对应）
			"",             // 路由参数，fanout类型交换机，自动忽略路由参数，填了也没用。
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "text/plain", // 消息内容类型，这里是普通文本
				Body:        []byte(body), // 消息内容
			})
		if err != nil {
			logger.DefaultLogger.Error(err.Error())
		}

		log.Printf("发送内容 %s", body)
	}
}

func NewFanoutConsumer(queueName string, ctx context.Context) chan []byte {
	messages := make(chan []byte)
	// 连接rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.DefaultLogger.Error(err.Error() + " Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	// 创建信道，通常一个消费者一个
	ch, err := conn.Channel()
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to open a channel")
	}
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"simplefanout", // 交换机名，需要跟消息发送方保持一致
		"fanout",       // 交换机类型
		true,           // 是否持久化
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to declare an exchange")
	}

	// 声明需要操作的队列
	q, err := ch.QueueDeclare(
		"",    // 队列名字，不填则随机生成一个
		false, // 是否持久化队列
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to declare an queue")
	}

	// 队列绑定指定的交换机
	err = ch.QueueBind(
		q.Name,         // 队列名
		"",             // 路由参数，fanout类型交换机，自动忽略路由参数
		"simplefanout", // 交换机名字，需要跟消息发送端定义的交换器保持一致
		false,
		nil)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to bind a queue")
	}

	// 创建消费者
	msgs, err := ch.Consume(
		q.Name, // 引用前面的队列名
		"",     // 消费者名字，不填自动生成一个
		true,   // 自动向队列确认消息已经处理
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error() + " Failed to register a consumer")
	}

	// 循环消费队列中的消息
	go func() {
		for {
			select {
			case d := <-msgs:
				messages <- d.Body
			case <-ctx.Done():
				return
			}
		}
	}()
	return messages
}
