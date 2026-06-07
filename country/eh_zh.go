//go:build country_africa || country_all || country_eh || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Chinese, "西撒哈拉")
	dataWesternSahara.RegisterOfficialName(xlanguage.Chinese, "撒拉威阿拉伯民主共和国")
	dataWesternSahara.RegisterCapital(xlanguage.Chinese, "阿尤恩")
}
