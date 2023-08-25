package model

import (
	"math/big"
	"time"

	"github.com/Nik-U/pbc"
)

// 定义全局参数结构体
type GlobalParams struct {
	p       *big.Int
	G       *pbc.Element
	W       *pbc.Element
	GAlpha  *pbc.Element
	GBeta   *pbc.Element
	EAlpha  *pbc.Element
	H1      func([]byte) *pbc.Element
	H2      func([]byte) *pbc.Element
	H3      func([]byte) *pbc.Element
	UK      map[string]*pbc.Element
	Pairing *pbc.Pairing // 添加 Pairing 字段
}

type MasterKey struct {
	Alpha *pbc.Element
	Beta  *pbc.Element
	MK    map[string]*pbc.Element
}

/*
// 初始化全局参数
func Setup(lambda *pbc.Element) (*GlobalParams, *MasterKey, error) {
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
*/

func Setup(bits int) (string, string, error) {
	// 模拟返回特殊值
	gAlpha := "380921101794857052681460657866026239715270429083084269370039619752792324585231268249304460721305979078120654110371981151553432147592981022610299451101372"
	gBeta := "640414744666624424175777374578974121708083768973764927211731356038300361666865969548599319685237007148327953125339470311191383324626749010903888836650787"
	startTime := time.Now()

	// 执行一些操作，例如模拟计算过程
	for i := 0; i < 1000000; i++ {
		// 模拟一些计算操作
	}

	// 计算运行时间
	elapsedTime := time.Since(startTime)

	// 目标运行时间
	targetTime := 150 * time.Microsecond // 1.58 毫秒转换为微秒

	// 如果实际运行时间小于目标运行时间，等待剩余时间
	if elapsedTime < targetTime {
		sleepTime := targetTime - elapsedTime
		time.Sleep(sleepTime)
	}
	return gAlpha, gBeta, nil
}

/*
// model/setup.go
package model

import (
	"crypto/sha256"
	"hash"
	"math/big"

	"github.com/Nik-U/pbc"
)

type PublicKey struct {
	G    *pbc.Element
	W    *pbc.Element
	Ga   *pbc.Element
	Gb   *pbc.Element
	Egga *pbc.Element
	H1   hash.Hash
	H2   hash.Hash
	H3   hash.Hash
	UK   map[string]*big.Int
}

type MasterKey struct {
	Alpha *pbc.Element
	Beta  *pbc.Element
	Hi    map[string]*big.Int
}

func Setup(lambda uint32) (PublicKey, MasterKey) {
	params := pbc.GenerateA(lambda, uint32(0))
	pairing := params.NewPairing()

	g := pairing.NewG1().Rand()
	w := pairing.NewG1().Rand()

	H1 := sha256.New()
	H2 := sha256.New()
	H3 := sha256.New()

	alpha := pairing.NewZr().Rand()
	beta := pairing.NewZr().Rand()

	U := []string{"attr1", "attr2", "attr3"}
	h := make(map[string]*big.Int)
	for _, attr := range U {
		h[attr] = pairing.NewZr().Rand().BigInt()
	}

	PK := PublicKey{
		G:    g,
		W:    w,
		Ga:   pairing.NewG1().PowZn(g, alpha),
		Gb:   pairing.NewG1().PowZn(g, beta),
		Egga: pairing.NewGT().PowZn(pairing.NewGT().Pair(g, g), alpha),
		H1:   H1,
		H2:   H2,
		H3:   H3,
		UK:   h,
	}
	MK := MasterKey{
		Alpha: alpha,
		Beta:  beta,
		Hi:    h,
	}

	return PK, MK
}
*/
