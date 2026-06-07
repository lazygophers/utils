//go:build country_africa || country_all || country_ly || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Chinese, "利比亚")
	dataLibya.RegisterOfficialName(xlanguage.Chinese, "利比亚国")
	dataLibya.RegisterCapital(xlanguage.Chinese, "的黎波里")
}
