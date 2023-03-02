// 解密 lark 发过来的消息

package lark

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

// Decrypt 解密飞书加密事件消息
func Decrypt(payload []byte, key string) ([]byte, error) {
	var d struct {
		Encrypt string `json:"encrypt"`
	}
	err := sonic.Unmarshal(payload, &d)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal error")
	}
	buf, err := base64.StdEncoding.DecodeString(d.Encrypt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(buf) < aes.BlockSize {
		return nil, errors.New("cipher too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return nil, errors.Wrap(err, "aes.NewCipher error")
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := bytes.Index(buf, []byte{'{'})
	if n == -1 {
		n = 0
	}
	m := bytes.LastIndex(buf, []byte{'}'})
	if m == -1 {
		m = len(buf) - 1
	}
	return buf[n : m+1], nil
}
