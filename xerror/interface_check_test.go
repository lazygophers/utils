package xerror_test

// 跨包接口契约校验：i18n.I18n 必须满足 xerror.Localizer。
// 放外部 _test 包避免 production 代码引入 i18n 依赖。

import (
	"github.com/lazygophers/utils/i18n"
	"github.com/lazygophers/utils/xerror"
)

var _ xerror.Localizer = (*i18n.I18n)(nil)
