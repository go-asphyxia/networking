package HTTP

type (
	URI struct {
		Original []byte

		Scheme string

		Username string
		Password string

		Host string

		Path  string
		Query ArgumentList
		Hash  string
	}
)

// type (
// 	Header [2]string

// 	HeaderList struct {
// 		List []Header
// 	}
// )

// type (
// 	Request struct {
// 		Method string
// 		URI
// 		HeaderList HeaderList
// 	}
// )

// type (
// 	Response struct{}
// )
