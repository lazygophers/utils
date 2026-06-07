package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Chinese, "瑞士")
	dataSwitzerland.RegisterOfficialName(xlanguage.Chinese, "瑞士联邦")
	dataSwitzerland.RegisterCapital(xlanguage.Chinese, "伯尔尼")
}
