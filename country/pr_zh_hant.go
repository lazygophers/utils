//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.MustParse("zh-Hant"), "波多黎各")
	dataPuertoRico.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "波多黎各自由邦")
	dataPuertoRico.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖胡安")
}
