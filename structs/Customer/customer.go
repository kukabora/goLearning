package customer

type Customer struct {
	Name        string
	Age         int
	balance     int
	Debt        int
	Discount    bool
	CalcDicount func() (int, error)
}

func New(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func CalcPrice(cust Customer, price int) (int, error) {
	discount, err := cust.CalcDicount()
	if err != nil {
		return 0, err
	}
	return discount - price, nil
}
