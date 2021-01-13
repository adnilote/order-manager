package kafka

import "errors"

var (
	ErrInvalidFieldType = errors.New("Unexpected field type")
	ErrInvalidMsgType   = errors.New("Unexpected message type")
	ErrInvalidMsgLength = errors.New("Unexpected ksql msg length")
)
