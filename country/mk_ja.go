//go:build (lang_ja || lang_all) && (country_all || country_europe || country_mk || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Japanese, "北マケドニア")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Japanese, "北マケドニア共和国")
	dataNorthMacedonia.RegisterCapital(xlanguage.Japanese, "スコピエ")
}
