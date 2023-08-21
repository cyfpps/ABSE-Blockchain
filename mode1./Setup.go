package model

import (
	"github.com/Nik-U/pbc"
)

// 初始化全局参数
func xianmen(lambda *pbc.Element) (*GlobalParams, *MasterKey, error) {
	p, err := generatePrime(lambda)
	if err != nil {
		return nil, nil, err
	}

	g, err := generateRandomElement(p)
	if err != nil {
		return nil, nil, err
	}

	w, err := generateRandomElement(p)
	if err != nil {
		return nil, nil, err
	}

	alpha, err := generateRandomElement(p)
	if err != nil {
		return nil, nil, err
	}

	beta, err := generateRandomElement(p)
	if err != nil {
		return nil, nil, err
	}
	pairing := pbc.NewPairing()
	gAlpha := pairing.NewG1().PowZn(g, alpha)

	gBeta := new(pbc.Element).Exp(g, beta)

	eAlpha := pairing.NewG1().Pair(g, g).PowZn(alpha)

	// 定义哈希函数
	h1 := func(data []byte) *pbc.Element {
		hash := generateHashFunction(p, data)
		return hash
	}

	h2 := func(data []byte) *pbc.Element {
		hash := generateHashFunction(p, data)
		return hash
	}

	h3 := func(data []byte) *pbc.Element {
		hash := generateHashFunction(p, data)
		return hash
	}

	// 固定的五个属性集合（五个大学）
	attributes := []string{"cyf", "patient", "male", "female", "ten"}

	// 随机选择一个属性作为满足条件的属性
	attrIndex := rand.Intn(len(attributes))
	attr := attributes[attrIndex]

	// 使用属性集合 attributes
	S := attributes

	// 初始化用户密钥
	UK := make(map[string]*pbc.Element)
	MK := make(map[string]*pbc.Element)
	for _, attr := range S {
		h := generateRandomElement(p) // 生成随机元素作为 h_i
		MK[attr] = h
		UK[attr] = new(pbc.Element).Exp(g, h)
	}

	// 初始化用户密钥
	masterKey := &MasterKey{
		Alpha: alpha,
		Beta:  beta,
		MK:    MK,
	}

	params := &GlobalParams{
		p:       p,
		G:       g,
		W:       w,
		GAlpha:  gAlpha,
		GBeta:   gBeta,
		EAlpha:  eAlpha,
		H1:      h1,
		H2:      h2,
		H3:      h3,
		UK:      UK,
		Pairing: pairing,
	}

	return params, masterKey, nil
}

// 生成指定位数的素数
func generatePrime(zr *pbc.Element) (*pbc.Element, error) {
	bits := 256
	for {
		num, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, err
		}

		if zr.SetBytes(num.Bytes()).ProbablyPrime(20) {
			return zr, nil
		}
	}
}

// 其他函数保持不变
// generateHashFunction 生成伪随机哈希函数的示例实现
func generateHashFunction(p *pbc.Element, data []byte) *pbc.Element {
	// 在实际应用中，您需要使用真正的哈希函数来生成哈希值
	// 这里只是一个示例，使用简单的取模运算模拟哈希函数
	hashValue := p.NewZr().SetFromBytes(data)
	hashValue.Mod(hashValue, p)
	return hashValue
}

func generateRandomElement(p *pbc.Element) (*pbc.Element, error) {
	// 生成一个随机的整数作为元素的字节表示
	randomBytes := make([]byte, p.Bytes())
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// 创建一个新的元素并将随机字节设置为其值
	randomElement := p.NewElement().SetBytes(randomBytes)
	return randomElement, nil
}
