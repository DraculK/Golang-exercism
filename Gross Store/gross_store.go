package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit := map[string]int{}
    unit["quarter_of_a_dozen"] = 3
    unit["half_of_a_dozen"] = 6
    unit["dozen"] = 12
    unit["small_gross"] = 120
    unit["gross"] = 144
    unit["great_gross"] = 1728
    return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    aux_count := 0
	for i := range units{
        if unit == i {
            aux_count = 1
            bill[item] += units[i]
        }
    }
	if aux_count == 0{
        return false
    }
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    //_ == item and ok is a bool that returns true if the item is in the map
    _, ok := bill[item]
    if !ok {
        return false
    }

    _, ok = units[unit]
    if !ok {
        return false
    }

    qtd := (bill[item])-(units[unit])
    if qtd == 0{
        delete(bill, item)
        return true
    }else if qtd < 0 {
    	return false
    }else{
    	bill[item] -= units[unit]
    }
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, ok := bill[item]
    if !ok{
        return 0, false
    }else{
    	return bill[item], true
    }
}
