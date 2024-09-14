package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/argon2"
)

// PW Hash/Salt using OAuth & Argon2id
// PW Length = 8-96 chars
// No char restrictions

// params holds all parameters for the argon2id hash
type params struct {
	mem      uint32
	iter     uint32
	parallel uint8
	saltLen  uint32
	keyLen   uint32
}

// HashPW creates struct w/ params and sends back the hashed password to the var that called the func
func HashPW(pw string) (string, error) {
	p := &params{
		mem:      64 * 1024,
		iter:     3,
		parallel: 2,
		saltLen:  16,
		keyLen:   32,
	}

	hash, err := genPW(pw, p)
	if err != nil {
		return "", err
	}

	log.Println(hash)
	return hash, nil
}

// genPW actually hashes pw w/ argon2id
func genPW(pw string, p *params) (encHash string, err error) {
	salt, err := genBytes(p.saltLen)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(pw), salt, p.iter, p.mem, p.parallel, p.keyLen)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	encHash  = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.mem, p.iter, p.parallel, b64salt, b64hash)

	return encHash, nil
}

// genBytes generates random salts for each PW
func genBytes(len uint32) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}