package cart

import "testing"

func TestCartEmpty(t *testing.T) {
	cart := &Cart{}
	if cart.TotalPrice() != 0 {
		t.Errorf("expected 0")
	}
}

func TestChocolate(t *testing.T) {
	cart := &Cart{}
	cart.AddItems(LineItem{
		Name:      "Chocolate",
		UnitPrice: 10,
	})

	if cart.TotalPrice() != 10 {
		t.Error("Expect 10")
	}
}

func TestItemsWithPromotion(t *testing.T) {
	cart := &Cart{}
	cart.AddItems(LineItem{
		Name:      "Chocolate",
		UnitPrice: 10,
		Quantity:  2,
	})
	cart.AddItems(LineItem{
		Name:      "Pocky",
		UnitPrice: 20,
		Quantity:  2,
	})
	cart.AddItems(LineItem{
		Name:      "Bento",
		UnitPrice: 15,
		Quantity:  2,
	})

	cart.AddPromotion(CashPromotion{
		Name:  "Bento",
		Price: 5,
	})

	if cart.TotalPrice() != 80 {
		t.Error("Expect 80")
	}
}

func TestItemsWithSpecialPromotion(t *testing.T) {
	cart := &Cart{}
	cart.AddItems(LineItem{
		Name:      "Chocolate",
		UnitPrice: 10,
		Quantity:  2,
	})
	cart.AddItems(LineItem{
		Name:      "Pocky",
		UnitPrice: 20,
		Quantity:  2,
	})
	cart.AddItems(LineItem{
		Name:      "Bento",
		UnitPrice: 15,
		Quantity:  2,
	})

	cart.AddPromotion(CashPromotion{
		Name:  "Bento",
		Price: 5,
	})

	cart.AddPromotion(SpecialPromotion{
		Discount: 10,
	})

	if cart.TotalPrice() != 70 {
		t.Error("Expect 70")
	}
}
