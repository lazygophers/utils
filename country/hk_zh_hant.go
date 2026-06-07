package country

import xlanguage "golang.org/x/text/language"

// Traditional Chinese is one of Hong Kong's official languages; this file
// is unguarded so it loads in default builds.
func init() {
	dataHongKong.RegisterName(xlanguage.MustParse("zh-Hant"), "香港")
	dataHongKong.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "中華人民共和國香港特別行政區")
	dataHongKong.RegisterCapital(xlanguage.MustParse("zh-Hant"), "香港")
}
