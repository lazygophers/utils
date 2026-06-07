//go:build country_all || country_melanesia || country_oceania || country_pg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Chinese, "巴布亚新几内亚")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Chinese, "巴布亚新几内亚独立国")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Chinese, "莫尔兹比港")
}
