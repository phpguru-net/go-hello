package gross_store

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := units[unit]
	if !ok {
		return false
	}
	currentQty, ok := bill[item]
	if !ok {
		bill[item] = qty
		return true
	}
	currentQty += qty
	bill[item] = currentQty
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	productQty, ok := bill[item]
	// not found product
	if !ok {
		return false
	}

	removeQty, ok := units[unit]
	// invalid unit
	if !ok {
		return false
	}
	// reduce qty
	productQty -= removeQty
	if productQty < 0 {
		return false
	}
	// remove product from bill if qty is less than 1
	if productQty == 0 {
		delete(bill, item)
		return true
	}
	// otherwise update qty in bill
	bill[item] = productQty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	if !ok {
		return 0, false
	}
	return qty, true
}
