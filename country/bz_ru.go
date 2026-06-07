//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bz || country_central_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Russian, "Белиз")
	dataBelize.RegisterOfficialName(xlanguage.Russian, "Белиз")
	dataBelize.RegisterCapital(xlanguage.Russian, "Бельмопан")
}
