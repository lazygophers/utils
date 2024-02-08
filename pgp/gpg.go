package pgp

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/jchavannes/go-pgp/pgp"
	"github.com/lazygophers/log"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"io"
)

var EntityList openpgp.EntityList

func GetEntities(pub, priv []byte) (openpgp.EntityList, error) {
	entity, err := pgp.GetEntity(pub, priv)
	if err != nil {
		return nil, err
	}

	return openpgp.EntityList{entity}, nil
}

// Register 注册私钥、公钥
// 如果不注册私钥，无法解密
func Register(pub, priv []byte) (err error) {
	EntityList, err = GetEntities(pub, priv)
	return
}

func AppendEntity(pub, priv []byte) error {
	entity, err := pgp.GetEntity(pub, priv)
	if err != nil {
		return err
	}

	EntityList = append(EntityList, entity)

	return nil
}

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func Gen(fullname string, comment string, email string, config *packet.Config) (*KeyPair, error) {
	var key KeyPair

	e, err := openpgp.NewEntity(fullname, comment, email, config)
	if err != nil {
		return nil, err
	}

	for _, id := range e.Identities {
		err = id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, config)
		if err != nil {
			return nil, err
		}
	}

	buf := log.GetBuffer()
	defer log.PutBuffer(buf)

	w, err := armor.Encode(buf, openpgp.PublicKeyType, nil)
	if err != nil {
		return nil, err
	}
	_ = e.Serialize(w)
	_ = w.Close()

	key.PublicKey = make([]byte, buf.Len())
	copy(key.PublicKey, buf.Bytes())
	buf.Reset()

	w, err = armor.Encode(buf, openpgp.PrivateKeyType, nil)
	if err != nil {
		return nil, err
	}
	_ = e.SerializePrivate(w, nil)
	_ = w.Close()

	key.PrivateKey = make([]byte, buf.Len())
	copy(key.PrivateKey, buf.Bytes())

	return &key, nil
}

func Encrypt(message []byte) ([]byte, error) {
	return EncryptWithEntities(message, EntityList)
}

func EncryptWithEntities(message []byte, entities openpgp.EntityList) ([]byte, error) {
	buf := log.GetBuffer()
	defer log.PutBuffer(buf)

	//encoderWriter, err := armor.Encode(buf, "Message", make(map[string]string))
	//if err != nil {
	//	return []byte{}, fmt.Errorf("error creating OpenPGP armor: %v", err)
	//}

	encryptorWriter, err := openpgp.Encrypt(buf, entities, nil, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating entity for encryption: %v", err)
	}

	compressorWriter := brotli.NewWriterLevel(encryptorWriter, brotli.BestCompression)

	_, err = io.Copy(compressorWriter, bytes.NewReader(message))
	if err != nil {
		return []byte{}, fmt.Errorf("error writing data to compressor: %v", err)
	}

	_ = compressorWriter.Close()
	_ = encryptorWriter.Close()
	//_ = encoderWriter.Close()

	return buf.Bytes(), nil
}

func Decrypt(encrypted []byte) ([]byte, error) {
	return DecryptWithEntities(encrypted, EntityList)
}

func DecryptWithEntities(encrypted []byte, entities openpgp.EntityList) ([]byte, error) {
	//block, err := armor.Decode(bytes.NewReader(encrypted))
	//if err != nil {
	//	return []byte{}, fmt.Errorf("error decoding: %v", err)
	//}
	//if block.Type != "Message" {
	//	return []byte{}, errors.New("invalid message type")
	//}

	messageReader, err := openpgp.ReadMessage(bytes.NewReader(encrypted), entities, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading message: %v", err)
	}

	return io.ReadAll(brotli.NewReader(messageReader.UnverifiedBody))
}

func EncryptText(tag string, message []byte) ([]byte, error) {
	return EncryptTextWithEntities(tag, message, EntityList)
}

func EncryptTextWithEntities(tag string, message []byte, entities openpgp.EntityList) ([]byte, error) {
	buf := log.GetBuffer()
	defer log.PutBuffer(buf)

	encoderWriter, err := armor.Encode(buf, tag, make(map[string]string))
	if err != nil {
		return []byte{}, fmt.Errorf("error creating OpenPGP armor: %v", err)
	}

	encryptorWriter, err := openpgp.Encrypt(encoderWriter, entities, nil, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating entity for encryption: %v", err)
	}

	compressorWriter := brotli.NewWriterLevel(encryptorWriter, brotli.BestCompression)

	_, err = io.Copy(compressorWriter, bytes.NewReader(message))
	if err != nil {
		return []byte{}, fmt.Errorf("error writing data to compressor: %v", err)
	}

	_ = compressorWriter.Close()
	_ = encryptorWriter.Close()
	_ = encoderWriter.Close()

	return buf.Bytes(), nil
}

func DecryptText(tag string, encrypted []byte) ([]byte, error) {
	return DecryptTextWithEntities(tag, encrypted, EntityList)
}

func DecryptTextWithEntities(tag string, encrypted []byte, entities openpgp.EntityList) ([]byte, error) {
	block, err := armor.Decode(bytes.NewReader(encrypted))
	if err != nil {
		return []byte{}, fmt.Errorf("error decoding: %v", err)
	}
	if block.Type != tag {
		return []byte{}, errors.New("invalid message type")
	}

	messageReader, err := openpgp.ReadMessage(block.Body, entities, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading message: %v", err)
	}

	return io.ReadAll(brotli.NewReader(messageReader.UnverifiedBody))
}
