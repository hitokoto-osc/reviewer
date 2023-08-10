package amqp

type sAMQP struct{}

func New() *sAMQP {
	return &sAMQP{}
}

func (*sAMQP) GetConnection() {

}
