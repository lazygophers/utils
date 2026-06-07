//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.MustParse("zh-Hant"), "馬約特")
	dataMayotte.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬約特省")
	dataMayotte.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬穆楚")
}
