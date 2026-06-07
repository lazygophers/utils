//go:build country_all || country_asia || country_central_asia || country_tm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Chinese, "土库曼斯坦")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Chinese, "土库曼斯坦")
	dataTurkmenistan.RegisterCapital(xlanguage.Chinese, "阿什哈巴德")
}
