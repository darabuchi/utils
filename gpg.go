package utils

import (
	"bytes"
	"crypto"
	"github.com/darabuchi/log"
	"github.com/jchavannes/go-pgp/pgp"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"math/rand"
	"time"
)

func GpgEncrypt(pubKey string, msg string) ([]byte, error) {
	pubEntity, err := pgp.GetEntity([]byte(pubKey), []byte{})
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	encrypted, err := pgp.Encrypt(pubEntity, []byte(msg))
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	return encrypted, nil
}

func GpgDecrypt(privKey, pubKey string, msg []byte) ([]byte, error) {
	privEntity, err := pgp.GetEntity([]byte(pubKey), []byte(privKey))
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	decrypted, err := pgp.Decrypt(privEntity, msg)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return decrypted, nil
}

func GenGpg() (*pgp.PGPKeyPair, error) {
	var e *openpgp.Entity
	e, err := openpgp.NewEntity("etsuko", "etsuko-check", "etsuko@darabuchi.com", &packet.Config{
		Rand:                   rand.New(rand.NewSource(time.Now().UnixNano())),
		DefaultHash:            crypto.SHA512,
		DefaultCipher:          packet.CipherAES256,
		Time:                   nil,
		DefaultCompressionAlgo: packet.CompressionZLIB,
		CompressionConfig:      &packet.CompressionConfig{Level: 9},
		S2KCount:               4096,
		RSABits:                4096,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	for _, id := range e.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, nil)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}

	buf := new(bytes.Buffer)
	w, err := armor.Encode(buf, openpgp.PublicKeyType, nil)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	defer w.Close()

	err = e.Serialize(w)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	pubKey := buf.String()

	buf = new(bytes.Buffer)
	w, err = armor.Encode(buf, openpgp.PrivateKeyType, nil)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	_ = e.SerializePrivate(w, nil)
	_ = w.Close()

	privateKey := buf.String()

	return &pgp.PGPKeyPair{
		PublicKey:  pubKey,
		PrivateKey: privateKey,
	}, nil
}
