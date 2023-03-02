package query

import (
	"github.com/go-asphyxia/core/bytes"
)

type (
	Query []Argument

	Argument struct {
		Key, Value bytes.Buffer
	}
)

func Decode(source bytes.Buffer) Query {
	l := len(source)

	if l < 0 {
		return nil
	}

	query := make([]Argument, 0, (l >> 4))

	key := true

	keyStart := 0
	keyEnd := 0

	valueStart := 0

	for i, b := range source {
		if b == '=' && key != false {
			key = false
			keyEnd = i
			valueStart = i + 1
		} else if b == '&' && key == false {
			key = true
			query = append(query, Argument{Key: source[keyStart:keyEnd], Value: source[valueStart:i]})
			keyStart = i + 1
		}
	}

	if key == false {
		query = append(query, Argument{Key: source[keyStart:keyEnd], Value: source[valueStart:l]})
	}

	return query
}

func Encode(query Query) bytes.Buffer {
	i := len(query) - 1

	if i < 0 {
		return nil
	}

	buffer := make(bytes.Buffer, 0, (i << 4))

start:
	argument := query[i]

	buffer.Write(argument.Key)
	buffer.WriteByte('=')
	buffer.Write(argument.Value)

	i -= 1

	if i >= 0 {
		buffer.WriteByte('&')
		goto start
	}

	return buffer
}
