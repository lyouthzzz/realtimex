package gateway

type ConnectAuthn interface {
	Check(clientID, username, password string) bool
}

type TopicAuthn interface {
	Check(action, clientID, username, ip, topic string) bool
}
