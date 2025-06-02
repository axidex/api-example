package logger

import "go.uber.org/zap"

type Attribute struct {
	Key string
	Val interface{}
}

func NewAttribute(key string, val interface{}) Attribute {
	return Attribute{Key: key, Val: val}
}

func transformAttributes(attribute []Attribute) []zap.Field {
	if len(attribute) == 0 {
		return nil
	}

	var fields []zap.Field
	for _, attr := range attribute {
		fields = append(fields, zap.Any(attr.Key, attr.Val))
	}

	return fields
}
