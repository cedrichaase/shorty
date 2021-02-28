package generator

import(
	"time"
	"strings"
	"math/rand"
	"github.com/oklog/ulid/v2"

	"github.com/dongri/go-mnemonic"
)

func GenerateMnemonic() string {
	var words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
	return strings.Join(strings.Split(words, " ")[0:2], "-")
}

func GenerateUlid() string {
	t := time.Unix(time.Now().Unix(), 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
