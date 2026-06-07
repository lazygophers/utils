//go:build country_all || country_asia || country_kw || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Chinese, "科威特")
	dataKuwait.RegisterOfficialName(xlanguage.Chinese, "科威特国")
	dataKuwait.RegisterCapital(xlanguage.Chinese, "科威特城")
}
