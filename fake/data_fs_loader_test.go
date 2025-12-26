package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDataFSAndLoader 测试DataFS和DataLoader的功能
func TestDataFSAndLoader(t *testing.T) {
	// 测试DataFS的基本功能
	t.Run("test_data_fs", func(t *testing.T) {
		// 测试HasLanguage
		assert.True(t, dataFS.HasLanguage("en"))
		assert.False(t, dataFS.HasLanguage("invalid"))

		// 测试GetAvailableLanguages
		languages := dataFS.GetAvailableLanguages()
		assert.Greater(t, len(languages), 0)
	})

	// 测试DataLoader的功能
	t.Run("test_data_loader", func(t *testing.T) {
		manager := getDataManager()

		// 测试LoadDataSet
		dataset, err := manager.LoadDataSet(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.NotNil(t, dataset)

		// 测试GetItems
		items, err := manager.GetItems(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.Greater(t, len(items), 0)

		// 测试GetItemValues
		values, err := manager.GetItemValues(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.Greater(t, len(values), 0)

		// 测试GetWeightedItems
		weightedValues, weights, err := manager.GetWeightedItems(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.Greater(t, len(weightedValues), 0)
		assert.Equal(t, len(weightedValues), len(weights))

		// 测试GetCacheStats
		stats := manager.GetCacheStats()
		assert.Greater(t, stats["cached_datasets"], 0)

		// 测试ClearCache
		manager.ClearCache()
		stats = manager.GetCacheStats()
		assert.Equal(t, 0, stats["cached_datasets"])

		// 重新加载数据集以测试缓存
		_, err = manager.LoadDataSet(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		stats = manager.GetCacheStats()
		assert.Greater(t, stats["cached_datasets"], 0)

		// 测试ListAvailableDataSets
		datasets, err := manager.ListAvailableDataSets()
		assert.NoError(t, err)
		assert.Greater(t, len(datasets), 0)
	})

	// 测试快速访问函数
	t.Run("test_fast_access_functions", func(t *testing.T) {
		// 测试loadDataSet
		dataset, err := loadDataSet(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.NotNil(t, dataset)

		// 测试getItemValues
		values, err := getItemValues(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.Greater(t, len(values), 0)

		// 测试getWeightedItems
		weightedValues, weights, err := getWeightedItems(LanguageEnglish, "names", "first_male")
		assert.NoError(t, err)
		assert.Greater(t, len(weightedValues), 0)
		assert.Equal(t, len(weightedValues), len(weights))
	})

	// 测试PreloadData
	t.Run("test_preload_data", func(t *testing.T) {
		manager := getDataManager()
		manager.ClearCache()

		// 预加载英文数据
		err := manager.PreloadData(LanguageEnglish)
		assert.NoError(t, err)

		// 检查缓存是否已填充
		stats := manager.GetCacheStats()
		assert.Greater(t, stats["cached_datasets"], 0)
	})
}
