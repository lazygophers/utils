package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Chinese, "德国")
	dataGermany.RegisterOfficialName(xlanguage.Chinese, "德意志联邦共和国")
	dataGermany.RegisterCapital(xlanguage.Chinese, "柏林")
}
