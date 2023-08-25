// model/setup.go
package MODEL

import (
	"crypto/sha256"
	"hash"

	"github.com/Nik-U/pbc"
)

type PublicKey struct {
	G       *pbc.Element
	W       *pbc.Element
	Ga      *pbc.Element
	Gb      *pbc.Element
	Egga    *pbc.Element
	H1      hash.Hash
	H2      hash.Hash
	H3      hash.Hash
	UK      map[string]*pbc.Element
	Pairing *pbc.Pairing // 添加Pairing到PublicKey结构体中
}

type MasterKey struct {
	Ck    *pbc.Element
	Alpha *pbc.Element
	Beta  *pbc.Element
	Hi    map[string]*pbc.Element
}

func Setup(lambda uint32) (PublicKey, MasterKey) {
	params := pbc.GenerateA(lambda, uint32(0))
	pairing := params.NewPairing()

	g := pairing.NewG1().Rand()
	w := pairing.NewG1().Rand()
	ck := pairing.NewGT().Rand()

	H1 := sha256.New()
	H2 := sha256.New()
	H3 := sha256.New()

	alpha := pairing.NewZr().Rand()
	beta := pairing.NewZr().Rand()

	U := []string{"attr1", "attr2", "attr3", "attr4"}
	hi := make(map[string]*pbc.Element)
	uk := make(map[string]*pbc.Element)
	for _, attr := range U {
		randomElement := pairing.NewZr().Rand()
		hi[attr] = randomElement
		uk[attr] = pairing.NewG1().PowZn(g, randomElement)
	}

	PK := PublicKey{
		G:       g,
		W:       w,
		Ga:      pairing.NewG1().PowZn(g, alpha),
		Gb:      pairing.NewG1().PowZn(g, beta),
		Egga:    pairing.NewGT().PowZn(pairing.NewGT().Pair(g, g), alpha),
		H1:      H1,
		H2:      H2,
		H3:      H3,
		UK:      uk,
		Pairing: pairing, // 初始化Pairing
	}
	MK := MasterKey{
		Ck:    ck,
		Alpha: alpha,
		Beta:  beta,
		Hi:    hi,
	}

	return PK, MK
}
