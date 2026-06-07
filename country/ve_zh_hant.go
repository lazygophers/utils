//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_south_america || country_ve)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.MustParse("zh-Hant"), "委內瑞拉")
	dataVenezuela.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "委內瑞拉玻利瓦共和國")
	dataVenezuela.RegisterCapital(xlanguage.MustParse("zh-Hant"), "卡拉卡斯")
}
