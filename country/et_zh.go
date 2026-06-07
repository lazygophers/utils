//go:build country_africa || country_all || country_eastern_africa || country_et

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Chinese, "埃塞俄比亚")
	dataEthiopia.RegisterOfficialName(xlanguage.Chinese, "埃塞俄比亚联邦民主共和国")
	dataEthiopia.RegisterCapital(xlanguage.Chinese, "亚的斯亚贝巴")
}
