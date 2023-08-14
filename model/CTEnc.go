package model

import (
	"github.com/Nik-U/pbc"
)

// Ciphertext 密文结构体
type Ciphertext struct {
	C  *pbc.Element
	C0 *pbc.Element
	C1 *pbc.Element
	Cy map[string]*pbc.Element
}

/*// CTEncrypt 根据访问策略P和ck生成密文CT
func CTEncrypt(globalParams *GlobalParams, ck *pbc.Element) (*Ciphertext, error) {
		r2, err := generateRandomElement(globalParams.G)
			if err != nil {
				return nil, err
			}
			r1, err := generateRandomElement(globalParams.G)
			if err != nil {
				return nil, err
			}

				C := globalParams.Pairing.NewG2().Set2(globalParams.G, globalParams.G).PowZn(globalParams.EAlpha).PowZn(r2)

				C0 := globalParams.Pairing.NewG2().Set(globalParams.G).PowZn(r2)

				C1 := globalParams.Pairing.NewG1().PowZn(new(big.Int).Mul(globalParams.GBeta, r2)).AddZn(globalParams.Pairing.NewG1().PowZn(r1))

				Cy := make(map[string]*pbc.Element)
				for _, y := range P {
					qx := P[0].Qx.Eval(new(big.Int).SetInt64(int64(y.Index)))
					hashedQx := new(big.Int).SetBytes(generateHashFunction(globalParams.G.Bytes(), qx))
					for _, attr := range y.Attributes {
						h := globalParams.UK[attr]
						Cy[attr] = globalParams.G.NewElement().PowZn(hashedQx).PowZn(h)
					}
				}

		return &Ciphertext{
			C:  C,
			C0: C0,
			C1: C1,
			Cy: Cy,
		}, nil
}*/
