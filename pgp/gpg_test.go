package pgp_test

import (
	"bytes"
	"github.com/ProtonMail/gopenpgp/v3/crypto"
	"github.com/lazygophers/utils/pgp"
	"golang.org/x/crypto/openpgp/packet"
	"testing"
)

func TestGpg(t *testing.T) {
	key, err := pgp.Gen("11", "22", "33", &packet.Config{
		DefaultHash:            crypto.SHA512,
		DefaultCipher:          packet.CipherAES256,
		DefaultCompressionAlgo: packet.CompressionZLIB,
		CompressionConfig: &packet.CompressionConfig{
			Level: packet.BestCompression,
		},
		RSABits: 1024,
	})
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	err = pgp.Register(key.PublicKey, key.PrivateKey)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	encrypt, err := pgp.Encrypt([]byte("hello"))
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	//t.Log(string(encrypt))

	msg, err := pgp.Decrypt(encrypt)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	if !bytes.Equal([]byte("hello"), msg) {
		t.Fatalf("want hello, got %s", msg)
	}

	t.Log("success")
}

func TestGpgText(t *testing.T) {
	key, err := pgp.Gen("11", "22", "33", &packet.Config{
		DefaultHash:            crypto.SHA512,
		DefaultCipher:          packet.CipherAES256,
		DefaultCompressionAlgo: packet.CompressionZLIB,
		CompressionConfig: &packet.CompressionConfig{
			Level: packet.BestCompression,
		},
		RSABits: 1024,
	})
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	err = pgp.Register(key.PublicKey, key.PrivateKey)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	encrypt, err := pgp.EncryptText("Tag", []byte("hello"))
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	msg, err := pgp.DecryptText("Tag", encrypt)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	if !bytes.Equal([]byte("hello"), msg) {
		t.Fatalf("want hello, got %s", msg)
	}

	t.Log("success")
}
