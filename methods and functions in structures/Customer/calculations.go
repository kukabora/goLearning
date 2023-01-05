package customer

import "errors"

const DEFAULT_DISCOUNT = 30

func (c Customer) CalculateDiscount() (Customer, error) {
	if !c.discount {
		return c, errors.New("DISCOUNT IS NOT AVAILABLE")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return c, nil
	}

	c.discountValue = result
	return c, nil
}

func (c *Customer) CalculateDiscountPointer() error {
	if !c.discount {
		return errors.New("DISCOUNT IS NOT AVAILABLE")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return nil
	}

	c.discountValue = result
	return nil
}

func (c Customer) CalcPrice(cust Customer, price int) (int, error) {
	discount, err := c.CalcDicount()
	if err != nil {
		return 0, err
	}
	return discount - price, nil
}

func (c Customer) CalcDicount() (int, error) {
	if !c.discount {
		return 0, errors.New("DISCOUNT IS NOT AVAILABLE")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}
