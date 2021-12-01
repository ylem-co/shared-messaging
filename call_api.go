package messaging

type CallApi struct {
	Task
	URL              string
	Method           string
	Payload          []byte
	ContentType      string
	AttachedFileName string
}
