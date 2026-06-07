package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.MustParse("zh-Hant"), "臺灣")
	dataTaiwan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "中華民國（臺灣）")
	dataTaiwan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "臺北")
}
