package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cny.RegisterName(xlanguage.Chinese, "人民币")
}
