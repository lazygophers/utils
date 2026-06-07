package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Chinese, "蒙古")
	dataMongolia.RegisterOfficialName(xlanguage.Chinese, "蒙古国")
	dataMongolia.RegisterCapital(xlanguage.Chinese, "乌兰巴托")
}
