package compress

import "github.com/golang/snappy"

func Encode(raw []byte) []byte {
	return snappy.Encode(nil, raw)
}

func Decode(raw []byte) ([]byte, error) {
	return snappy.Decode(nil, raw)
}
