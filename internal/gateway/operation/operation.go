package operation

const (
	Connect = iota + 1
	Connack
	Publish
	Puback
	Pubrec
	Pubrel
	Pubcomp
	Subscribe
	Suback
	Unsubscribe
	Unsuback
	Pingreq
	Pingresp
	Disconnect
)

var status = map[int]string{
	Connect:     "connect",
	Connack:     "connack",
	Publish:     "publish",
	Puback:      "puback",
	Pubrec:      "pubrec",
	Pubrel:      "pubrel",
	Pubcomp:     "pubcomp",
	Subscribe:   "subscribe",
	Suback:      "suback",
	Unsubscribe: "unsubscribe",
	Unsuback:    "unsuback",
	Pingreq:     "pingreq",
	Pingresp:    "pingresp",
	Disconnect:  "disconnect",
}

func Text(op int) string {
	return status[op]
}
