package bank

import "testing"

func TestGetAccountNumber(t *testing.T) {
	// iban := "IR130560088381002399976001"
	specs := []struct {
		name string
		iban string
		exp  string
	}{
		{name: "first", iban: "IR130560088381002399976001", exp: "883-810-2399976-1"},
		{name: "second", iban: "IR130560088381002399972001", exp: "883-810-2399972-1"},
	}

	handler := NewSaman()

	for _, v := range specs {
		t.Run(v.name, func(t *testing.T) {
			if got := handler.GetAccountNumber(v.iban); got != v.exp {
				t.Fatalf("expected result to be %v \n got: \n %v \n", v.exp, got)
			}
		})
	}
	// t.Run("first", func(t *testing.T) {
	// 	if got := handler.GetAccountNumber(iban); got != "883-810-2399976-1" {
	// 		t.Fatalf("expected result to be %v \n got: \n %v \n", "883-810-2399976-1", got)
	// 	}
	// })
	// t.Run("second", func(t *testing.T) {
	// 	if got := handler.GetAccountNumber(iban); got == "883-810-2399976-01" {
	// 		t.Fatalf("expected result to be %v \n got: \n %v \n", "883-810-2399976-1", got)
	// 	}
	// })

}
