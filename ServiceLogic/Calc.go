package ServiceLogic

type ICalc interface {
	Add()
	Minus()
	Multiply()
	Division()
}

type Calc struct {
}

func (p *Calc) Add(a, b int32) int32 {

	return a + b
}

func (p *Calc) Minus(a, b int32) int32 {

	return a - b
}

func (p *Calc) Multiply(a, b int32) int32 {

	return a * b
}

func (p *Calc) Division(a, b int32) int32 {
	return (int32)(float32(a) / float32(b))
}
