package randtoken

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunesWX = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

//LetterRunesHex runes set for trans id generation
var LetterRunesHex = []rune("1234567890ABCDEFabcdef")

//RandStringRunes generate string with fixed length of provided rune set
func RandStringRunes(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
