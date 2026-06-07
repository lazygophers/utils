//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_md)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Russian, "Молдавия")
	dataMoldova.RegisterOfficialName(xlanguage.Russian, "Республика Молдова")
	dataMoldova.RegisterCapital(xlanguage.Russian, "Кишинёв")
}
