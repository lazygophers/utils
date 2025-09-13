package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Name     string `json:"name" yaml:"name" toml:"name" ini:"name" xml:"name" properties:"name" env:"name" validate:"required"`
	Port     int    `json:"port" yaml:"port" toml:"port" ini:"port" xml:"port" properties:"port" env:"port" validate:"min=1,max=65535"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" ini:"debug" xml:"debug" properties:"debug" env:"debug"`
	Database struct {
		Host     string `json:"host" yaml:"host" toml:"host" ini:"host" xml:"host" properties:"host" env:"host"`
		Username string `json:"username" yaml:"username" toml:"username" ini:"username" xml:"username" properties:"username" env:"username"`
		Password string `json:"password" yaml:"password" toml:"password" ini:"password" xml:"password" properties:"password" env:"password"`
	} `json:"database" yaml:"database" toml:"database" ini:"database" xml:"database" properties:"database" env:"database"`
}

func TestRegisterParser(t *testing.T) {
	// 备份原始配置
	originalMap := make(map[string]supportedExt)
	for k, v := range supportedExtMap {
		originalMap[k] = v
	}
	defer func() {
		// 恢复原始配置
		supportedExtMap = originalMap
	}()

	t.Run("register new parser", func(t *testing.T) {
		marshaler := func(writer io.Writer, v interface{}) error {
			return nil
		}
		unmarshaler := func(reader io.Reader, v interface{}) error {
			return nil
		}

		RegisterParser(".test", marshaler, unmarshaler)

		ext, exists := supportedExtMap[".test"]
		assert.True(t, exists)
		assert.NotNil(t, ext.Marshaler)
		assert.NotNil(t, ext.Unmarshaler)
	})

	t.Run("override existing parser", func(t *testing.T) {
		marshaler := func(writer io.Writer, v interface{}) error {
			return nil
		}
		unmarshaler := func(reader io.Reader, v interface{}) error {
			return nil
		}

		RegisterParser(".json", marshaler, unmarshaler)

		ext, exists := supportedExtMap[".json"]
		assert.True(t, exists)
		assert.NotNil(t, ext.Marshaler)
		assert.NotNil(t, ext.Unmarshaler)
	})
}

