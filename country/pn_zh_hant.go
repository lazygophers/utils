//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_pn || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.MustParse("zh-Hant"), "皮特肯群島")
	dataPitcairn.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "皮特肯、亨德森、迪西和奧埃諾群島")
	dataPitcairn.RegisterCapital(xlanguage.MustParse("zh-Hant"), "亞當斯敦")
}
