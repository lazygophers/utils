package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.Chinese, "印度")
	dataIndia.RegisterOfficialName(xlanguage.Chinese, "印度共和国")
	dataIndia.RegisterCapital(xlanguage.Chinese, "新德里")
}
