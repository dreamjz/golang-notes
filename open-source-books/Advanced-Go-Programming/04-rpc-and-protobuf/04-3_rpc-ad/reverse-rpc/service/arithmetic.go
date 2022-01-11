package service

type ArithmeticInterface interface {
	Multiply(args ArithmeticService, reply *int) error
}

type ArithmeticService struct {
	X int
	Y int
}

var _ ArithmeticInterface = ArithmeticService{}

func (ArithmeticService) Multiply(args ArithmeticService, reply *int) error {
	*reply = args.X * args.Y
	return nil
}

func NewArithmeticService() *ArithmeticService {
	return new(ArithmeticService)
}
