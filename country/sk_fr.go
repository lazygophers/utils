//go:build (lang_fr || lang_all) && (country_all || country_eastern_europe || country_europe || country_sk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.French, "Slovaquie")
	dataSlovakia.RegisterOfficialName(xlanguage.French, "République slovaque")
	dataSlovakia.RegisterCapital(xlanguage.French, "Bratislava")
}
