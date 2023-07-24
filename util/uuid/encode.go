package uuid

import "errors"

const encodeBase58Map = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"

var decodeBase58Map [256]byte

func base58(f int64) string {

	if f < 58 {
		return string(encodeBase58Map[f])
	}

	b := make([]byte, 0, 11)
	for f >= 58 {
		b = append(b, encodeBase58Map[f%58])
		f /= 58
	}
	b = append(b, encodeBase58Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

var ErrInvalidBase58 = errors.New("invalid base58")

func ParseBase58(b []byte) (int64, error) {

	var id int64

	for i := range b {
		if decodeBase58Map[b[i]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		id = id*58 + int64(decodeBase58Map[b[i]])
	}

	return id, nil
}
