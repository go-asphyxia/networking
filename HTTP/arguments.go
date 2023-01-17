package HTTP

import (
	"unsafe"

	"github.com/go-asphyxia/core/bytes"
)

type (
	ArgumentList struct {
		List [][2][]byte
	}
)

func (arguments *ArgumentList) Reset() {
	arguments.List = arguments.List[:0]
}

func (arguments *ArgumentList) Decode(source []byte) {
	arguments.List = arguments.List[:0]

	key := true

	keyStart := 0
	keyEnd := 0

	valueStart := 0

	l := len(source)

	i := 0
	b := byte(0)

	for i < l {
		b = source[i]

		if b == '=' && key && keyStart <= i {
			key = false

			keyEnd = i
			valueStart = i + 1
		} else if b == '&' && !key && valueStart <= i {
			key = true

			arguments.List = append(arguments.List, [2][]byte{source[keyStart:keyEnd], source[valueStart:i]})
			keyStart = i + 1
		}

		i++
	}

	if !key && valueStart <= l {
		arguments.List = append(arguments.List, [2][]byte{source[keyStart:keyEnd], source[valueStart:i]})
	}
}

func (arguments *ArgumentList) DecodeString(source string) {
	arguments.Decode(unsafe.Slice(unsafe.StringData(source), len(source)))
}

func (arguments *ArgumentList) Encode() []byte {
	l := len(arguments.List)

	buffer := &bytes.Buffer{List: make([]byte, 0, (l << 4))}

	l -= 1
	i := 0

	argument := arguments.List[i]

start:
	if i < l {
		buffer.Write(argument[0])
		buffer.WriteRune('=')
		buffer.Write(argument[1])
		buffer.WriteRune('&')

		i++

		argument = arguments.List[i]

		goto start
	} else if i == l {
		buffer.Write(argument[0])
		buffer.WriteRune('=')
		buffer.Write(argument[1])
	}

	return buffer.List
}

func (AL *ArgumentList) EncodeString() string {
	target := AL.Encode()
	return unsafe.String(unsafe.SliceData(target), len(target))
}
