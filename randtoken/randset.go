package randtoken

import (
	"math/rand"
	"strings"
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
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteRune(runes[rand.Intn(len(runes))])
	}
	return sb.String()
}
