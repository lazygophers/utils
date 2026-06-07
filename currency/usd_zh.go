package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Usd.RegisterName(xlanguage.Chinese, "美元")
}
