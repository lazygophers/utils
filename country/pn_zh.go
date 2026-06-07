//go:build country_all || country_oceania || country_pn || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Chinese, "皮特凯恩群岛")
	dataPitcairn.RegisterOfficialName(xlanguage.Chinese, "皮特凯恩、亨德森、迪西和奥埃诺群岛")
	dataPitcairn.RegisterCapital(xlanguage.Chinese, "亚当斯敦")
}
