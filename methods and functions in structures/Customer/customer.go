package customer

type Customer struct {
	Name          string
	Age           int
	balance       int
	Debt          int
	discount      bool
	discountValue int
}

func New(name string, age int, balance int, debt int, discount bool) Customer {
	return Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

func NewPointer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}
