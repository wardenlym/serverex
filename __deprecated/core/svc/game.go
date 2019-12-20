package svc

type Game interface {
	Service
	// msg.Parser
	// msg.Dispatcher
	MqClient
}
