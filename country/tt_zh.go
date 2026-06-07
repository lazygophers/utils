//go:build country_all || country_americas || country_caribbean || country_tt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Chinese, "特立尼达和多巴哥")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Chinese, "特立尼达和多巴哥共和国")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Chinese, "西班牙港")
}
