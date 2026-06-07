//go:build country_ae || country_all || country_asia || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Chinese, "阿拉伯联合酋长国")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Chinese, "阿拉伯联合酋长国")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Chinese, "阿布扎比")
}
