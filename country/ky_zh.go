//go:build country_all || country_americas || country_caribbean || country_ky

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Chinese, "开曼群岛")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Chinese, "开曼群岛")
	dataCaymanIslands.RegisterCapital(xlanguage.Chinese, "乔治敦")
}
