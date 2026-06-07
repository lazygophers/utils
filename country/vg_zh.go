//go:build country_all || country_americas || country_caribbean || country_vg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Chinese, "英属维尔京群岛")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Chinese, "英属维尔京群岛")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Chinese, "罗德城")
}
