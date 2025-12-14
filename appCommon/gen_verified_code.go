package appCommon

import (
	"math/rand"
	"time"
)

func GenVerifiedCode(length int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	s := make([]rune, 6)
	for i := 0; i < 6; i++ {
		s[i] = letters[r1.Intn(99999)%len(letters)]
	}
	return string(s)
}
