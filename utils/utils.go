package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// import (
// 	"crypto/rand"
// 	"math/big"
// 	"time"

// 	"github.com/goombaio/namegenerator"
// )

// func RandomString(n int) string {
// 	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
// 	ret := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
// 		if err != nil {
// 			return ""
// 		}
// 		ret[i] = letters[num.Int64()]
// 	}

// 	return string(ret)
// }

// func RandomName() string {
// 	seed := time.Now().UTC().UnixNano()
// 	nameGenerator := namegenerator.NewNameGerator(seed)
// 	name := nameGenerator.Generate()
// 	return name
// }
