//go:build (lang_zh_hant || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.MustParse("zh-Hant"), "安地卡及巴布達")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "安地卡及巴布達")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖約翰")
}
