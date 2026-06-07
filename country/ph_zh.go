package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Chinese, "菲律宾")
	dataPhilippines.RegisterOfficialName(xlanguage.Chinese, "菲律宾共和国")
	dataPhilippines.RegisterCapital(xlanguage.Chinese, "马尼拉")
}
