package fake

import (
	"fmt"
	"sync"
)

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
	registered map[string]*DataSet
	mu         sync.RWMutex
}

var (
	dataManager     *DataManager
	dataManagerOnce sync.Once
)

// getDataManager 获取全局数据管理器实例
func getDataManager() *DataManager {
	dataManagerOnce.Do(func() {
		dataManager = &DataManager{
			registered: make(map[string]*DataSet),
		}
	})
	return dataManager
}

// registerDataSet registers a data set directly (used by data_xx.go init functions)
func registerDataSet(lang, dataType, subType string, ds *DataSet) {
	dm := getDataManager()
	key := fmt.Sprintf("%s:%s:%s", lang, dataType, subType)
	dm.registered[key] = ds
}

// LoadDataSet 加载指定的数据集
func (dm *DataManager) LoadDataSet(language Language, dataType, subType string) (*DataSet, error) {
	key := fmt.Sprintf("%s:%s:%s", language, dataType, subType)

	dm.mu.RLock()
	ds, ok := dm.registered[key]
	dm.mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("data set not found: %s", key)
	}

	return ds, nil
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
			weight = 1.0
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

// ListAvailableDataSets 列出可用的数据集
func (dm *DataManager) ListAvailableDataSets() []string {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	var dataSets []string
	for key := range dm.registered {
		dataSets = append(dataSets, key)
	}
	return dataSets
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
