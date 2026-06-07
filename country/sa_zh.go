//go:build country_all || country_asia || country_sa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Chinese, "沙特阿拉伯")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Chinese, "沙特阿拉伯王国")
	dataSaudiArabia.RegisterCapital(xlanguage.Chinese, "利雅得")
}