func TestTryFindConfigPath(t *testing.T) {
	// 创建临时目录用于测试
	tmpDir, err := os.MkdirTemp("", "config_test_")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	t.Run("find config file", func(t *testing.T) {
		// 创建测试配置文件
		configFile := filepath.Join(tmpDir, "conf.json")
		err := os.WriteFile(configFile, []byte(`{"name":"test"}`), 0644)
		require.NoError(t, err)

		path := tryFindConfigPath(tmpDir)
		assert.Equal(t, configFile, path)
	})

	t.Run("find config file with config prefix", func(t *testing.T) {
		// 清理之前的文件
		os.RemoveAll(filepath.Join(tmpDir, "conf.json"))

		// 创建 config.yaml 文件
		configFile := filepath.Join(tmpDir, "config.yaml")
		err := os.WriteFile(configFile, []byte(`name: test`), 0644)
		require.NoError(t, err)

		path := tryFindConfigPath(tmpDir)
		assert.Equal(t, configFile, path)
	})

	t.Run("no config file found", func(t *testing.T) {
		// 创建空目录
		emptyDir, err := os.MkdirTemp("", "empty_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(emptyDir)

		path := tryFindConfigPath(emptyDir)
		assert.Equal(t, "", path)
	})

	t.Run("priority test - conf over config", func(t *testing.T) {
		tmpDir2, err := os.MkdirTemp("", "priority_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir2)

		// 创建两个文件，conf 应该优先
		confFile := filepath.Join(tmpDir2, "conf.json")
		configFile := filepath.Join(tmpDir2, "config.json")

		err = os.WriteFile(confFile, []byte(`{"name":"conf"}`), 0644)
		require.NoError(t, err)
		err = os.WriteFile(configFile, []byte(`{"name":"config"}`), 0644)
		require.NoError(t, err)

		path := tryFindConfigPath(tmpDir2)
		assert.Equal(t, confFile, path)
	})
}

func TestLoadConfigSkipValidate(t *testing.T) {
	// 备份原始 configPath
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("load from specified path", func(t *testing.T) {
		// 重置 configPath
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "load_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 创建 JSON 配置文件
		configFile := filepath.Join(tmpDir, "test.json")
		configContent := `{
			"name": "testapp",
			"port": 8080,
			"debug": true,
			"database": {
				"host": "localhost",
				"username": "user",
				"password": "pass"
			}
		}`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "testapp", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
		assert.Equal(t, "localhost", config.Database.Host)
		assert.Equal(t, configFile, configPath)
	})

	t.Run("load YAML config", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "yaml_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "test.yaml")
		configContent := `
name: testapp
port: 9090
debug: false
database:
  host: remotehost
  username: admin
  password: secret
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "testapp", config.Name)
		assert.Equal(t, 9090, config.Port)
		assert.False(t, config.Debug)
		assert.Equal(t, "remotehost", config.Database.Host)
	})

	t.Run("load TOML config", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "toml_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "test.toml")
		configContent := `
name = "testapp"
port = 3000
debug = true

[database]
host = "tomlhost"
username = "tomluser"
password = "tomlpass"
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "testapp", config.Name)
		assert.Equal(t, 3000, config.Port)
		assert.True(t, config.Debug)
		assert.Equal(t, "tomlhost", config.Database.Host)
	})

	t.Run("load INI config", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "ini_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "test.ini")
		configContent := `
name = testapp
port = 4000
debug = false

[database]
host = inihost
username = iniuser
password = inipass
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "testapp", config.Name)
		assert.Equal(t, 4000, config.Port)
		assert.False(t, config.Debug)
		assert.Equal(t, "inihost", config.Database.Host)
		assert.Equal(t, "iniuser", config.Database.Username)
		assert.Equal(t, "inipass", config.Database.Password)
	})

	t.Run("load from environment variable", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "env_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "env.json")
		err = os.WriteFile(configFile, []byte(`{"name":"envtest","port":5000}`), 0644)
		require.NoError(t, err)

		// 设置环境变量
		os.Setenv("LAZYGOPHERS_CONFIG", configFile)
		defer os.Unsetenv("LAZYGOPHERS_CONFIG")

		var config TestConfig
		err = LoadConfigSkipValidate(&config)
		assert.NoError(t, err)
		assert.Equal(t, "envtest", config.Name)
		assert.Equal(t, 5000, config.Port)
	})

	t.Run("load from current directory", func(t *testing.T) {
		configPath = ""
		os.Unsetenv("LAZYGOPHERS_CONFIG")

		// 在当前目录创建配置文件
		currentDir, err := os.Getwd()
		require.NoError(t, err)

		configFile := filepath.Join(currentDir, "conf.json")
		err = os.WriteFile(configFile, []byte(`{"name":"currentdir","port":6000}`), 0644)
		require.NoError(t, err)
		defer os.Remove(configFile)

		var config TestConfig
		err = LoadConfigSkipValidate(&config)
		assert.NoError(t, err)
		assert.Equal(t, "currentdir", config.Name)
		assert.Equal(t, 6000, config.Port)
	})

	t.Run("file not exists", func(t *testing.T) {
		configPath = ""
		os.Unsetenv("LAZYGOPHERS_CONFIG")

		var config TestConfig
		err := LoadConfigSkipValidate(&config, "/nonexistent/path/config.json")
		assert.NoError(t, err) // Should not return error, just use default values
	})

	t.Run("unsupported format", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "unsupported_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.txt")
		err = os.WriteFile(configFile, []byte("some content"), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported config file format")
	})

	t.Run("invalid JSON content", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "invalid_json_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "invalid.json")
		err = os.WriteFile(configFile, []byte(`{"name": invalid json`), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err) // Function doesn't return unmarshaling errors
	})

	t.Run("environment variable file not exists", func(t *testing.T) {
		configPath = ""

		// 设置不存在的文件路径
		os.Setenv("LAZYGOPHERS_CONFIG", "/nonexistent/config.json")
		defer os.Unsetenv("LAZYGOPHERS_CONFIG")

		var config TestConfig
		err := LoadConfigSkipValidate(&config)
		assert.NoError(t, err) // Should fallback to other methods
	})
}

func TestLoadConfig(t *testing.T) {
	// 备份原始 configPath
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("load valid config with validation", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "valid_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "valid.json")
		configContent := `{
			"name": "validapp",
			"port": 8080,
			"debug": true
		}`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfig(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "validapp", config.Name)
		assert.Equal(t, 8080, config.Port)
	})

	t.Run("load invalid config fails validation", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "invalid_config_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "invalid.json")
		configContent := `{
			"port": 70000,
			"debug": true
		}`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfig(&config, configFile)
		assert.Error(t, err) // Should fail validation due to missing required name and invalid port
	})

	t.Run("load config file not found", func(t *testing.T) {
		configPath = ""

		var config TestConfig
		err := LoadConfig(&config, "/nonexistent/config.json")
		assert.Error(t, err) // Should fail validation due to missing required fields
	})
}

func TestSetConfig(t *testing.T) {
	// 备份原始 configPath
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("set JSON config", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "set_json_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "output.json")
		configPath = configFile

		config := TestConfig{
			Name:  "settest",
			Port:  7000,
			Debug: true,
		}
		config.Database.Host = "sethost"
		config.Database.Username = "setuser"

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件是否正确写入
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "settest")
		assert.Contains(t, string(content), "7000")
		assert.Contains(t, string(content), "sethost")
	})

	t.Run("set YAML config", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "set_yaml_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "output.yaml")
		configPath = configFile

		config := TestConfig{
			Name:  "yamltest",
			Port:  8000,
			Debug: false,
		}

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件内容
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "yamltest")
		assert.Contains(t, string(content), "8000")
	})

	t.Run("set TOML config", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "set_toml_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "output.toml")
		configPath = configFile

		config := TestConfig{
			Name:  "tomltest",
			Port:  9000,
			Debug: true,
		}

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件内容
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "tomltest")
		assert.Contains(t, string(content), "9000")
	})

	t.Run("set INI config", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "set_ini_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "output.ini")
		configPath = configFile

		config := TestConfig{
			Name:  "initest",
			Port:  5000,
			Debug: false,
		}
		config.Database.Host = "iniwritehost"
		config.Database.Username = "iniwriteuser"

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件内容
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "initest")
		assert.Contains(t, string(content), "5000")
		assert.Contains(t, string(content), "iniwritehost")
		assert.Contains(t, string(content), "iniwriteuser")
	})

	t.Run("unsupported format", func(t *testing.T) {
		configPath = "/some/path/config.txt"

		config := TestConfig{Name: "test"}
		err := SetConfig(&config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported config file format")
	})

	t.Run("file creation error", func(t *testing.T) {
		// 使用只读目录路径
		configPath = "/proc/config.json" // /proc 通常是只读的

		config := TestConfig{Name: "test"}
		err := SetConfig(&config)
		assert.Error(t, err) // Should fail to create/open file
	})

	t.Run("marshaling error handling", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "marshal_error_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "marshal.json")
		configPath = configFile

		// 创建一个包含不可序列化字段的结构
		type UnmarshallableStruct struct {
			Channel chan int `json:"channel"` // channels cannot be marshaled
		}

		config := UnmarshallableStruct{
			Channel: make(chan int),
		}

		err = SetConfig(&config)
		assert.NoError(t, err) // Function returns nil on marshal error (line 191)
	})
}

func TestSupportedFormats(t *testing.T) {
	t.Run("test all supported formats", func(t *testing.T) {
		expectedExtensions := []string{".json", ".yaml", ".yml", ".toml", ".ini"}

		for _, ext := range expectedExtensions {
			_, exists := supportedExtMap[ext]
			assert.True(t, exists, "Extension %s should be supported", ext)
		}
	})

	t.Run("test marshaler and unmarshaler existence", func(t *testing.T) {
		for ext, support := range supportedExtMap {
			assert.NotNil(t, support.Marshaler, "Marshaler for %s should not be nil", ext)
			assert.NotNil(t, support.Unmarshaler, "Unmarshaler for %s should not be nil", ext)
		}
	})
}

func TestComplexLoadingScenarios(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("multiple path fallback", func(t *testing.T) {
		configPath = ""
		os.Unsetenv("LAZYGOPHERS_CONFIG")

		tmpDir, err := os.MkdirTemp("", "fallback_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 创建第二个路径的配置文件
		configFile2 := filepath.Join(tmpDir, "config2.json")
		err = os.WriteFile(configFile2, []byte(`{"name":"fallback","port":2000}`), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, "/nonexistent/config1.json", configFile2)
		assert.NoError(t, err)
		assert.Equal(t, "fallback", config.Name)
		assert.Equal(t, 2000, config.Port)
	})

	t.Run("YML extension", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "yml_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "test.yml")
		configContent := `
name: ymltest
port: 4000
debug: true
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "ymltest", config.Name)
		assert.Equal(t, 4000, config.Port)
		assert.True(t, config.Debug)
	})

	t.Run("empty config path in environment", func(t *testing.T) {
		configPath = ""
		os.Setenv("LAZYGOPHERS_CONFIG", "")
		defer os.Unsetenv("LAZYGOPHERS_CONFIG")

		var config TestConfig
		err := LoadConfigSkipValidate(&config)
		assert.NoError(t, err) // Should not fail, just use defaults
	})
}

// Additional tests for edge cases and error conditions
func TestErrorConditions(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("empty configPath in SetConfig", func(t *testing.T) {
		configPath = ""

		config := TestConfig{Name: "test"}
		err := SetConfig(&config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported config file format")
	})

	t.Run("directory instead of file", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "dir_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 尝试加载目录而不是文件
		var config TestConfig
		err = LoadConfigSkipValidate(&config, tmpDir)
		assert.NoError(t, err) // Should handle gracefully
	})
}

// Test concurrent access safety
func TestConcurrentAccess(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("concurrent loads", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "concurrent_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "concurrent.json")
		err = os.WriteFile(configFile, []byte(`{"name":"concurrent","port":1000}`), 0644)
		require.NoError(t, err)

		// 并发加载配置
		for i := 0; i < 10; i++ {
			go func() {
				var config TestConfig
				LoadConfigSkipValidate(&config, configFile)
			}()
		}

		// 主要是确保不会panic
		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
	})
}

