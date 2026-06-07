//go:build (lang_zh_hant || lang_all) && (country_all || country_antarctic || country_hm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "赫德島和麥克唐納群島")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "赫德島和麥克唐納群島領地")
}
