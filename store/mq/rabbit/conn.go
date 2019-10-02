package rabbit

import (
	"api-framework/core/config"
	"api-framework/core/logging"
	"github.com/streadway/amqp"
)


var conn *amqp.Connection
var channel *amqp.Channel

//如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

//更新mq host
func UpdateRabbitHost(host string) {
	config.RabbitConfig.RabbitURL = host
}

//初始化MQ连接信息
func Init() {

	if initChannel(config.RabbitConfig.RabbitURL) {
		channel.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				logging.Error(msg)
				initChannel(config.RabbitConfig.RabbitURL)
			}
		}
	}()
}


func initChannel(rabbitHost string) bool {
	if channel != nil {
		return true
	}

	conn, err := amqp.Dial(rabbitHost)
	if err != nil {
		logging.Info(err.Error())
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		logging.Info(err.Error())
		return false
	}
	return true
}

