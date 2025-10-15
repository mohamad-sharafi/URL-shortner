package shortner

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256OF(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58encoded(byte []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(byte)
	if err != nil {
		log.Panicf("failed to encode %v", err)
		os.Exit(1)
	}
	return string(encoded)
}
func GenerateShortUrl(initialLink, userId string) string {
	urlhashByte := sha256OF(initialLink + userId)
	generatedNum := new(big.Int).SetBytes(urlhashByte).Uint64()
	finalStr := base58encoded([]byte(
		fmt.Sprintf("%d", generatedNum),
	))
	return finalStr[:8]
}
