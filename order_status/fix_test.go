package order_status

import (
	"fmt"
	"testing"
	"time"
)

// 测试修复方案是否能解决阶段流转问题
func TestFixStageFlow(t *testing.T) {
	// 创建测试订单
	orderID := "test-order-fix"
	CreateOrder(orderID)
	
	// 模拟问题场景：跳过初始化，直接完成访问页面
	fmt.Println("\n[测试] 模拟问题场景：跳过初始化，直接完成访问页面")
	UpdateStageStatus(orderID, StageVisitPage, StatusCompleted, 100)
	
	// 等待修复完成
	time.Sleep(500 * time.Millisecond)
	
	// 验证修复结果
	order, _ := GetOrder(orderID)
	if order.Stages[StageInitialization].Status != StatusCompleted {
		t.Errorf("初始化阶段应该被修复为已完成，但实际状态是 %s", order.Stages[StageInitialization].Status)
	}
	
	if order.Stages[StageVisitPage].Status != StatusCompleted {
		t.Errorf("访问页面阶段应该保持已完成，但实际状态是 %s", order.Stages[StageVisitPage].Status)
	}
	
	fmt.Println("\n[测试] 修复方案验证通过！")
}

// 测试多阶段流转修复
func TestMultiStageFlowFix(t *testing.T) {
	// 创建测试订单
	orderID := "test-order-multi-fix"
	CreateOrder(orderID)
	
	// 模拟跳过前面多个阶段，直接完成等待渲染
	fmt.Println("\n[测试] 模拟问题场景：跳过初始化和访问页面，直接完成等待渲染")
	UpdateStageStatus(orderID, StageWaitingRender, StatusCompleted, 100)
	
	// 等待修复完成
	time.Sleep(500 * time.Millisecond)
	
	// 验证修复结果
	order, _ := GetOrder(orderID)
	if order.Stages[StageInitialization].Status != StatusCompleted {
		t.Errorf("初始化阶段应该被修复为已完成，但实际状态是 %s", order.Stages[StageInitialization].Status)
	}
	
	if order.Stages[StageVisitPage].Status != StatusCompleted {
		t.Errorf("访问页面阶段应该被修复为已完成，但实际状态是 %s", order.Stages[StageVisitPage].Status)
	}
	
	if order.Stages[StageWaitingRender].Status != StatusCompleted {
		t.Errorf("等待渲染阶段应该保持已完成，但实际状态是 %s", order.Stages[StageWaitingRender].Status)
	}
	
	fmt.Println("\n[测试] 多阶段修复验证通过！")
}

// 测试手动调用修复函数
func TestManualFix(t *testing.T) {
	// 创建测试订单
	orderID := "test-order-manual-fix"
	CreateOrder(orderID)
	
	// 模拟问题场景
	UpdateStageStatus(orderID, StageVisitPage, StatusCompleted, 100)
	
	// 手动调用修复函数
	fmt.Println("\n[测试] 手动调用修复函数")
	FixOrderStageFlow(orderID)
	
	// 等待修复完成
	time.Sleep(500 * time.Millisecond)
	
	// 验证修复结果
	order, _ := GetOrder(orderID)
	if order.Stages[StageInitialization].Status != StatusCompleted {
		t.Errorf("初始化阶段应该被修复为已完成，但实际状态是 %s", order.Stages[StageInitialization].Status)
	}
	
	fmt.Println("\n[测试] 手动修复验证通过！")
}


