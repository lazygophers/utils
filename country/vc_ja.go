//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_vc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Japanese, "セントビンセント・グレナディーン")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Japanese, "セントビンセント・グレナディーン")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Japanese, "キングスタウン")
}
