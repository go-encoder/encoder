package encoder

import (
	"encoding/base64"
	"gopkg.in/encoder.v1/types"
	"testing"
)

func TestNew(t *testing.T) {
	data := "hello world"

	encoding := New(types.Pbkdf2)
	encoding.Hash(data)
	hash, err := encoding.Encode(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	salt, err := encoding.GetSalt()
	if err != nil {
		t.Error(err.Error())
		return
	}

	verify, err := encoding.Verify(hash, data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("hash: ", hash)
	t.Log("verify: ", verify)
	t.Log("salt: ", base64.StdEncoding.EncodeToString(salt))
}
