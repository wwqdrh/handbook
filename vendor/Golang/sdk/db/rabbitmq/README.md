> 三方库使用github.com/rabbitmq/amqp091-go

示例代码参考: https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.7.9-management-alpine 
```

创建完之后可以创建新用户，这里为了简单直接使用administrator用户

创建virual host -> handbook

## 组件

### 交换机

默认交换机、直连交换机、扇形交换机、主题交换机、头交换机

## 消息模式

- simple(简单工作队列), 使用空交换机，一条消息由一个消费者进行消费。（当有多个消费者时，默认使用轮训机制把消息分配给消费者）
- publish/subscribe，使用扇形交换机， 一条消息由注册的多个消费者使用
- routing，使用直连交换机
- topic(按照规则匹配)，使用主题交换机