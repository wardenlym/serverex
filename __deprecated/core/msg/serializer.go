package msg

type ProtobufSerializer struct{}

func (ProtobufSerializer) Marshal(interface{}) (*Msg, error) {
	return nil, nil
}
func (ProtobufSerializer) Unmarshal(*Msg, interface{}) error {
	return nil
}
