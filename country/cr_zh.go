//go:build country_all || country_americas || country_central_america || country_cr

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Chinese, "哥斯达黎加")
	dataCostaRica.RegisterOfficialName(xlanguage.Chinese, "哥斯达黎加共和国")
	dataCostaRica.RegisterCapital(xlanguage.Chinese, "圣何塞")
}
