package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Chinese, "比利时")
	dataBelgium.RegisterOfficialName(xlanguage.Chinese, "比利时王国")
	dataBelgium.RegisterCapital(xlanguage.Chinese, "布鲁塞尔")
}
