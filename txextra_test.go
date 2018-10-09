package moneroutil

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func extraFromHex(extraHex string) (*TransactionExtra, error) {
	buf, err := hex.DecodeString(extraHex)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(buf)

	txExtra, err := ParseTransactionExtra(reader)
	return txExtra, err
}

func TestTxPubKey(t *testing.T) {
	txExtra, err := extraFromHex("01433bece9dacbe62c3edc98719fe8025bb28167af38118e0745e1c659066f7225")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(txExtra.PubKeys))
	assert.Equal(t, "433bece9dacbe62c3edc98719fe8025bb28167af38118e0745e1c659066f7225", txExtra.PubKeys[0].String())
}

func TestExtraNonce(t *testing.T) {
	txExtra, err := extraFromHex("02210097a20b5049454816847238e3578238907741846a5de3484a90d2b81f62bda714017874131b1189bfd20dd004444f4b91c807b119ed50a4051a3049ed09d1364f75")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(txExtra.PubKeys))
	assert.Equal(t, "7874131b1189bfd20dd004444f4b91c807b119ed50a4051a3049ed09d1364f75", txExtra.PubKeys[0].String())
}

//TODO: test extras with additional public keys, etc...