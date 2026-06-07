//go:build (lang_ru || lang_all) && (country_all || country_asia || country_central_asia || country_tm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Russian, "Туркмения")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Russian, "Туркменистан")
	dataTurkmenistan.RegisterCapital(xlanguage.Russian, "Ашхабад")
}
