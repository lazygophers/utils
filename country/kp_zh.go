package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Chinese, "朝鲜")
	dataNorthKorea.RegisterOfficialName(xlanguage.Chinese, "朝鲜民主主义人民共和国")
	dataNorthKorea.RegisterCapital(xlanguage.Chinese, "平壤")
}
