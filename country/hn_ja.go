//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_hn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Japanese, "ホンジュラス")
	dataHonduras.RegisterOfficialName(xlanguage.Japanese, "ホンジュラス共和国")
	dataHonduras.RegisterCapital(xlanguage.Japanese, "テグシガルパ")
}
