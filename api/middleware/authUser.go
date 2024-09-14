package middleware

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func AuthenticatePW(pw, encHash string) (match bool, err error) {
	p, salt, hash, err := decHash(encHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(pw), salt, p.iter, p.mem, p.parallel, p.keyLen)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1{
		return true, nil
	}
	return false, nil
}

func decHash(encHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("hash is in wrong format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("different argon version")
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d, t=%d, p=%d", &p.mem, &p.iter, &p.parallel)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLen = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLen = uint32(len(hash))

	return p, salt, hash, nil
}