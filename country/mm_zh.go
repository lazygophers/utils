package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Chinese, "缅甸")
	dataMyanmar.RegisterOfficialName(xlanguage.Chinese, "缅甸联邦共和国")
	dataMyanmar.RegisterCapital(xlanguage.Chinese, "内比都")
}
