//go:build country_all || country_by || country_eastern_europe || country_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Russian, "Беларусь")
	dataBelarus.RegisterOfficialName(xlanguage.Russian, "Республика Беларусь")
	dataBelarus.RegisterCapital(xlanguage.Russian, "Минск")
}
