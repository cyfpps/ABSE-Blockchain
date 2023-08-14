package model

// ProxyDecrypt 进行服务器代理解密
/*func ProxyDecrypt(SK *UserDecryptionKey, CT *Ciphertext, f *big.Int, globalParams *GlobalParams) (*big.Int, error) {
	// 计算 Froot
	Froot := new(pairing).Pair(globalParams.G, globalParams.G).Exp(new(big.Int).Mul(SK.S2, f), globalParams.G)
	Froot.Exp(Froot, new(big.Int).Mul(globalParams.Alpha, CT.R2), globalParams.G)

	// 计算 D
	D := new(pairing).Mul(Froot, new(pairing).Pair(SK.S1, CT.C1))
	D = new(pairing).Div(D, new(pairing).Pair(SK.S2, CT.C))

	return D, nil
}

// UserDecrypt 进行数据使用者二次解密
func UserDecrypt(CT *Ciphertext, D *big.Int, globalParams *GlobalParams) ([]byte, error) {
	// 计算密钥ck
	ck := new(pairing).Mul(CT.C, new(pairing).Pair(globalParams.G, globalParams.G).Exp(globalParams.Alpha, D.Denominator()))

	// 数据使用者解密
	m, err := decrypt(CT.CT, ck, globalParams)
	if err != nil {
		return nil, err
	}

	return m, nil
}*/
