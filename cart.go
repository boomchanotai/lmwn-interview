package cart

type LineItem struct {
	Name      string
	UnitPrice int
	Quantity  int
}

type CashPromotion struct {
	Name  string
	Price int
}

type SpecialPromotion struct {
}

type Cart struct {
	Items            []LineItem
	Promotion        map[string]CashPromotion
	SpecialPromotion []SpecialPromotion
}

func (c *Cart) TotalPrice() int {
	totalPrice := 0
	for i := 0; i < len(c.Items); i++ {
		item := c.Items[i]
		pro, ok := c.Promotion[item.Name]
		if ok {
			totalPrice += (item.UnitPrice - pro.Price) * item.Quantity
		} else {
			totalPrice += item.UnitPrice * item.Quantity
		}
	}

	return totalPrice
}

func (c *Cart) AddItems(item LineItem) {
	if item.Quantity == 0 {
		item.Quantity = 1
	}
	c.Items = append(c.Items, item)
}

func (c *Cart) AddPromotion(cashPromotion interface{}) {
	// parse

	if ok {
		// CashPromotion
		if c.Promotion == nil {
			c.Promotion = map[string]CashPromotion{}
		}
		c.Promotion[cashPromotion.Name] = cashPromotion
	} else {
		// Special
	}
}

// 20 + 40 + 30 = 90
// 20 + 40 + (30-10) = 80
