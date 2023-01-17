package query

import (
	"unsafe"

	"github.com/go-asphyxia/core/bytes"
)

type (
	List [][2][]byte
)

func Decode(source []byte) List {
	l := len(source)

	if l < 0 {
		return nil
	}

	list := make([][2][]byte, 0, (l >> 4))

	key := true

	keyStart := 0
	keyEnd := 0

	valueStart := 0

	i := 0

start:
	b := source[i]

	if b == '=' && key != false && keyStart <= i {
		key = false

		keyEnd = i
		valueStart = i + 1
	} else if b == '&' && key == false && valueStart <= i {
		key = true

		list = append(list, [2][]byte{source[keyStart:keyEnd], source[valueStart:i]})
		keyStart = i + 1
	}

	i += 1

	if i < l {
		goto start
	}

	if key == false && valueStart <= l {
		list = append(list, [2][]byte{source[keyStart:keyEnd], source[valueStart:i]})
	}

	return list
}

func DecodeString(source string) List {
	return Decode(unsafe.Slice(unsafe.StringData(source), len(source)))
}

func (list *List) Reset() {
	*list = (*list)[:0]
}

func (list List) Encode() []byte {
	l := len(list)

	if l == 0 {
		return nil
	}

	buffer := &bytes.Buffer{List: make([]byte, 0, (l << 4))}

	i := 0

start:
	argument := list[i]

	buffer.Write(argument[0])
	buffer.WriteByte('=')
	buffer.Write(argument[1])

	i += 1

	if i < l {
		buffer.WriteByte('&')

		goto start
	}

	return buffer.List
}

func (list List) EncodeString() string {
	encoded := list.Encode()

	return unsafe.String(unsafe.SliceData(encoded), len(encoded))
}
