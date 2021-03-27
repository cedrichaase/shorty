package generator

import (
	"fmt"
	"strings"

	"github.com/dongri/go-mnemonic"
	"github.com/teris-io/shortid"
)

type AlgoGenerator func() string
type AlgoName string

const (
	Mnemonic AlgoName = "mnemonic"
	Ulid              = "sid"
)

var generators = map[AlgoName]AlgoGenerator{
	Mnemonic: generateMnemonic,
	Ulid:     generateSid,
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

func generateSid() string {
	var sid, _ = shortid.Generate()
	return sid
}
