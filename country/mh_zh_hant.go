//go:build (lang_zh_hant || lang_all) && (country_all || country_mh || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "馬紹爾群島")
	dataMarshallIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬紹爾群島共和國")
	dataMarshallIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬久羅")
}
