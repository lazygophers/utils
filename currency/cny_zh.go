package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CNY.RegisterName(xlanguage.Chinese, "人民币")
}