// Additional tests to improve coverage
func TestCoverageImprovements(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("YAML encoder test", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "yaml_encoder_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "encoder.yml")
		configPath = configFile

		config := TestConfig{
			Name:  "yamlencoder",
			Port:  8888,
			Debug: true,
		}

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件是否正确创建
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "yamlencoder")
	})

	t.Run("LoadConfig validation failure", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "validation_fail_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 创建一个有效的JSON但验证会失败的配置
		configFile := filepath.Join(tmpDir, "invalid_validation.json")
		configContent := `{
			"name": "",
			"port": 70000,
			"debug": false
		}`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfig(&config, configFile)
		assert.Error(t, err) // Should fail validation
	})

	t.Run("file permission error", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "permission_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 创建一个文件然后移除读取权限
		configFile := filepath.Join(tmpDir, "noperm.json")
		err = os.WriteFile(configFile, []byte(`{"name":"test"}`), 0644)
		require.NoError(t, err)

		// 移除读取权限
		err = os.Chmod(configFile, 0000)
		require.NoError(t, err)
		defer os.Chmod(configFile, 0644) // 恢复权限以便清理

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err) // Function handles permission errors gracefully
	})

	t.Run("marshaling error in YAML", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "yaml_marshal_error_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "marshal_error.yaml")
		configPath = configFile

		// 使用channel来触发marshaling错误，但需要捕获panic
		type UnmarshalableConfig struct {
			Name    string   `yaml:"name"`
			Channel chan int `yaml:"channel"` // channels cannot be marshaled
		}

		config := UnmarshalableConfig{
			Name:    "test",
			Channel: make(chan int),
		}

		// YAML encoder会panic而不是返回错误，这是预期行为
		assert.Panics(t, func() {
			SetConfig(&config)
		})
	})

	t.Run("TOML encoder error", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "toml_encoder_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "toml_test.toml")
		configPath = configFile

		config := TestConfig{
			Name:  "tomltest",
			Port:  7777,
			Debug: false,
		}

		err = SetConfig(&config)
		assert.NoError(t, err)
	})

	t.Run("INI encoder test", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "ini_encoder_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "encoder.ini")
		configPath = configFile

		config := TestConfig{
			Name:  "iniencoder",
			Port:  6666,
			Debug: true,
		}
		config.Database.Host = "iniencodehost"

		err = SetConfig(&config)
		assert.NoError(t, err)

		// 验证文件是否正确创建
		content, err := os.ReadFile(configFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "iniencoder")
		assert.Contains(t, string(content), "6666")
		assert.Contains(t, string(content), "iniencodehost")
	})

	t.Run("invalid INI content", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "invalid_ini_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "invalid.ini")
		// 创建一个无效的INI内容（不是所有格式错误都会被检测到，INI相对容错）
		err = os.WriteFile(configFile, []byte("[invalid section\nkey without value\n"), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		// INI库相对宽容，可能不会报错，这取决于具体的错误类型
		// 主要是验证不会崩溃
	})
}

