//go:build country_all || country_americas || country_south_america || country_ve

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Chinese, "委内瑞拉")
	dataVenezuela.RegisterOfficialName(xlanguage.Chinese, "委内瑞拉玻利瓦尔共和国")
	dataVenezuela.RegisterCapital(xlanguage.Chinese, "加拉加斯")
}
