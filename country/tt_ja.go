//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_tt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Japanese, "トリニダード・トバゴ")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Japanese, "トリニダード・トバゴ共和国")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Japanese, "ポートオブスペイン")
}
