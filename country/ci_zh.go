//go:build country_africa || country_all || country_ci || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Chinese, "科特迪瓦")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Chinese, "科特迪瓦共和国")
	dataIvoryCoast.RegisterCapital(xlanguage.Chinese, "亚穆苏克罗")
}
