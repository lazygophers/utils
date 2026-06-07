//go:build country_all || country_asia || country_ps || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Chinese, "巴勒斯坦")
	dataPalestine.RegisterOfficialName(xlanguage.Chinese, "巴勒斯坦国")
	dataPalestine.RegisterCapital(xlanguage.Chinese, "东耶路撒冷")
}
