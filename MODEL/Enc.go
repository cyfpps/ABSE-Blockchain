package MODEL

import (
	"crypto/rand"
	"fmt"
	"hash"
	"math/big"

	"github.com/Nik-U/pbc"
)

type Element = *pbc.Element
type Sigma struct {
	Sigma1 Element
	Sigma2 Element
}
type Index struct {
	I1 Element
	I2 Element
	I3 Element
}
type CT struct {
	C  Element
	C0 Element
	C1 Element
	Cy map[string]*pbc.Element
}

type Attribute string

// 设置多项式并计算q_x(0)的值
func setPolynomialAndComputeQ0(act *ACT, node *Node, parentValue int) {
	if node == nil {
		return
	}

	// 如果是根节点
	if node == act.Root {
		r1, _ := rand.Int(rand.Reader, big.NewInt(100)) // 选择一个随机值r1
		node.Polynomial = make([]int, node.Threshold)
		node.Polynomial[0] = int(r1.Int64())
	} else {
		node.Polynomial = make([]int, node.Threshold)
		node.Polynomial[0] = parentValue
	}

	for _, child := range node.Children {
		setPolynomialAndComputeQ0(act, child, node.EvaluatePolynomial(child.Threshold))
	}
}

// 计算Cy的值
/*func computeCy(node *Node, g *pbc.Element, MK *MasterKey, Cy map[string]*pbc.Element) {
	if node.Attribute != "" { // 只对叶子节点进行操作
		polyValue := new(pbc.Element)
		polyValue.SetInt(node.EvaluatePolynomial(0)) // 将整数转换为pbc.Element
		value := new(pbc.Element).Mul(g, new(pbc.Element).Mul(MK.Hi[node.Attribute], polyValue))
		Cy[node.Attribute] = value
		return
	}

	for _, child := range node.Children {
		computeCy(child, g, MK, Cy)
	}
}
*/
// CTEncrypt 根据公钥PK、访问控制策略P和对称加密密钥ck生成密钥密文CT
func CTEncrypt(PK PublicKey, ck *pbc.Element, MK MasterKey) CT {
	// 使用从PublicKey中提取的值
	g := PK.G
	Egga := PK.Egga
	Beta := MK.Beta
	// 获取配对的群
	pairing := PK.Pairing
	//act := &ACT{Root: ParsePolicy(P)}
	fmt.Println("PK.G1", PK.G)
	// 设置多项式并计算q_x(0)的值
	//	setPolynomialAndComputeQ0(act, act.Root, 0)

	// 计算Cy的值

	//	computeCy(act.Root, g, MK, Cy)

	// 随机生成r1和r2
	r1 := pairing.NewZr().Rand()
	r2 := pairing.NewZr().Rand()

	// 使用pbc库进行计算
	C := pairing.NewGT().Mul(MK.Ck, Egga.PowZn(Egga, r2))
	C0 := pairing.NewG1().PowZn(g, r2)
	// 计算 beta * r2
	beta_r2 := pairing.NewZr().Mul(Beta, r2)
	// 计算 beta * r2 + r1
	tempExpForC1 := pairing.NewZr().Add(beta_r2, r1)
	// 计算 C1 = g^{beta * r2 + r1}
	C1 := pairing.NewG1().PowZn(g, tempExpForC1)
	// 初始化Cy
	Cy := make(map[string]*pbc.Element)
	fmt.Println("PK.G22", PK.G)
	return CT{C, C0, C1, Cy}
}

func hashToBigInt(h hash.Hash, data string) *big.Int {
	hashed := h.Sum([]byte(data))
	return new(big.Int).SetBytes(hashed)
}

func SigEncrypt(PK PublicKey, CT string, sk_DO *pbc.Element) Sigma {
	pairing := PK.G.Pairing()
	h3 := hashToBigInt(PK.H3, CT)
	h3Element := pairing.NewZr().SetBig(h3)
	fmt.Println("PK.G2", PK.G)
	// 计算σ1i = g^H3(CTi)
	sigma1 := pairing.NewG1().PowZn(PK.G, h3Element)
	fmt.Println("sigma.Sigma1", sigma1)
	// 计算σ2i = (w^H3(CTi))^ξ
	temp := pairing.NewG1().PowZn(PK.W, h3Element)
	sigma2 := pairing.NewG1().PowZn(temp, sk_DO)
	l := pairing.NewGT()
	//	c := pairing.NewGT()
	l.Pair(PK.W, PK.G)
	fmt.Println("l:", l)
	te := pairing.NewGT().PowZn(l, sk_DO)
	fmt.Println("sk_DO", sk_DO)
	fmt.Println("te", te)
	tes := pairing.NewGT().PowZn(te, h3Element)
	fmt.Println("tes", tes)
	fmt.Println("sigma.Sigma2", sigma2)
	fmt.Println("PK.G2222", PK.G)
	return Sigma{Sigma1: sigma1, Sigma2: sigma2}
}

func IndexGen(KW []string, PK PublicKey, sk_DO *pbc.Element) Index {
	pairing := PK.Pairing
	fmt.Println("PK.G222", PK.G)
	fmt.Println("sk_DO1", sk_DO)
	// 随机选择ρ
	rho := pairing.NewZr().Rand()
	fmt.Println("rho", rho)
	fmt.Println("sk_DO34", sk_DO)
	// 计算temp1 = ξ - Σ H2(kwi)
	temp1 := pairing.NewZr().Set(sk_DO)

	for _, kw := range KW {
		h2 := hashToBigInt(PK.H2, kw)
		h2Element := pairing.NewZr().SetBig(h2)

		temp1 = temp1.Sub(temp1, h2Element)
	}

	fmt.Println("temp1", temp1)
	fmt.Println("rho4", rho)
	// 计算temp2 = w * g^(-ρ)
	temp3 := pairing.NewZr().Neg(rho)
	temp2 := pairing.NewG1().Mul(PK.W, pairing.NewG1().PowZn(PK.G, temp3))
	fmt.Println("rho3", rho)
	fmt.Println("sk_DO3", sk_DO)
	// 计算I1, I2, I3
	I1 := temp2.PowZn(temp2, temp1.Invert(temp1))
	fmt.Println("rho2", rho)
	I2 := pairing.NewGT().PowZn(pairing.NewGT().Pair(PK.G, PK.G), rho)
	fmt.Println("rho12312", rho)
	fmt.Println("PK.G33456", PK.G)
	I3 := pairing.NewG1().PowZn(PK.W, sk_DO)
	fmt.Println("rho12wqeqwe2", rho)
	fmt.Println("I3", I3)
	fmt.Println("sk_DO4", sk_DO)
	fmt.Println("rho12wqeqwe2", rho)
	fmt.Println("PK.W", PK.W)
	return Index{I1: I1, I2: I2, I3: I3}
}
