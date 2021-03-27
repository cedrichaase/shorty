package generator

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/dongri/go-mnemonic"
)

func GenerateByName(name string) string {
	if name == "mnemonic" {
		return GenerateMnemonic()
	} else if name == "ulid" {
		return GenerateUlid()
	}

	return ""
}

func GenerateMnemonic() string {
	var words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
	return strings.Join(strings.Split(words, " ")[0:2], "-")
}

func GenerateUlid() string {
	t := time.Unix(time.Now().Unix(), 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
