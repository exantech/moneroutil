package moneroutil

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeThenEncode(t *testing.T) {
	addrs := []string{
		"9xttFdWis9x5BbDj24X2YLHVYB4DTsdic25SHpRFGSvKeLS2tZ1xL3eAgZYLYMqnr213b4jVy7huzQhM7gL7YGG67cTddJg",
		"4AdUndXHHZ6cfufTMvppY6JwXNouMBzSkbLYfpAV5Usx3skxNgYeYTRj5UzqtReoS44qo9mtmXCqY45DJ852K5Jv2684Rge",
	}

	for _, addr := range addrs {
		raw := DecodeMoneroBase58(addr)

		enc := EncodeMoneroBase58(raw)
		assert.Equal(t, addr, enc)
	}
}

func TestByFuzzing(t *testing.T) {
	for i := 0; i < 20000; i++ {
		ln := rand.Intn(300) + 1
		arr := make([]byte, ln)
		rand.Read(arr)

		encoded := EncodeMoneroBase58(arr)
		raw := DecodeMoneroBase58(encoded)
		reproduced := EncodeMoneroBase58(raw)

		assert.Equal(t, encoded, reproduced)
	}
}

func TestIllegalChar(t *testing.T) {
	assert.Nil(t, DecodeMoneroBase58(""))
	assert.Nil(t, DecodeMoneroBase58("abcd#"))
	assert.Nil(t, DecodeMoneroBase58("@abcd"))
	assert.Nil(t, DecodeMoneroBase58("!AAAA"))
	assert.Nil(t, DecodeMoneroBase58("01234O"))
}
