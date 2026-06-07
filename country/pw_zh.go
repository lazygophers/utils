//go:build country_all || country_micronesia || country_oceania || country_pw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Chinese, "帕劳")
	dataPalau.RegisterOfficialName(xlanguage.Chinese, "帕劳共和国")
	dataPalau.RegisterCapital(xlanguage.Chinese, "恩吉鲁穆德")
}
