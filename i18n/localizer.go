package i18n

import (
	"strings"
	"sync"

	"github.com/lazygophers/utils/json"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Localizer 抽象多格式反序列化
type Localizer interface {
	Unmarshal(body []byte, v any) error
}

// LocalizerHandle 用函数适配 Localizer
type LocalizerHandle struct {
	unmarshal func([]byte, any) error
}

// Unmarshal 调用底层反序列化函数
func (h *LocalizerHandle) Unmarshal(body []byte, v any) error {
	return h.unmarshal(body, v)
}

// NewLocalizerHandle 用任意 unmarshal 函数构造 Localizer
func NewLocalizerHandle(unmarshal func([]byte, any) error) *LocalizerHandle {
	return &LocalizerHandle{unmarshal: unmarshal}
}

var (
	jsonHandle = NewLocalizerHandle(json.Unmarshal)
	yamlHandle = NewLocalizerHandle(yaml.Unmarshal)
	tomlHandle = NewLocalizerHandle(toml.Unmarshal)

	localizerMu sync.RWMutex
	localizers  = map[string]Localizer{
		"json": jsonHandle,
		"yaml": yamlHandle,
		"yml":  yamlHandle,
		"toml": tomlHandle,
	}
)

// RegisterLocalizer 注册指定扩展名（"json" 或 ".json" 均接受）的解析器
func RegisterLocalizer(ext string, l Localizer) {
	localizerMu.Lock()
	defer localizerMu.Unlock()
	localizers[normExt(ext)] = l
}

// GetLocalizer 按扩展名查询解析器
func GetLocalizer(ext string) (Localizer, bool) {
	localizerMu.RLock()
	defer localizerMu.RUnlock()
	l, ok := localizers[normExt(ext)]
	return l, ok
}

func normExt(ext string) string {
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}
