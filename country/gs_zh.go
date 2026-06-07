//go:build country_all || country_antarctic || country_gs

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Chinese, "南乔治亚和南桑威奇群岛")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Chinese, "南乔治亚和南桑威奇群岛")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Chinese, "爱德华王角")
}
