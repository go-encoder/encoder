package encoder

import (
	"encoding/base64"
	"github.com/go-encoder/encoder/types"
	"testing"
)

func TestNew(t *testing.T) {
	data := "hello world"

	encoder := New(types.Scrypt)

	hash, err := encoder.Encode(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	salt, err := encoder.GetSalt()
	if err != nil {
		t.Error(err.Error())
		return
	}

	verify, err := encoder.Verify(hash, data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("hash: ", hash)
	t.Log("verify: ", verify)
	t.Log("salt: ", base64.StdEncoding.EncodeToString(salt))
}
