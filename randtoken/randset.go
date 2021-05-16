package randtoken

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var letterRunesWX = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

//LetterRunesHex runes set for trans id generation
var LetterRunesHex = []rune("1234567890ABCDEFabcdef")

//RandStringRunes generate string with fixed length of provided rune set
func RandStringRunes(n int, runes []rune) (string, error) {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}
		sb.WriteRune(runes[num.Int64()])
	}
	return sb.String(), nil
}
