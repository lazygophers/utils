//go:build country_all || country_asia || country_central_asia || country_kz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Russian, "Казахстан")
	dataKazakhstan.RegisterOfficialName(xlanguage.Russian, "Республика Казахстан")
	dataKazakhstan.RegisterCapital(xlanguage.Russian, "Астана")
}
