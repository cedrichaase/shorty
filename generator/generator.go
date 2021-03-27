package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dongri/go-mnemonic"
	"github.com/oklog/ulid/v2"
)

type AlgoGenerator func() string
type AlgoName string

const (
	Mnemonic AlgoName = "mnemonic"
	Ulid              = "ulid"
)

var generators = map[AlgoName]AlgoGenerator{
	Mnemonic: generateMnemonic,
	Ulid:     generateUlid,
}

func Generate(name AlgoName) (string, error) {
	if generator, ok := generators[name]; ok {
		return generator(), nil
	}

	return "", fmt.Errorf("generator: Unknown algorithm %v", name)
}

func generateMnemonic() string {
	var words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
	return strings.Join(strings.Split(words, " ")[0:2], "-")
}

func generateUlid() string {
	t := time.Unix(time.Now().Unix(), 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
