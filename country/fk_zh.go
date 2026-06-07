//go:build country_all || country_americas || country_fk || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Chinese, "福克兰群岛")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Chinese, "福克兰群岛")
	dataFalklandIslands.RegisterCapital(xlanguage.Chinese, "斯坦利港")
}