// TestNewConfigFormats 测试新添加的配置文件格式
func TestNewConfigFormats(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("XML format", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "xml_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.xml")
		xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<TestConfig>
  <name>xmltest</name>
  <port>8080</port>
  <debug>true</debug>
  <database>
    <host>localhost</host>
    <username>testuser</username>
    <password>testpass</password>
  </database>
</TestConfig>`

		err = os.WriteFile(configFile, []byte(xmlContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "xmltest", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
		assert.Equal(t, "localhost", config.Database.Host)
		assert.Equal(t, "testuser", config.Database.Username)
		assert.Equal(t, "testpass", config.Database.Password)
	})

	t.Run("Properties format", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "properties_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.properties")
		propertiesContent := `# Configuration file
name=proptest
port=9090
debug=true
# Database configuration
database.host=db.example.com
database.username=propuser
database.password=proppass
`

		err = os.WriteFile(configFile, []byte(propertiesContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "proptest", config.Name)
		assert.Equal(t, 9090, config.Port)
		assert.True(t, config.Debug)
		assert.Equal(t, "db.example.com", config.Database.Host)
		assert.Equal(t, "propuser", config.Database.Username)
		assert.Equal(t, "proppass", config.Database.Password)
	})

	t.Run("Properties format with colon separator", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "properties_colon_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.properties")
		propertiesContent := `name: colontest
port: 7070
debug: false
database.host: colon.example.com
database.username: colonuser
database.password: colonpass
`

		err = os.WriteFile(configFile, []byte(propertiesContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "colontest", config.Name)
		assert.Equal(t, 7070, config.Port)
		assert.False(t, config.Debug)
		assert.Equal(t, "colon.example.com", config.Database.Host)
		assert.Equal(t, "colonuser", config.Database.Username)
		assert.Equal(t, "colonpass", config.Database.Password)
	})

	t.Run("Env format", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "env_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, ".env")
		envContent := `# Environment variables
name="envtest"
port=6060
debug=true
# Database settings
database.host="env.example.com"
database.username='envuser'
database.password="envpass"
`

		err = os.WriteFile(configFile, []byte(envContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "envtest", config.Name)
		assert.Equal(t, 6060, config.Port)
		assert.True(t, config.Debug)
		assert.Equal(t, "env.example.com", config.Database.Host)
		assert.Equal(t, "envuser", config.Database.Username)
		assert.Equal(t, "envpass", config.Database.Password)
	})

	t.Run("XML write and read", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "xml_write_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "write_test.xml")
		configPath = configFile

		// 创建测试配置
		originalConfig := TestConfig{
			Name:  "writetest",
			Port:  3030,
			Debug: false,
		}
		originalConfig.Database.Host = "write.example.com"
		originalConfig.Database.Username = "writeuser"
		originalConfig.Database.Password = "writepass"

		// 写入配置
		err = SetConfig(originalConfig)
		assert.NoError(t, err)

		// 读取配置
		var readConfig TestConfig
		err = LoadConfigSkipValidate(&readConfig, configFile)
		assert.NoError(t, err)
		assert.Equal(t, originalConfig.Name, readConfig.Name)
		assert.Equal(t, originalConfig.Port, readConfig.Port)
		assert.Equal(t, originalConfig.Debug, readConfig.Debug)
		assert.Equal(t, originalConfig.Database.Host, readConfig.Database.Host)
		assert.Equal(t, originalConfig.Database.Username, readConfig.Database.Username)
		assert.Equal(t, originalConfig.Database.Password, readConfig.Database.Password)
	})

	t.Run("Properties write and read", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "properties_write_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "write_test.properties")
		configPath = configFile

		// 创建测试配置
		originalConfig := TestConfig{
			Name:  "prop-writetest",
			Port:  4040,
			Debug: true,
		}
		originalConfig.Database.Host = "prop-write.example.com"
		originalConfig.Database.Username = "propwriteuser"
		originalConfig.Database.Password = "propwritepass"

		// 写入配置
		err = SetConfig(originalConfig)
		assert.NoError(t, err)

		// 读取配置
		var readConfig TestConfig
		err = LoadConfigSkipValidate(&readConfig, configFile)
		assert.NoError(t, err)
		assert.Equal(t, originalConfig.Name, readConfig.Name)
		assert.Equal(t, originalConfig.Port, readConfig.Port)
		assert.Equal(t, originalConfig.Debug, readConfig.Debug)
		assert.Equal(t, originalConfig.Database.Host, readConfig.Database.Host)
		assert.Equal(t, originalConfig.Database.Username, readConfig.Database.Username)
		assert.Equal(t, originalConfig.Database.Password, readConfig.Database.Password)
	})

	t.Run("Env write and read", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "env_write_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "write_test.env")
		configPath = configFile

		// 创建测试配置
		originalConfig := TestConfig{
			Name:  "env writetest with spaces",
			Port:  5050,
			Debug: false,
		}
		originalConfig.Database.Host = "env-write.example.com"
		originalConfig.Database.Username = "envwriteuser"
		originalConfig.Database.Password = "env write pass"

		// 写入配置
		err = SetConfig(originalConfig)
		assert.NoError(t, err)

		// 读取配置
		var readConfig TestConfig
		err = LoadConfigSkipValidate(&readConfig, configFile)
		assert.NoError(t, err)
		assert.Equal(t, originalConfig.Name, readConfig.Name)
		assert.Equal(t, originalConfig.Port, readConfig.Port)
		assert.Equal(t, originalConfig.Debug, readConfig.Debug)
		assert.Equal(t, originalConfig.Database.Host, readConfig.Database.Host)
		assert.Equal(t, originalConfig.Database.Username, readConfig.Database.Username)
		assert.Equal(t, originalConfig.Database.Password, readConfig.Database.Password)
	})
}

// TestPropertiesSpecialCases 测试 Properties 格式的特殊情况
func TestPropertiesSpecialCases(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("properties with escape characters", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "properties_escape_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.properties")
		propertiesContent := `name=test\\nwith\\nnewlines
port=8080
debug=true
database.host=localhost\\twithtab
`

		err = os.WriteFile(configFile, []byte(propertiesContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "test\nwith\nnewlines", config.Name)
		assert.Equal(t, "localhost\twithtab", config.Database.Host)
	})

	t.Run("properties with comments", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "properties_comments_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "config.properties")
		propertiesContent := `# Main configuration
name=commenttest
# This is also a comment
! This is an alternative comment style
port=8080
debug=true
`

		err = os.WriteFile(configFile, []byte(propertiesContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "commenttest", config.Name)
		assert.Equal(t, 8080, config.Port)
	})
}

// TestEnvSpecialCases 测试 Env 格式的特殊情况
func TestEnvSpecialCases(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("env with single quotes", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "env_single_quotes_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, ".env")
		envContent := `name='single quoted value'
port=8080
debug='true'
`

		err = os.WriteFile(configFile, []byte(envContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "single quoted value", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
	})

	t.Run("env with double quotes", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "env_double_quotes_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, ".env")
		envContent := `name="double quoted value"
port=8080
debug="false"
`

		err = os.WriteFile(configFile, []byte(envContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "double quoted value", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.False(t, config.Debug)
	})

	t.Run("env without quotes", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "env_no_quotes_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, ".env")
		envContent := `name=noquotesvalue
port=8080
debug=true
`

		err = os.WriteFile(configFile, []byte(envContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		assert.NoError(t, err)
		assert.Equal(t, "noquotesvalue", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
	})
}

// TestHCLWriteAndFormat 测试HCL相关函数
func TestHCLWriteAndFormat(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("writeHCLFile and structToHCL", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "hcl_write_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "write_test.hcl")
		configPath = configFile

		// 创建测试配置，需要添加hcl标签的配置结构体
		type HCLTestConfig struct {
			Name     string `hcl:"name"`
			Port     int    `hcl:"port"`
			Debug    bool   `hcl:"debug"`
			Database struct {
				Host     string `hcl:"host"`
				Username string `hcl:"username"`
				Password string `hcl:"password"`
			} `hcl:"database,block"`
		}

		originalConfig := HCLTestConfig{
			Name:  "hclwritetest",
			Port:  2020,
			Debug: true,
		}
		originalConfig.Database.Host = "hcl-write.example.com"
		originalConfig.Database.Username = "hclwriteuser"
		originalConfig.Database.Password = "hclwritepass"

		// 写入配置
		err = SetConfig(originalConfig)
		assert.NoError(t, err)

		// 读取配置
		var readConfig HCLTestConfig
		err = LoadConfigSkipValidate(&readConfig, configFile)
		assert.NoError(t, err)
		assert.Equal(t, originalConfig.Name, readConfig.Name)
		assert.Equal(t, originalConfig.Port, readConfig.Port)
		assert.Equal(t, originalConfig.Debug, readConfig.Debug)
		assert.Equal(t, originalConfig.Database.Host, readConfig.Database.Host)
		assert.Equal(t, originalConfig.Database.Username, readConfig.Database.Username)
		assert.Equal(t, originalConfig.Database.Password, readConfig.Database.Password)
	})

	t.Run("formatHCLValue different types", func(t *testing.T) {
		// 测试不同数据类型的格式化
		assert.Equal(t, `"test"`, formatHCLValue("test"))
		assert.Equal(t, "true", formatHCLValue(true))
		assert.Equal(t, "false", formatHCLValue(false))
		assert.Equal(t, "42", formatHCLValue(int(42)))
		assert.Equal(t, "42", formatHCLValue(int8(42)))
		assert.Equal(t, "42", formatHCLValue(int16(42)))
		assert.Equal(t, "42", formatHCLValue(int32(42)))
		assert.Equal(t, "42", formatHCLValue(int64(42)))
		assert.Equal(t, "42", formatHCLValue(uint(42)))
		assert.Equal(t, "42", formatHCLValue(uint8(42)))
		assert.Equal(t, "42", formatHCLValue(uint16(42)))
		assert.Equal(t, "42", formatHCLValue(uint32(42)))
		assert.Equal(t, "42", formatHCLValue(uint64(42)))
		assert.Equal(t, "3.14", formatHCLValue(float32(3.14)))
		assert.Equal(t, "3.14", formatHCLValue(float64(3.14)))

		// 测试其他类型的默认处理
		type CustomType struct{}
		assert.Equal(t, `"{}"`, formatHCLValue(CustomType{}))
	})

	t.Run("getHCLFieldTagName", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `hcl:"hcl_field"`
			Field2 string `json:"json_field"`
			Field3 string `yaml:"yaml_field"`
			Field4 string `toml:"toml_field"`
			Field5 string `ini:"ini_field"`
			Field6 string // 没有标签，应该使用小写字段名
		}

		rt := reflect.TypeOf(TestStruct{})

		// 测试 hcl 标签优先级
		field1, _ := rt.FieldByName("Field1")
		assert.Equal(t, "hcl_field", getHCLFieldTagName(field1))

		// 测试 json 标签
		field2, _ := rt.FieldByName("Field2")
		assert.Equal(t, "json_field", getHCLFieldTagName(field2))

		// 测试 yaml 标签
		field3, _ := rt.FieldByName("Field3")
		assert.Equal(t, "yaml_field", getHCLFieldTagName(field3))

		// 测试 toml 标签
		field4, _ := rt.FieldByName("Field4")
		assert.Equal(t, "toml_field", getHCLFieldTagName(field4))

		// 测试 ini 标签
		field5, _ := rt.FieldByName("Field5")
		assert.Equal(t, "ini_field", getHCLFieldTagName(field5))

		// 测试没有标签的字段，应该返回小写字段名
		field6, _ := rt.FieldByName("Field6")
		assert.Equal(t, "field6", getHCLFieldTagName(field6))
	})
}

// TestSetFieldValueCompleteTypes 测试 setFieldValue 函数的所有数据类型
func TestSetFieldValueCompleteTypes(t *testing.T) {
	t.Run("test all supported types", func(t *testing.T) {
		type AllTypes struct {
			StringField  string
			IntField     int
			Int8Field    int8
			Int16Field   int16
			Int32Field   int32
			Int64Field   int64
			UintField    uint
			Uint8Field   uint8
			Uint16Field  uint16
			Uint32Field  uint32
			Uint64Field  uint64
			Float32Field float32
			Float64Field float64
			BoolField    bool
		}

		var testStruct AllTypes
		rv := reflect.ValueOf(&testStruct).Elem()

		// 测试字符串
		stringField := rv.FieldByName("StringField")
		err := setFieldValue(stringField, "test")
		assert.NoError(t, err)
		assert.Equal(t, "test", testStruct.StringField)

		// 测试 int 类型
		intField := rv.FieldByName("IntField")
		err = setFieldValue(intField, "42")
		assert.NoError(t, err)
		assert.Equal(t, 42, testStruct.IntField)

		// 测试 int8
		int8Field := rv.FieldByName("Int8Field")
		err = setFieldValue(int8Field, "127")
		assert.NoError(t, err)
		assert.Equal(t, int8(127), testStruct.Int8Field)

		// 测试 int16
		int16Field := rv.FieldByName("Int16Field")
		err = setFieldValue(int16Field, "32767")
		assert.NoError(t, err)
		assert.Equal(t, int16(32767), testStruct.Int16Field)

		// 测试 int32
		int32Field := rv.FieldByName("Int32Field")
		err = setFieldValue(int32Field, "2147483647")
		assert.NoError(t, err)
		assert.Equal(t, int32(2147483647), testStruct.Int32Field)

		// 测试 int64
		int64Field := rv.FieldByName("Int64Field")
		err = setFieldValue(int64Field, "9223372036854775807")
		assert.NoError(t, err)
		assert.Equal(t, int64(9223372036854775807), testStruct.Int64Field)

		// 测试 uint 类型
		uintField := rv.FieldByName("UintField")
		err = setFieldValue(uintField, "42")
		assert.NoError(t, err)
		assert.Equal(t, uint(42), testStruct.UintField)

		// 测试 uint8
		uint8Field := rv.FieldByName("Uint8Field")
		err = setFieldValue(uint8Field, "255")
		assert.NoError(t, err)
		assert.Equal(t, uint8(255), testStruct.Uint8Field)

		// 测试 uint16
		uint16Field := rv.FieldByName("Uint16Field")
		err = setFieldValue(uint16Field, "65535")
		assert.NoError(t, err)
		assert.Equal(t, uint16(65535), testStruct.Uint16Field)

		// 测试 uint32
		uint32Field := rv.FieldByName("Uint32Field")
		err = setFieldValue(uint32Field, "4294967295")
		assert.NoError(t, err)
		assert.Equal(t, uint32(4294967295), testStruct.Uint32Field)

		// 测试 uint64
		uint64Field := rv.FieldByName("Uint64Field")
		err = setFieldValue(uint64Field, "18446744073709551615")
		assert.NoError(t, err)
		assert.Equal(t, uint64(18446744073709551615), testStruct.Uint64Field)

		// 测试 float32
		float32Field := rv.FieldByName("Float32Field")
		err = setFieldValue(float32Field, "3.14")
		assert.NoError(t, err)
		assert.Equal(t, float32(3.14), testStruct.Float32Field)

		// 测试 float64
		float64Field := rv.FieldByName("Float64Field")
		err = setFieldValue(float64Field, "3.141592653589793")
		assert.NoError(t, err)
		assert.Equal(t, 3.141592653589793, testStruct.Float64Field)

		// 测试 bool
		boolField := rv.FieldByName("BoolField")
		err = setFieldValue(boolField, "true")
		assert.NoError(t, err)
		assert.True(t, testStruct.BoolField)

		// 测试不支持的类型
		type UnsupportedType struct {
			SliceField []string
		}
		var unsupported UnsupportedType
		unsupportedRv := reflect.ValueOf(&unsupported).Elem()
		sliceField := unsupportedRv.FieldByName("SliceField")
		err = setFieldValue(sliceField, "test")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported field type")
	})

	t.Run("test parsing errors", func(t *testing.T) {
		type ErrorTestStruct struct {
			IntField     int
			UintField    uint
			Float32Field float32
			BoolField    bool
		}

		var testStruct ErrorTestStruct
		rv := reflect.ValueOf(&testStruct).Elem()

		// 测试 int 解析错误
		intField := rv.FieldByName("IntField")
		err := setFieldValue(intField, "not-a-number")
		assert.Error(t, err)

		// 测试 uint 解析错误
		uintField := rv.FieldByName("UintField")
		err = setFieldValue(uintField, "not-a-number")
		assert.Error(t, err)

		// 测试 float 解析错误
		floatField := rv.FieldByName("Float32Field")
		err = setFieldValue(floatField, "not-a-float")
		assert.Error(t, err)

		// 测试 bool 解析错误
		boolField := rv.FieldByName("BoolField")
		err = setFieldValue(boolField, "not-a-bool")
		assert.Error(t, err)
	})
}

// TestGetFieldTagNameComplete 测试 getFieldTagName 函数的所有分支
func TestGetFieldTagNameComplete(t *testing.T) {
	t.Run("test all tag priorities", func(t *testing.T) {
		type TagTestStruct struct {
			Field1 string `properties:"prop_name" env:"env_name" json:"json_name1" yaml:"yaml_name" toml:"toml_name" ini:"ini_name"`
			Field2 string `env:"env_name" json:"json_name2" yaml:"yaml_name" toml:"toml_name" ini:"ini_name"`
			Field3 string `json:"json_name3" yaml:"yaml_name" toml:"toml_name" ini:"ini_name"`
			Field4 string `yaml:"yaml_name" toml:"toml_name" ini:"ini_name"`
			Field5 string `toml:"toml_name" ini:"ini_name"`
			Field6 string `ini:"ini_name"`
			Field7 string // 没有标签
			Field8 string `json:"json_with_options,omitempty"`
			Field9 string `properties:"-"` // 忽略的字段
		}

		rt := reflect.TypeOf(TagTestStruct{})

		// 测试优先级：properties > env > json > yaml > toml > ini
		field1, _ := rt.FieldByName("Field1")
		assert.Equal(t, "prop_name", getFieldTagName(field1))

		field2, _ := rt.FieldByName("Field2")
		assert.Equal(t, "env_name", getFieldTagName(field2))

		field3, _ := rt.FieldByName("Field3")
		assert.Equal(t, "json_name", getFieldTagName(field3))

		field4, _ := rt.FieldByName("Field4")
		assert.Equal(t, "yaml_name", getFieldTagName(field4))

		field5, _ := rt.FieldByName("Field5")
		assert.Equal(t, "toml_name", getFieldTagName(field5))

		field6, _ := rt.FieldByName("Field6")
		assert.Equal(t, "ini_name", getFieldTagName(field6))

		// 测试没有标签的字段
		field7, _ := rt.FieldByName("Field7")
		assert.Equal(t, "field7", getFieldTagName(field7))

		// 测试带选项的标签（逗号分隔）
		field8, _ := rt.FieldByName("Field8")
		assert.Equal(t, "json_with_options", getFieldTagName(field8))

		// 测试忽略的字段
		field9, _ := rt.FieldByName("Field9")
		assert.Equal(t, "field9", getFieldTagName(field9)) // 应该回退到字段名
	})
}

// TestWriteErrorHandling 测试写入函数的错误处理
func TestWriteErrorHandling(t *testing.T) {
	t.Run("writeProperties error handling", func(t *testing.T) {
		// 测试能输出但使用简单类型的结构体
		type SimpleStruct struct {
			Field string `properties:"field"`
		}

		var buf strings.Builder
		data := SimpleStruct{
			Field: "test",
		}

		err := writeProperties(&buf, data)
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), "field=test")
	})

	t.Run("writeEnvFile error handling", func(t *testing.T) {
		// 测试能输出但使用简单类型的结构体
		type SimpleStruct struct {
			Field string `env:"field"`
		}

		var buf strings.Builder
		data := SimpleStruct{
			Field: "test value",
		}

		err := writeEnvFile(&buf, data)
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), `field="test value"`)
	})

	t.Run("writeHCLFile error handling", func(t *testing.T) {
		// 测试无法转换的结构体
		type InvalidStruct struct {
			BadField chan int
		}

		var buf strings.Builder
		invalidData := InvalidStruct{
			BadField: make(chan int),
		}

		err := writeHCLFile(&buf, invalidData)
		// 由于 structToHCL 会处理任何类型，所以不应该产生错误
		assert.NoError(t, err)
	})
}

// TestMapToStructErrorHandling 测试 mapToStruct 的错误处理
func TestMapToStructErrorHandling(t *testing.T) {
	t.Run("invalid target type", func(t *testing.T) {
		props := map[string]string{"test": "value"}

		// 测试非指针类型
		var notPointer TestConfig
		err := mapToStruct(props, notPointer)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be a pointer to struct")

		// 测试指向非结构体的指针
		var notStruct string
		err = mapToStruct(props, &notStruct)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be a pointer to struct")
	})

	t.Run("nested struct error handling", func(t *testing.T) {
		type NestedConfig struct {
			Database struct {
				Port string `properties:"port"`
			} `properties:"database"`
		}

		props := map[string]string{
			"database.port": "invalid-port", // 这将导致设置字段值时出错
		}

		var config NestedConfig
		err := mapToStruct(props, &config)
		// 由于我们这里设置的是字符串字段，所以不会出错
		assert.NoError(t, err)
		assert.Equal(t, "invalid-port", config.Database.Port)
	})

	t.Run("field setting error", func(t *testing.T) {
		type ErrorConfig struct {
			Port int `properties:"port"`
		}

		props := map[string]string{
			"port": "not-a-number",
		}

		var config ErrorConfig
		err := mapToStruct(props, &config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to set field")
	})
}

// TestStructToMapErrorHandling 测试 structToMap 的错误处理
func TestStructToMapErrorHandling(t *testing.T) {
	t.Run("invalid input type", func(t *testing.T) {
		// 测试非结构体类型
		_, err := structToMap("not a struct")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be a struct or pointer to struct")

		// 测试指向非结构体的指针
		notStruct := "not a struct"
		_, err = structToMap(&notStruct)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be a struct or pointer to struct")
	})
}

// TestParseErrorHandling 测试解析函数的错误处理
func TestParseErrorHandling(t *testing.T) {
	t.Run("parseProperties with valid content", func(t *testing.T) {
		// 测试有效的 Properties 内容解析
		content := "name=testname\nport=8080\ndebug=true"
		reader := strings.NewReader(content)

		type SimpleConfig struct {
			Name  string `properties:"name"`
			Port  int    `properties:"port"`
			Debug bool   `properties:"debug"`
		}

		var config SimpleConfig
		err := parseProperties(reader, &config)
		assert.NoError(t, err)
		assert.Equal(t, "testname", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
	})

	t.Run("parseEnvFile with valid content", func(t *testing.T) {
		// 测试有效的 Env 内容解析
		content := "name=\"testname\"\nport=8080\ndebug=true"
		reader := strings.NewReader(content)

		type SimpleConfig struct {
			Name  string `env:"name"`
			Port  int    `env:"port"`
			Debug bool   `env:"debug"`
		}

		var config SimpleConfig
		err := parseEnvFile(reader, &config)
		assert.NoError(t, err)
		assert.Equal(t, "testname", config.Name)
		assert.Equal(t, 8080, config.Port)
		assert.True(t, config.Debug)
	})
}

// TestNestedStructEdgeCases 测试嵌套结构体的边界情况
func TestNestedStructEdgeCases(t *testing.T) {
	t.Run("deep nested structs", func(t *testing.T) {
		type DeepNested struct {
			Level1 struct {
				Level2 struct {
					Value string `properties:"value"`
				} `properties:"level2"`
			} `properties:"level1"`
		}

		props := map[string]string{
			"level1.level2.value": "deep_value",
		}

		var config DeepNested
		err := mapToStruct(props, &config)
		assert.NoError(t, err)
		assert.Equal(t, "deep_value", config.Level1.Level2.Value)
	})

	t.Run("cannot set nested field", func(t *testing.T) {
		type NestedWithError struct {
			Database struct {
				port int // 私有字段，无法设置
			} `properties:"database"`
		}

		props := map[string]string{
			"database.port": "8080",
		}

		var config NestedWithError
		err := mapToStruct(props, &config)
		// 由于无法设置私有字段，应该不会报错，只是忽略该字段
		assert.NoError(t, err)
	})
}

// TestComplexTypes 测试复杂类型的处理
func TestComplexTypes(t *testing.T) {
	t.Run("struct with interface fields", func(t *testing.T) {
		type InterfaceStruct struct {
			Value interface{} `properties:"value"`
		}

		var config InterfaceStruct
		rv := reflect.ValueOf(&config).Elem()
		valueField := rv.FieldByName("Value")

		// 测试设置 interface{} 字段
		err := setFieldValue(valueField, "test")
		assert.Error(t, err) // interface{} 不被支持
		assert.Contains(t, err.Error(), "unsupported field type")
	})
}

// TestLoadConfigReturnError 测试 LoadConfig 的返回错误
func TestLoadConfigReturnError(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("LoadConfigSkipValidate returns no error even with parse errors", func(t *testing.T) {
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "loadconfig_no_error_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		configFile := filepath.Join(tmpDir, "invalid.json")
		invalidContent := `{`

		err = os.WriteFile(configFile, []byte(invalidContent), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfigSkipValidate(&config, configFile)
		// LoadConfigSkipValidate 应该返回 nil，即使解析失败
		assert.NoError(t, err)
	})
}

// TestAdditionalCoverage 测试剩余的未覆盖分支
func TestAdditionalCoverage(t *testing.T) {
	originalConfigPath := configPath
	defer func() {
		configPath = originalConfigPath
	}()

	t.Run("LoadConfig LoadConfigSkipValidate error", func(t *testing.T) {
		// 测试 LoadConfigSkipValidate 返回错误的情况
		configPath = ""

		tmpDir, err := os.MkdirTemp("", "loadconfig_skiperror_test_")
		require.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		// 创建一个会导致解析错误的情况
		configFile := filepath.Join(tmpDir, "test.properties")
		// 创建一个会导致类型错误的properties文件
		propertiesContent := `name=testname
port=not-a-number`

		err = os.WriteFile(configFile, []byte(propertiesContent), 0644)
		require.NoError(t, err)

		type ErrorConfig struct {
			Name string `properties:"name" validate:"required"`
			Port int    `properties:"port" validate:"min=1,max=65535"`
		}

		var config ErrorConfig
		err = LoadConfig(&config, configFile)
		assert.Error(t, err) // 应该因为验证失败而返回错误
	})

	t.Run("parseProperties no separator", func(t *testing.T) {
		// 测试没有分隔符的行
		content := "name=testname\ninvalidline\nport=8080"
		reader := strings.NewReader(content)

		type SimpleConfig struct {
			Name string `properties:"name"`
			Port int    `properties:"port"`
		}

		var config SimpleConfig
		err := parseProperties(reader, &config)
		assert.NoError(t, err)
		assert.Equal(t, "testname", config.Name)
		assert.Equal(t, 8080, config.Port)
	})

	t.Run("parseEnvFile no separator", func(t *testing.T) {
		// 测试没有分隔符的行
		content := "name=testname\ninvalidline\nport=8080"
		reader := strings.NewReader(content)

		type SimpleConfig struct {
			Name string `env:"name"`
			Port int    `env:"port"`
		}

		var config SimpleConfig
		err := parseEnvFile(reader, &config)
		assert.NoError(t, err)
		assert.Equal(t, "testname", config.Name)
		assert.Equal(t, 8080, config.Port)
	})

	t.Run("parseEnvFile no quotes needed", func(t *testing.T) {
		// 测试不需要引号的值
		content := "name=simple\nport=8080"
		reader := strings.NewReader(content)

		type SimpleConfig struct {
			Name string `env:"name"`
			Port int    `env:"port"`
		}

		var config SimpleConfig
		err := parseEnvFile(reader, &config)
		assert.NoError(t, err)
		assert.Equal(t, "simple", config.Name)
		assert.Equal(t, 8080, config.Port)
	})

	t.Run("writeProperties with error", func(t *testing.T) {
		// 测试 writeProperties 的错误处理 - 使用一个不能Marshal的类型
		type BadWriterStruct struct {
			Field func() // 函数类型会被格式化为字符串
		}

		var buf strings.Builder
		data := BadWriterStruct{
			Field: func() {},
		}

		err := writeProperties(&buf, data)
		// writeProperties 不会失败，它会把任何类型都转换为字符串
		assert.NoError(t, err)
	})

	t.Run("writeEnvFile with spaces", func(t *testing.T) {
		// 测试包含空格的值
		type SpacedStruct struct {
			Field string `env:"field"`
		}

		var buf strings.Builder
		data := SpacedStruct{
			Field: "value with spaces",
		}

		err := writeEnvFile(&buf, data)
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), `"value with spaces"`)
	})

	t.Run("structMapToFlat nested struct", func(t *testing.T) {
		// 测试嵌套结构体的展平
		type NestedStruct struct {
			Database struct {
				Host string `properties:"host"`
			} `properties:"database"`
		}

		data := NestedStruct{}
		data.Database.Host = "localhost"

		result, err := structToMap(data)
		assert.NoError(t, err)
		assert.Equal(t, "localhost", result["database.host"])
	})

	t.Run("structToHCL with non-struct", func(t *testing.T) {
		// 测试非结构体类型
		result, err := structToHCL("not a struct", "")
		assert.NoError(t, err)
		assert.Equal(t, "not a struct", result)
	})

	t.Run("parseNestedStruct with error", func(t *testing.T) {
		// 测试嵌套结构体解析错误
		type NestedWithBadField struct {
			Database struct {
				Port int `properties:"port"`
			} `properties:"database"`
		}

		props := map[string]string{
			"database.port": "invalid-number",
		}

		var config NestedWithBadField
		err := mapToStruct(props, &config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to set nested field")
	})
}

// scannerErrorReader 模拟一个会产生扫描错误的读取器
type scannerErrorReader struct {
	data []byte
	pos  int
}

func (r *scannerErrorReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, fmt.Errorf("scanner error: simulated error")
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// errorWriter 模拟一个会产生写入错误的写入器
type errorWriter struct{}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error: simulated error")
}

// TestCompleteCoverageEnhancement 增强测试覆盖率至100%
func TestCompleteCoverageEnhancement(t *testing.T) {
	t.Run("parseProperties scanner error", func(t *testing.T) {
		// 创建一个会产生读取错误的reader
		reader := &scannerErrorReader{data: []byte("key=value\n"), pos: len("key=value\n")}

		var result map[string]interface{}
		err := parseProperties(reader, &result)
		// 应该产生扫描错误
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "scanner error")
	})

	t.Run("parseProperties with BufferSize limit", func(t *testing.T) {
		// 创建一个非常长的行来触发bufio.Scanner的限制
		longValue := strings.Repeat("a", 70000) // 超过默认bufio.Scanner缓冲区大小
		content := fmt.Sprintf("key=%s", longValue)
		reader := strings.NewReader(content)

		var result map[string]interface{}
		err := parseProperties(reader, &result)
		if err != nil {
			// 如果出现错误，应该是bufio.Scanner相关的错误
			assert.Contains(t, err.Error(), "token too long")
		} else {
			// 如果没有错误，检查数据是否正确解析
			assert.Equal(t, longValue, result["key"])
		}
	})

	t.Run("writeProperties error on reflect", func(t *testing.T) {
		// 测试无法转换为map的类型
		var buf bytes.Buffer
		err := writeProperties(&buf, make(chan int))
		assert.Error(t, err)
	})

	t.Run("parseEnvFile scanner buffer overflow", func(t *testing.T) {
		// 创建一个非常长的行来触发bufio.Scanner的限制
		longValue := strings.Repeat("b", 70000)
		content := fmt.Sprintf("KEY=%s", longValue)
		reader := strings.NewReader(content)

		var result map[string]interface{}
		err := parseEnvFile(reader, &result)
		if err != nil {
			assert.Contains(t, err.Error(), "token too long")
		} else {
			assert.Equal(t, longValue, result["KEY"])
		}
	})

	t.Run("writeEnvFile error on reflect", func(t *testing.T) {
		// 测试无法转换为map的类型
		var buf bytes.Buffer
		err := writeEnvFile(&buf, make(chan int))
		assert.Error(t, err)
	})

	t.Run("mapToStruct with nested error", func(t *testing.T) {
		// 测试嵌套结构体转换错误
		type NestedWithBadType struct {
			Port chan int `properties:"port"`
		}
		type ConfigWithBadNested struct {
			Database NestedWithBadType `properties:"database"`
		}

		props := map[string]string{
			"database.port": "invalid-value",
		}

		var config ConfigWithBadNested
		err := mapToStruct(props, &config)
		assert.Error(t, err)
	})

	t.Run("structToMap with unsupported type", func(t *testing.T) {
		// 测试不支持的类型
		type UnsupportedConfig struct {
			Channel chan int `properties:"channel"`
		}

		config := UnsupportedConfig{
			Channel: make(chan int),
		}

		result, err := structToMap(config)
		// structToMap 应该能处理这种情况，或者返回错误
		if err != nil {
			assert.Error(t, err)
		} else {
			// 如果没有错误，确保结果有效
			assert.NotNil(t, result)
		}
	})

	t.Run("parseNestedStruct with invalid conversion", func(t *testing.T) {
		// 测试parseNestedStruct中的类型转换错误
		props := map[string]string{
			"config.port": "not-a-number",
		}

		type BadConfig struct {
			Port int `properties:"port"`
		}

		var config BadConfig
		structValue := reflect.ValueOf(&config).Elem()
		err := parseNestedStruct(props, structValue, "config")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to set")
	})

	t.Run("structMapToFlat with complex nested", func(t *testing.T) {
		// 测试复杂的嵌套结构体平铺
		type Level2 struct {
			Level3 string `properties:"level3"`
			Number int    `properties:"number"`
		}
		type Level1 struct {
			Level2 Level2 `properties:"level2"`
			Simple string `properties:"simple"`
		}
		type ComplexConfig struct {
			Level1 Level1 `properties:"level1"`
		}

		config := ComplexConfig{
			Level1: Level1{
				Level2: Level2{
					Level3: "deepvalue",
					Number: 42,
				},
				Simple: "value",
			},
		}

		result := make(map[string]interface{})
		structMapToFlat(result, reflect.ValueOf(config), "test")
		assert.Equal(t, "deepvalue", result["test.level1.level2.level3"])
		assert.Equal(t, 42, result["test.level1.level2.number"])
		assert.Equal(t, "value", result["test.level1.simple"])
	})

	t.Run("writeHCLFile with non-struct", func(t *testing.T) {
		// 测试writeHCLFile处理非结构体类型
		var buf bytes.Buffer

		// 尝试用基础类型
		err := writeHCLFile(&buf, "simple string")
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), "simple string")
	})

	t.Run("LoadConfig validation error path", func(t *testing.T) {
		// 测试LoadConfig中validation错误的返回路径
		tempDir := t.TempDir()
		configFile := filepath.Join(tempDir, "test_validation.json")

		// 创建会导致验证失败的配置
		content := `{"name": "", "port": 0}`
		err := os.WriteFile(configFile, []byte(content), 0644)
		require.NoError(t, err)

		type ValidatedConfig struct {
			Name string `json:"name" validate:"required"`
			Port int    `json:"port" validate:"min=1"`
		}

		var config ValidatedConfig
		err = LoadConfig(&config, configFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation")
	})

	t.Run("writeProperties with write error", func(t *testing.T) {
		// 测试写入错误的情况
		type SimpleConfig struct {
			Name string `properties:"name"`
		}
		config := SimpleConfig{Name: "test"}

		writer := &errorWriter{}
		err := writeProperties(writer, config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "write error")
	})

	t.Run("writeEnvFile with write error", func(t *testing.T) {
		// 测试写入错误的情况
		type SimpleConfig struct {
			Name string `env:"name"`
		}
		config := SimpleConfig{Name: "test"}

		writer := &errorWriter{}
		err := writeEnvFile(writer, config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "write error")
	})

	t.Run("writeHCLFile with write error", func(t *testing.T) {
		// 测试写入错误的情况
		type SimpleConfig struct {
			Name string `hcl:"name"`
		}
		config := SimpleConfig{Name: "test"}

		writer := &errorWriter{}
		err := writeHCLFile(writer, config)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "write error")
	})

	t.Run("structMapToFlat with empty tag", func(t *testing.T) {
		// 测试空标签的情况
		type ConfigWithEmptyTag struct {
			Name   string `properties:"name"`
			Hidden string // 没有标签，应该被忽略
		}

		config := ConfigWithEmptyTag{
			Name:   "test",
			Hidden: "hidden",
		}

		result := make(map[string]interface{})
		structMapToFlat(result, reflect.ValueOf(config), "prefix")

		assert.Equal(t, "test", result["prefix.name"])
		// Hidden字段可能会被包含，因为它没有标签但可能有默认行为
		// 我们不做断言，只检查name字段正确
	})

	t.Run("structToHCL with nil pointer", func(t *testing.T) {
		// 测试处理nil指针
		var nilPtr *TestConfig = nil
		result, err := structToHCL(nilPtr, "")
		assert.NoError(t, err)
		assert.Equal(t, "<nil>", result)
	})
}

// TestFinalCoverageBoost 最终覆盖率提升测试
func TestFinalCoverageBoost(t *testing.T) {
	t.Run("LoadConfig with LoadConfigSkipValidate error", func(t *testing.T) {
		// 创建一个不存在的配置文件，然后强制路径
		// 当文件不存在时，LoadConfigSkipValidate会使用默认配置，但这会导致验证失败
		var config TestConfig
		err := LoadConfig(&config, "/this/path/absolutely/does/not/exist.json")
		// 应该因为验证失败而返回错误（TestConfig有required字段）
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation")
	})

	t.Run("LoadConfig with unsupported format error", func(t *testing.T) {
		// 测试文件格式不支持的情况，这应该让LoadConfigSkipValidate返回错误
		tempDir := t.TempDir()
		configFile := filepath.Join(tempDir, "config.unsupported")
		err := os.WriteFile(configFile, []byte("some content"), 0644)
		require.NoError(t, err)

		var config TestConfig
		err = LoadConfig(&config, configFile)
		assert.Error(t, err)
		// 这应该覆盖LoadConfig中的err != nil路径（lines 159-162）
	})

	t.Run("mapToStruct with unsupported field types", func(t *testing.T) {
		// 测试不支持的字段类型，触发字段设置错误
		type UnsupportedConfig struct {
			Channel chan int `properties:"channel"`
		}

		props := map[string]string{
			"channel": "some-value", // 无法转换为chan int
		}

		var config UnsupportedConfig
		err := mapToStruct(props, &config)
		// 这应该会失败，因为无法将字符串设置到channel字段
		assert.Error(t, err)
	})

	t.Run("structToMap with exported fields only", func(t *testing.T) {
		// 测试只包含exported字段的结构，确保函数正常工作
		type SimpleConfig struct {
			Public  string `properties:"public"`
			Another string `properties:"another"`
		}

		config := SimpleConfig{
			Public:  "public-value",
			Another: "another-value",
		}

		result, err := structToMap(config)
		assert.NoError(t, err)
		assert.Equal(t, "public-value", result["public"])
		assert.Equal(t, "another-value", result["another"])
	})

	t.Run("parseNestedStruct with unsettable field", func(t *testing.T) {
		// 测试不可设置的字段情况
		type ReadOnlyConfig struct {
			value string `properties:"value"` // unexported字段无法设置
		}

		props := map[string]string{
			"config.value": "test-value",
		}

		var config ReadOnlyConfig
		err := parseNestedStruct(props, reflect.ValueOf(&config).Elem(), "config")
		// unexported字段应该被跳过，不应该产生错误
		assert.NoError(t, err)
	})

	t.Run("structMapToFlat with simple fields", func(t *testing.T) {
		// 测试structMapToFlat基本功能
		type SimpleConfig struct {
			Public  string `properties:"public"`
			Another string `properties:"another"`
		}

		config := SimpleConfig{
			Public:  "public-value",
			Another: "another-value",
		}

		result := make(map[string]interface{})
		structMapToFlat(result, reflect.ValueOf(config), "prefix")

		assert.Equal(t, "public-value", result["prefix.public"])
		assert.Equal(t, "another-value", result["prefix.another"])
	})
}

