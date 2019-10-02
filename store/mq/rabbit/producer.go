package rabbit

import (
	"api-framework/core/config"
	"github.com/streadway/amqp"
)

// 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel(config.RabbitConfig.RabbitURL) {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, 	// 如果没有对应的queue, 就会丢弃这条消息
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}
