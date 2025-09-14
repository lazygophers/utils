package fake

import (
	"embed"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"sync"
)

//go:embed data
var dataFS embed.FS

// DataItem 数据项结构
type DataItem struct {
	Value  string            `json:"value"`
	Weight float64           `json:"weight,omitempty"`
	Tags   []string          `json:"tags,omitempty"`
	Meta   map[string]string `json:"meta,omitempty"`
}

// DataSet 数据集结构
type DataSet struct {
	Language string     `json:"language"`
	Country  string     `json:"country,omitempty"`
	Type     string     `json:"type"`
	Items    []DataItem `json:"items"`
	Version  string     `json:"version,omitempty"`
}

// DataManager 数据管理器
type DataManager struct {
	cache sync.Map // map[string]*DataSet
	mu    sync.RWMutex
}

var (
	dataManager     *DataManager
	dataManagerOnce sync.Once
)

// getDataManager 获取全局数据管理器实例
func getDataManager() *DataManager {
	dataManagerOnce.Do(func() {
		dataManager = &DataManager{}
	})
	return dataManager
}

// LoadDataSet 加载指定的数据集
func (dm *DataManager) LoadDataSet(language Language, dataType, subType string) (*DataSet, error) {
	key := fmt.Sprintf("%s:%s:%s", language, dataType, subType)
	
	// 尝试从缓存获取
	if cached, ok := dm.cache.Load(key); ok {
		return cached.(*DataSet), nil
	}
	
	// 构建文件路径
	filePath := path.Join("data", string(language), dataType, subType+".json")
	
	// 读取嵌入的文件
	data, err := dataFS.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file %s: %w", filePath, err)
	}
	
	// 解析JSON
	var dataSet DataSet
	if err := json.Unmarshal(data, &dataSet); err != nil {
		return nil, fmt.Errorf("failed to parse data file %s: %w", filePath, err)
	}
	
	// 存入缓存
	dm.cache.Store(key, &dataSet)
	
	return &dataSet, nil
}

// GetItems 获取指定数据集的所有条目
func (dm *DataManager) GetItems(language Language, dataType, subType string) ([]DataItem, error) {
	dataSet, err := dm.LoadDataSet(language, dataType, subType)
	if err != nil {
		return nil, err
	}
	
	return dataSet.Items, nil
}

// GetItemValues 获取指定数据集的所有值
func (dm *DataManager) GetItemValues(language Language, dataType, subType string) ([]string, error) {
	items, err := dm.GetItems(language, dataType, subType)
	if err != nil {
		return nil, err
	}
	
	values := make([]string, len(items))
	for i, item := range items {
		values[i] = item.Value
	}
	
	return values, nil
}

// GetWeightedItems 获取带权重的数据项
func (dm *DataManager) GetWeightedItems(language Language, dataType, subType string) ([]string, []float64, error) {
	items, err := dm.GetItems(language, dataType, subType)
	if err != nil {
		return nil, nil, err
	}
	
	values := make([]string, len(items))
	weights := make([]float64, len(items))
	
	for i, item := range items {
		values[i] = item.Value
		weight := item.Weight
		if weight == 0 {
			weight = 1.0 // 默认权重
		}
		weights[i] = weight
	}
	
	return values, weights, nil
}

// GetItemsByTag 根据标签获取数据项
func (dm *DataManager) GetItemsByTag(language Language, dataType, subType, tag string) ([]DataItem, error) {
	items, err := dm.GetItems(language, dataType, subType)
	if err != nil {
		return nil, err
	}
	
	var filtered []DataItem
	for _, item := range items {
		for _, itemTag := range item.Tags {
			if itemTag == tag {
				filtered = append(filtered, item)
				break
			}
		}
	}
	
	return filtered, nil
}

// ClearCache 清空数据缓存
func (dm *DataManager) ClearCache() {
	dm.cache.Range(func(key, value interface{}) bool {
		dm.cache.Delete(key)
		return true
	})
}

// ListAvailableDataSets 列出可用的数据集
func (dm *DataManager) ListAvailableDataSets() ([]string, error) {
	var dataSets []string
	
	// 读取所有支持的语言目录
	for _, lang := range GetSupportedLanguages() {
		langDir := path.Join("data", string(lang))
		
		entries, err := dataFS.ReadDir(langDir)
		if err != nil {
			continue // 跳过不存在的语言目录
		}
		
		for _, entry := range entries {
			if entry.IsDir() {
				dataType := entry.Name()
				
				// 读取数据类型目录
				typeDir := path.Join(langDir, dataType)
				subEntries, err := dataFS.ReadDir(typeDir)
				if err != nil {
					continue
				}
				
				for _, subEntry := range subEntries {
					if !subEntry.IsDir() && strings.HasSuffix(subEntry.Name(), ".json") {
						subType := strings.TrimSuffix(subEntry.Name(), ".json")
						dataSetName := fmt.Sprintf("%s:%s:%s", lang, dataType, subType)
						dataSets = append(dataSets, dataSetName)
					}
				}
			}
		}
	}
	
	return dataSets, nil
}

// PreloadData 预加载指定语言的所有数据
func (dm *DataManager) PreloadData(language Language) error {
	langDir := path.Join("data", string(language))
	
	entries, err := dataFS.ReadDir(langDir)
	if err != nil {
		return fmt.Errorf("language %s not supported: %w", language, err)
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			dataType := entry.Name()
			
			typeDir := path.Join(langDir, dataType)
			subEntries, err := dataFS.ReadDir(typeDir)
			if err != nil {
				continue
			}
			
			for _, subEntry := range subEntries {
				if !subEntry.IsDir() && strings.HasSuffix(subEntry.Name(), ".json") {
					subType := strings.TrimSuffix(subEntry.Name(), ".json")
					
					// 预加载数据集
					_, err := dm.LoadDataSet(language, dataType, subType)
					if err != nil {
						return fmt.Errorf("failed to preload %s:%s:%s: %w", language, dataType, subType, err)
					}
				}
			}
		}
	}
	
	return nil
}

// GetCacheStats 获取缓存统计信息
func (dm *DataManager) GetCacheStats() map[string]int {
	count := 0
	dm.cache.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	
	return map[string]int{
		"cached_datasets": count,
	}
}

// 快速访问函数
func loadDataSet(language Language, dataType, subType string) (*DataSet, error) {
	return getDataManager().LoadDataSet(language, dataType, subType)
}

func getItemValues(language Language, dataType, subType string) ([]string, error) {
	return getDataManager().GetItemValues(language, dataType, subType)
}

func getWeightedItems(language Language, dataType, subType string) ([]string, []float64, error) {
	return getDataManager().GetWeightedItems(language, dataType, subType)
}

func getItemsByTag(language Language, dataType, subType, tag string) ([]DataItem, error) {
	return getDataManager().GetItemsByTag(language, dataType, subType, tag)
}