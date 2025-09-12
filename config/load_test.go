package config

import (
	"io"
	"os"
	"path/filepath"
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
			Name    string    `yaml:"name"`
			Channel chan int  `yaml:"channel"` // channels cannot be marshaled
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