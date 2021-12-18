package gross_store

import "testing"

type TestCase struct {
}

func TestUnits(t *testing.T) {
	tests := []struct {
		name string
		qty  int
	}{
		{"quarter_of_a_dozen", 3},
		{"half_of_a_dozen", 6},
		{"dozen", 12},
		{"small_gross", 120},
		{"gross", 144},
		{"great_gross", 1728},
	}
	units := Units()
	for _, tt := range tests {
		qty, ok := units[tt.name]
		if !ok {
			t.Errorf(`Unit "%s" not found!`, tt.name)
			continue
		}
		if qty != tt.qty {
			t.Errorf(`Unit "%s" should have quantity %d, found %d`, tt.name, tt.qty, qty)
		}
	}
}

type entry struct {
	item string
	unit string
	qty  int
}

func TestAddItem(t *testing.T) {
	tests := []struct {
		name     string
		entry    []entry
		expected bool
	}{
		{
			name: `Invalid measurement`,
			entry: []entry{
				{item: `Orange`, unit: `a little bit`, qty: 0},
				{item: `Orange`, unit: `half`, qty: 0},
				{item: `Orange`, unit: `full`, qty: 0},
			},
			expected: false,
		},
		{
			name: `Valid measurement unit`,
			entry: []entry{
				{item: `Orange`, unit: `quarter_of_a_dozen`, qty: 3},
				{item: `Apple`, unit: `half_of_a_dozen`, qty: 6},
				{item: `Lemon`, unit: `dozen`, qty: 12},
			},
			expected: true,
		},
		{
			name: `check quantity of item that add twice`,
			entry: []entry{
				{"peas", "quarter_of_a_dozen", 3},
				{"peas", "quarter_of_a_dozen", 6},
				{"tomato", "half_of_a_dozen", 6},
				{"tomato", "quarter_of_a_dozen", 9},
			},
			expected: true,
		},
	}
	units := Units()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bill := NewBill()
			for _, row := range tc.entry {
				ok := AddItem(bill, units, row.item, row.unit)
				if ok != tc.expected {
					t.Errorf("Expected %t from AddItem, found %t at %v", tc.expected, ok, row.item)
				}
				qty, ok := bill[row.item]
				// check item
				if ok != tc.expected {
					t.Errorf("Could not find item %s in customer bill", row.item)
				}
				// check qty
				if qty != row.qty {
					t.Errorf("Expected %s to have quantity %d in customer bill, found %d", row.item, row.qty, qty)
				}
			}
		})
	}
}
