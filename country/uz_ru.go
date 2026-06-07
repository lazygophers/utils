//go:build country_all || country_asia || country_central_asia || country_uz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Russian, "Узбекистан")
	dataUzbekistan.RegisterOfficialName(xlanguage.Russian, "Республика Узбекистан")
	dataUzbekistan.RegisterCapital(xlanguage.Russian, "Ташкент")
}
