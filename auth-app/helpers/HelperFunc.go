package helpers

import (
	"strings"
	"math/rand"
    "fmt"
	"regexp"
	"time"
)

type HelperFunc struct {}

func (helper *HelperFunc) GeneratePass() (string) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 4
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func (helper *HelperFunc) GenerateUsername(name string) (string){
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
    if err != nil {
        fmt.Println(err)
    }
    processedString := reg.ReplaceAllString(name, "")
	return processedString
}