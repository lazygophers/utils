//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.MustParse("zh-Hant"), "盧森堡")
	dataLuxembourg.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "盧森堡大公國")
	dataLuxembourg.RegisterCapital(xlanguage.MustParse("zh-Hant"), "盧森堡市")
}
