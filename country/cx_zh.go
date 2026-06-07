//go:build country_all || country_australia_and_new_zealand || country_cx || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Chinese, "圣诞岛")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Chinese, "圣诞岛领地")
	dataChristmasIsland.RegisterCapital(xlanguage.Chinese, "飞鱼湾")
}
