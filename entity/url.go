package entity

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
	idLength = 11
)

type URL struct {
	ID       string
	URL      string
	Expired  *int64
	HitCount int
}

func MakeURL(number uint64, url string, expired *int64) (new *URL) {
	return &URL{
		ID:      encode(number),
		URL:     url,
		Expired: expired,
	}
}

func encode(number uint64) (id string) {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(idLength)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}
