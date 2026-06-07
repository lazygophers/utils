//go:build country_all || country_americas || country_northern_america || country_pm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Chinese, "圣皮埃尔和密克隆")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Chinese, "圣皮埃尔和密克隆群岛")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Chinese, "圣皮埃尔")
}
