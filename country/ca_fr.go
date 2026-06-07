//go:build country_all || country_americas || country_ca || country_northern_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.French, "Canada")
	dataCanada.RegisterOfficialName(xlanguage.French, "Canada")
	dataCanada.RegisterCapital(xlanguage.French, "Ottawa")
}
