package str

import "github.com/leekchan/accounting"

// FormatCurrency ...
func FormatCurrency(value float64, currencySymbol, thousandSeparator, decimalSeparator string, precision int) string {
	if currencySymbol == "IDR" {
		// currencySymbol = "Rp "
		currencySymbol = "" // RP already handled on aplus notif template
	}
	ac := accounting.Accounting{Symbol: currencySymbol, Precision: precision, Thousand: thousandSeparator, Decimal: decimalSeparator}

	return ac.FormatMoney(value)
}
