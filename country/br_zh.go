package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Chinese, "巴西")
	dataBrazil.RegisterOfficialName(xlanguage.Chinese, "巴西联邦共和国")
	dataBrazil.RegisterCapital(xlanguage.Chinese, "巴西利亚")
}
