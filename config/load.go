package config

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils"
	"github.com/lazygophers/utils/json"
	"github.com/lazygophers/utils/osx"
	"github.com/lazygophers/utils/runtime"
	"github.com/pelletier/go-toml/v2"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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
	".ini": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			cfg, err := ini.Load(reader)
			if err != nil {
				return err
			}
			return cfg.MapTo(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			cfg := ini.Empty()
			err := cfg.ReflectFrom(v)
			if err != nil {
				return err
			}
			_, err = cfg.WriteTo(writer)
			return err
		},
	},
	".xml": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return xml.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			encoder := xml.NewEncoder(writer)
			encoder.Indent("", "  ")
			return encoder.Encode(v)
		},
	},
	".properties": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return parseProperties(reader, v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return writeProperties(writer, v)
		},
	},
	".env": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return parseEnvFile(reader, v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return writeEnvFile(writer, v)
		},
	},
	".hcl": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			content, err := io.ReadAll(reader)
			if err != nil {
				return err
			}
			return hclsimple.Decode("config.hcl", content, nil, v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return writeHCLFile(writer, v)
		},
	},
	".json5": {
		Unmarshaler: func(reader io.Reader, v interface{}) error {
			return json5.NewDecoder(reader).Decode(v)
		},
		Marshaler: func(writer io.Writer, v interface{}) error {
			return json5.NewEncoder(writer).Encode(v)
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

func LoadConfig(c any, paths ...string) (err error) {
	err = LoadConfigSkipValidate(c, paths...)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	err = utils.Validate(c)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	log.Info("load config success")

	return nil
}

func LoadConfigSkipValidate(c any, paths ...string) error {
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
	if configPath == "" {
		log.Warnf("Try to load config from environment variable(LAZYGOPHERS_CONFIG)")
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

	//// NOTE: 从用户目录中获取
	//if configPath == "" {
	//	log.Warnf("Try to load config from %s", runtime.UserHomeDir())
	//	configPath = tryFindConfigPath(filepath.Join(runtime.UserHomeDir(), app.Name))
	//}
	//
	//// NOTE: 从系统目录中获取
	//if configPath == "" {
	//	log.Warnf("Try to load config from %s", runtime.UserConfigDir())
	//	configPath = tryFindConfigPath(filepath.Join(runtime.UserConfigDir(), app.Name))
	//}

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
			return fmt.Errorf("unsupported config file format:%v", ext)
		}
	}

	log.Info("load config success")

	return nil
}

func SetConfig(c any) error {
	ext := filepath.Ext(configPath)
	if supported, ok := supportedExtMap[ext]; ok {
		file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
		defer file.Close()

		err = supported.Marshaler(file, c)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil
		}
	} else {
		log.Errorf("unsupported config file format:%v", ext)
		return fmt.Errorf("unsupported config file format:%v", ext)
	}

	return nil
}

// parseProperties 解析 .properties 文件格式
func parseProperties(reader io.Reader, v interface{}) error {
	props := make(map[string]string)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") {
			continue
		}

		// 找到第一个 = 或 : 作为分隔符
		sepIndex := -1
		for i, char := range line {
			if char == '=' || char == ':' {
				sepIndex = i
				break
			}
		}

		if sepIndex == -1 {
			continue
		}

		key := strings.TrimSpace(line[:sepIndex])
		value := strings.TrimSpace(line[sepIndex+1:])

		// 处理转义字符
		value = strings.ReplaceAll(value, "\\\\", "\\")
		value = strings.ReplaceAll(value, "\\n", "\n")
		value = strings.ReplaceAll(value, "\\t", "\t")
		value = strings.ReplaceAll(value, "\\r", "\r")

		props[key] = value
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return mapToStruct(props, v)
}

// writeProperties 写入 .properties 文件格式
func writeProperties(writer io.Writer, v interface{}) error {
	props, err := structToMap(v)
	if err != nil {
		return err
	}

	for key, value := range props {
		// 转义特殊字符
		valueStr := fmt.Sprintf("%v", value)
		valueStr = strings.ReplaceAll(valueStr, "\n", "\\n")
		valueStr = strings.ReplaceAll(valueStr, "\t", "\\t")
		valueStr = strings.ReplaceAll(valueStr, "\r", "\\r")

		_, err := fmt.Fprintf(writer, "%s=%s\n", key, valueStr)
		if err != nil {
			return err
		}
	}

	return nil
}

// parseEnvFile 解析 .env 文件格式
func parseEnvFile(reader io.Reader, v interface{}) error {
	props := make(map[string]string)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 找到第一个 = 作为分隔符
		sepIndex := strings.Index(line, "=")
		if sepIndex == -1 {
			continue
		}

		key := strings.TrimSpace(line[:sepIndex])
		value := strings.TrimSpace(line[sepIndex+1:])

		// 处理引号包围的值
		if len(value) >= 2 {
			if (value[0] == '"' && value[len(value)-1] == '"') ||
				(value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			}
		}

		props[key] = value
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return mapToStruct(props, v)
}

// writeEnvFile 写入 .env 文件格式
func writeEnvFile(writer io.Writer, v interface{}) error {
	props, err := structToMap(v)
	if err != nil {
		return err
	}

	for key, value := range props {
		valueStr := fmt.Sprintf("%v", value)

		// 如果值包含空格、特殊字符，则用双引号包围
		if strings.ContainsAny(valueStr, " \t\n\r\"'\\") {
			valueStr = strconv.Quote(valueStr)
		}

		_, err := fmt.Fprintf(writer, "%s=%s\n", key, valueStr)
		if err != nil {
			return err
		}
	}

	return nil
}

