package order_status

import (
	"sync"

	"github.com/lazygophers/utils/event"
)

// 修复订单阶段流转问题的解决方案

// 阶段流转修复器，确保订单阶段按照正确的顺序完成
type StageFlowFixer struct {
	orderMutex sync.RWMutex
}

var defaultFixer = &StageFlowFixer{}

// 修复订单阶段流转，确保前面的阶段在后面的阶段完成时也被标记为完成
func (fixer *StageFlowFixer) FixStageFlow(order *Order) {
	fixer.orderMutex.Lock()
	defer fixer.orderMutex.Unlock()
	
	// 阶段顺序：初始化 → 访问页面 → 等待渲染 → 截取图片 → 保存文件
	stageOrder := []OrderStage{
		StageInitialization,
		StageVisitPage,
		StageWaitingRender,
		StageScreenshot,
		StageSaveFile,
	}
	
	// 找到最后一个已完成的阶段
	lastCompletedIndex := -1
	for i, stage := range stageOrder {
		if order.Stages[stage].Status == StatusCompleted {
			lastCompletedIndex = i
		} else {
			// 一旦遇到未完成的阶段，就停止查找
			break
		}
	}
	
	// 如果有已完成的阶段，确保前面的所有阶段也都已完成
	if lastCompletedIndex > 0 {
		for i := 0; i <= lastCompletedIndex; i++ {
			stage := stageOrder[i]
			if order.Stages[stage].Status != StatusCompleted {
				order.Stages[stage].Status = StatusCompleted
				order.Stages[stage].Progress = 100
				
				// 触发阶段状态更新事件，通知前端更新UI
			event.Emit("order.stage.updated", map[string]interface{}{
					"order_id": order.ID,
					"stage":    stage,
					"status":   StatusCompleted,
					"progress": 100,
				})
			}
		}
	}
}

// 修复特定订单的阶段流转
func FixOrderStageFlow(orderID string) {
	order, exists := GetOrder(orderID)
	if exists {
		defaultFixer.FixStageFlow(order)
	}
}

// 监听订单阶段更新事件，自动修复阶段流转
func init() {
	// 注册事件监听器，当任何订单阶段更新时自动修复
	event.Register("order.stage.updated", func(args any) {
		data := args.(map[string]interface{})
		orderID := data["order_id"].(string)
		FixOrderStageFlow(orderID)
	})
}
