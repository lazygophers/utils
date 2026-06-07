//go:build country_all || country_americas || country_caribbean || country_do

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Chinese, "多米尼加")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Chinese, "多米尼加共和国")
	dataDominicanRepublic.RegisterCapital(xlanguage.Chinese, "圣多明各")
}
