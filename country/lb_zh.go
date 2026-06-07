//go:build country_all || country_asia || country_lb || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Chinese, "黎巴嫩")
	dataLebanon.RegisterOfficialName(xlanguage.Chinese, "黎巴嫩共和国")
	dataLebanon.RegisterCapital(xlanguage.Chinese, "贝鲁特")
}
