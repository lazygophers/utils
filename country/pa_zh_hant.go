//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_pa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.MustParse("zh-Hant"), "巴拿馬")
	dataPanama.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴拿馬共和國")
	dataPanama.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴拿馬市")
}