// mapToStruct 将 map 转换为结构体
func mapToStruct(props map[string]string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to struct")
	}

	rv = rv.Elem()
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		// 获取字段标签名
		tagName := getFieldTagName(field)
		if tagName == "" {
			continue
		}

		// 检查是否为嵌套结构体
		if fieldValue.Kind() == reflect.Struct {
			err := parseNestedStruct(props, fieldValue, tagName)
			if err != nil {
				return err
			}
			continue
		}

		// 从 map 中获取值
		if propValue, exists := props[tagName]; exists {
			err := setFieldValue(fieldValue, propValue)
			if err != nil {
				return fmt.Errorf("failed to set field %s: %w", field.Name, err)
			}
		}
	}

	return nil
}

// structToMap 将结构体转换为 map
func structToMap(v interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("v must be a struct or pointer to struct")
	}

	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		tagName := getFieldTagName(field)
		if tagName == "" {
			continue
		}

		if fieldValue.Kind() == reflect.Struct {
			// 处理嵌套结构体
			structMapToFlat(result, fieldValue, tagName)
			continue
		}

		result[tagName] = fieldValue.Interface()
	}

	return result, nil
}

// getFieldTagName 获取字段的标签名
func getFieldTagName(field reflect.StructField) string {
	// 优先级: properties > env > json > yaml > toml > ini
	tags := []string{"properties", "env", "json", "yaml", "toml", "ini"}

	for _, tag := range tags {
		if tagValue := field.Tag.Get(tag); tagValue != "" && tagValue != "-" {
			// 处理 json:",omitempty" 这样的格式
			if commaIndex := strings.Index(tagValue, ","); commaIndex != -1 {
				return tagValue[:commaIndex]
			}
			return tagValue
		}
	}

	return strings.ToLower(field.Name)
}

// setFieldValue 设置字段值
func setFieldValue(fieldValue reflect.Value, propValue string) error {
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(propValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(propValue, 10, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(propValue, 10, 64)
		if err != nil {
			return err
		}
		fieldValue.SetUint(val)
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(propValue, 64)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(propValue)
		if err != nil {
			return err
		}
		fieldValue.SetBool(val)
	default:
		return fmt.Errorf("unsupported field type: %s", fieldValue.Kind())
	}

	return nil
}

// parseNestedStruct 解析嵌套结构体
func parseNestedStruct(props map[string]string, structValue reflect.Value, prefix string) error {
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		tagName := getFieldTagName(field)
		if tagName == "" {
			continue
		}

		fullKey := prefix + "." + tagName

		if fieldValue.Kind() == reflect.Struct {
			err := parseNestedStruct(props, fieldValue, fullKey)
			if err != nil {
				return err
			}
		} else if propValue, exists := props[fullKey]; exists {
			err := setFieldValue(fieldValue, propValue)
			if err != nil {
				return fmt.Errorf("failed to set nested field %s: %w", field.Name, err)
			}
		}
	}

	return nil
}

// structMapToFlat 将嵌套结构体展平到 map 中
func structMapToFlat(result map[string]interface{}, structValue reflect.Value, prefix string) {
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		tagName := getFieldTagName(field)
		if tagName == "" {
			continue
		}

		fullKey := prefix + "." + tagName

		if fieldValue.Kind() == reflect.Struct {
			structMapToFlat(result, fieldValue, fullKey)
		} else {
			result[fullKey] = fieldValue.Interface()
		}
	}
}

// writeHCLFile 写入 HCL 文件格式
func writeHCLFile(writer io.Writer, v interface{}) error {
	hclContent, err := structToHCL(v, "")
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(hclContent))
	return err
}

// structToHCL 将结构体转换为 HCL 格式字符串
func structToHCL(v interface{}, indent string) (string, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v), nil
	}

	rt := rv.Type()
	var result strings.Builder

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		tagName := getHCLFieldTagName(field)
		if tagName == "" {
			continue
		}

		if fieldValue.Kind() == reflect.Struct {
			nestedContent, err := structToHCL(fieldValue.Interface(), indent+"  ")
			if err != nil {
				return "", err
			}
			result.WriteString(fmt.Sprintf("%s%s {\n%s%s}\n", indent, tagName, nestedContent, indent))
		} else {
			value := formatHCLValue(fieldValue.Interface())
			result.WriteString(fmt.Sprintf("%s%s = %s\n", indent, tagName, value))
		}
	}

	return result.String(), nil
}

// getHCLFieldTagName 获取 HCL 字段的标签名
func getHCLFieldTagName(field reflect.StructField) string {
	// 优先级: hcl > json > yaml > toml > ini
	tags := []string{"hcl", "json", "yaml", "toml", "ini"}

	for _, tag := range tags {
		if tagValue := field.Tag.Get(tag); tagValue != "" && tagValue != "-" {
			if commaIndex := strings.Index(tagValue, ","); commaIndex != -1 {
				return tagValue[:commaIndex]
			}
			return tagValue
		}
	}

	return strings.ToLower(field.Name)
}

// formatHCLValue 格式化 HCL 值
func formatHCLValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf(`"%s"`, v)
	case bool:
		return fmt.Sprintf("%t", v)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%g", v)
	default:
		return fmt.Sprintf(`"%v"`, v)
	}
}
