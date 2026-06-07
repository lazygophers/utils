package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Chinese, "澳门")
	dataMacao.RegisterOfficialName(xlanguage.Chinese, "中华人民共和国澳门特别行政区")
	dataMacao.RegisterCapital(xlanguage.Chinese, "澳门")
}
