package config

import (
	"fmt"
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils"
	"github.com/lazygophers/utils/app"
	"github.com/lazygophers/utils/json"
	"github.com/lazygophers/utils/osx"
	"github.com/lazygophers/utils/runtime"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

type Unmarshaler func(reader io.Reader, v interface{}) error
type Marshaler func(writer io.Writer, v interface{}) error

type supportedExt struct {
	Marshaler   Marshaler
	Unmarshaler Unmarshaler
}

var supportedExtMap = map[string]supportedExt{
	".json": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return json.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return json.NewEncoder(writer).Encode(v)
		},
	},
	".toml": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return toml.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return toml.NewEncoder(writer).Encode(v)
		},
	},
	".yaml": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return yaml.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return yaml.NewEncoder(writer).Encode(v)
		},
	},
	".yml": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return yaml.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return yaml.NewEncoder(writer).Encode(v)
		},
	},
}

func RegisterParser(ext string, m Marshaler, u Unmarshaler) {
	supportedExtMap[ext] = supportedExt{
		Marshaler:   m,
		Unmarshaler: u,
	}
}

func tryFindConfigPath(baseDir string) string {
	for ext := range supportedExtMap {
		path := filepath.Join(baseDir, "conf"+ext)
		if osx.IsFile(path) {
			return path
		}

		path = filepath.Join(baseDir, "config"+ext)
		if osx.IsFile(path) {
			return path
		}
	}

	return ""
}

var configPath string

func LoadConfig(c any, paths ...string) error {
	// 依次确认配置文件的位置
	if len(paths) > 0 {
		for _, path := range paths {
			if osx.IsFile(path) {
				configPath = path
				break
			}
		}
	}

	// NOTE: 从环境变量中获取
	log.Warnf("Try to load config from environment variable(LAZYGOPHERS_CONFIG)")
	if configPath != "" {
		configPath = os.Getenv("LAZYGOPHERS_CONFIG")
		if configPath != "" && !osx.IsFile(configPath) {
			log.Debugf("config file not found:%v", configPath)
			configPath = ""
		}
	}

	// NOTE: 从当前目录中获取
	if configPath == "" {
		log.Warnf("Try to load config from %s", runtime.Pwd())
		configPath = tryFindConfigPath(runtime.Pwd())
	}

	// NOTE: 从用户目录中获取
	if configPath == "" {
		log.Warnf("Try to load config from %s", runtime.UserHomeDir())
		configPath = tryFindConfigPath(filepath.Join(runtime.UserHomeDir(), app.Name))
	}

	// NOTE: 从系统目录中获取
	if configPath == "" {
		log.Warnf("Try to load config from %s", runtime.UserConfigDir())
		configPath = tryFindConfigPath(filepath.Join(runtime.UserConfigDir(), app.Name))
	}

	// NOTE: 从程序目录中获取
	if configPath == "" {
		log.Warnf("Try to load config from %s", runtime.ExecDir())
		configPath = tryFindConfigPath(runtime.ExecDir())
	}

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warnf("Config file not found, use default config")
			log.Debugf("config file not found:%v", configPath)
		} else {
			log.Warnf("Config file open failed, use default config")
			log.Errorf("err:%v", err)
		}
	} else {
		defer file.Close()
		log.Infof("Config file found, use config from %s", configPath)

		ext := filepath.Ext(configPath)
		if supported, ok := supportedExtMap[ext]; ok {
			err = supported.Unmarshaler(file, c)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		} else {
			log.Errorf("unsupported config file format:%v", ext)
			log.Errorf("unsupported config file format:%v", ext)
			return fmt.Errorf("unsupported config file format:%v", ext)
		}
	}

	err = utils.Validate(c)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	log.Info("load config success")

	return nil
}
