//go:build country_africa || country_all || country_bf || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Chinese, "布基纳法索")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Chinese, "布基纳法索")
	dataBurkinaFaso.RegisterCapital(xlanguage.Chinese, "瓦加杜古")
}
