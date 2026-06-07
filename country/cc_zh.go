//go:build country_all || country_australia_and_new_zealand || country_cc || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Chinese, "科科斯（基林）群岛")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Chinese, "科科斯（基林）群岛领地")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Chinese, "西岛")
}
