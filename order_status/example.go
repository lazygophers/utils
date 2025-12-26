package order_status

import (
	"fmt"
	"sync"
	"time"

	"github.com/lazygophers/utils/event"
	"github.com/lazygophers/utils/routine"
)

// 订单阶段类型
type OrderStage string

const (
	StageInitialization OrderStage = "初始化"
	StageVisitPage      OrderStage = "访问页面"
	StageWaitingRender  OrderStage = "等待渲染"
	StageScreenshot     OrderStage = "截取图片"
	StageSaveFile       OrderStage = "保存文件"
)

// 阶段状态
type StageStatus string

const (
	StatusPending   StageStatus = "pending"   // 等待中
	StatusInProgress StageStatus = "in_progress" // 进行中
	StatusCompleted StageStatus = "completed"  // 已完成
	StatusFailed    StageStatus = "failed"     // 失败
)

// 订单阶段信息
type OrderStageInfo struct {
	Stage  OrderStage  `json:"stage"`
	Status StageStatus `json:"status"`
	Progress int       `json:"progress"` // 0-100
}

// 订单信息
type Order struct {
	ID       string                      `json:"id"`
	Stages   map[OrderStage]*OrderStageInfo `json:"stages"`
	Status   string                      `json:"status"`
	TotalProgress int                    `json:"total_progress"`
	mutex    sync.RWMutex
}

// 订单管理器
type OrderManager struct {
	orders map[string]*Order
	mutex  sync.RWMutex
}

var defaultManager = NewOrderManager()

// 创建新的订单管理器
func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders: make(map[string]*Order),
	}
}

// 创建新订单
func (om *OrderManager) CreateOrder(orderID string) *Order {
	om.mutex.Lock()
	defer om.mutex.Unlock()
	
	order := &Order{
		ID:     orderID,
		Status: "downloading",
		Stages: map[OrderStage]*OrderStageInfo{
			StageInitialization: {Stage: StageInitialization, Status: StatusPending, Progress: 0},
			StageVisitPage:      {Stage: StageVisitPage, Status: StatusPending, Progress: 0},
			StageWaitingRender:  {Stage: StageWaitingRender, Status: StatusPending, Progress: 0},
			StageScreenshot:     {Stage: StageScreenshot, Status: StatusPending, Progress: 0},
			StageSaveFile:       {Stage: StageSaveFile, Status: StatusPending, Progress: 0},
		},
		TotalProgress: 0,
	}
	
	om.orders[orderID] = order
	return order
}

// 获取订单
func (om *OrderManager) GetOrder(orderID string) (*Order, bool) {
	om.mutex.RLock()
	defer om.mutex.RUnlock()
	
	order, exists := om.orders[orderID]
	return order, exists
}

// 更新订单阶段状态
func (om *OrderManager) UpdateStageStatus(orderID string, stage OrderStage, status StageStatus, progress int) {
	order, exists := om.GetOrder(orderID)
	if !exists {
		return
	}
	
	order.mutex.Lock()
	defer order.mutex.Unlock()
	
	stageInfo, exists := order.Stages[stage]
	if !exists {
		return
	}
	
	// 更新阶段状态
	stageInfo.Status = status
	stageInfo.Progress = progress
	
	// 确保状态流转的正确性
	om.ensureStageFlow(order, stage, status)
	
	// 更新总进度
	om.updateTotalProgress(order)
	
	// 触发阶段状态更新事件
	event.Emit("order.stage.updated", map[string]interface{}{
		"order_id": orderID,
		"stage":    stage,
		"status":   status,
		"progress": progress,
	})
	
	// 检查订单是否完成
	om.checkOrderCompletion(order)
}

// 确保阶段流转的正确性
func (om *OrderManager) ensureStageFlow(order *Order, currentStage OrderStage, currentStatus StageStatus) {
	// 阶段顺序：初始化 → 访问页面 → 等待渲染 → 截取图片 → 保存文件
	stageOrder := []OrderStage{
		StageInitialization,
		StageVisitPage,
		StageWaitingRender,
		StageScreenshot,
		StageSaveFile,
	}
	
	// 如果当前阶段已完成，确保前面的所有阶段也标记为已完成
	if currentStatus == StatusCompleted {
		currentIndex := -1
		for i, stage := range stageOrder {
			if stage == currentStage {
				currentIndex = i
				break
			}
		}
		
		// 确保前面的所有阶段都已完成
		for i := 0; i <= currentIndex; i++ {
			stage := stageOrder[i]
			if order.Stages[stage].Status != StatusCompleted {
				order.Stages[stage].Status = StatusCompleted
				order.Stages[stage].Progress = 100
			}
		}
	}
}

// 更新订单总进度
func (om *OrderManager) updateTotalProgress(order *Order) {
	total := 0
	completed := 0
	
	for _, stage := range order.Stages {
		total++
		if stage.Status == StatusCompleted {
			completed++
		}
	}
	
	// 计算总进度：已完成阶段数 / 总阶段数 * 100
	order.TotalProgress = (completed * 100) / total
}

// 检查订单是否完成
func (om *OrderManager) checkOrderCompletion(order *Order) {
	allCompleted := true
	for _, stage := range order.Stages {
		if stage.Status != StatusCompleted {
			allCompleted = false
			break
		}
	}
	
	if allCompleted && order.Status != "completed" {
		order.Status = "completed"
		event.Emit("order.completed", map[string]interface{}{
			"order_id": order.ID,
		})
	}
}

// 示例：使用订单管理器
func Example() {
	// 初始化订单
	orderID := "05736e4541ae49ffa8b33204db8d055a"
	defaultManager.CreateOrder(orderID)
	
	// 注册订单事件处理器
	event.Register("order.stage.updated", func(args any) {
		data := args.(map[string]interface{})
		fmt.Printf("订单 %s 的阶段 %s 状态更新为 %s，进度 %d%%\n",
			data["order_id"].(string),
			data["stage"].(OrderStage),
			data["status"].(StageStatus),
			data["progress"].(int))
	})
	
	event.Register("order.completed", func(args any) {
		data := args.(map[string]interface{})
		fmt.Printf("订单 %s 已完成！\n", data["order_id"].(string))
	})
	
	// 模拟订单处理流程
	routine.Go(func() error {
		// 1. 开始初始化
		defaultManager.UpdateStageStatus(orderID, StageInitialization, StatusInProgress, 50)
		time.Sleep(1 * time.Second)
		// 初始化完成
		defaultManager.UpdateStageStatus(orderID, StageInitialization, StatusCompleted, 100)
		
		// 2. 开始访问页面
		defaultManager.UpdateStageStatus(orderID, StageVisitPage, StatusInProgress, 50)
		time.Sleep(1 * time.Second)
		// 访问页面完成
		defaultManager.UpdateStageStatus(orderID, StageVisitPage, StatusCompleted, 100)
		
		// 3. 开始等待渲染
		defaultManager.UpdateStageStatus(orderID, StageWaitingRender, StatusInProgress, 74)
		time.Sleep(2 * time.Second)
		// 等待渲染完成
		defaultManager.UpdateStageStatus(orderID, StageWaitingRender, StatusCompleted, 100)
		
		// 4. 开始截取图片
		defaultManager.UpdateStageStatus(orderID, StageScreenshot, StatusInProgress, 50)
		time.Sleep(1 * time.Second)
		// 截取图片完成
		defaultManager.UpdateStageStatus(orderID, StageScreenshot, StatusCompleted, 100)
		
		// 5. 开始保存文件
		defaultManager.UpdateStageStatus(orderID, StageSaveFile, StatusInProgress, 50)
		time.Sleep(1 * time.Second)
		// 保存文件完成
		defaultManager.UpdateStageStatus(orderID, StageSaveFile, StatusCompleted, 100)
		
		return nil
	})
	
	// 模拟问题场景：访问页面已完成，但初始化未标记为完成
	routine.Go(func() error {
		orderID2 := "test-order-2"
		defaultManager.CreateOrder(orderID2)
		
		// 跳过初始化，直接完成访问页面
		defaultManager.UpdateStageStatus(orderID2, StageVisitPage, StatusCompleted, 100)
		
		return nil
	})
	
	// 等待一段时间，观察事件输出
	time.Sleep(10 * time.Second)
}

// 创建订单
func CreateOrder(orderID string) *Order {
	return defaultManager.CreateOrder(orderID)
}

// 获取订单
func GetOrder(orderID string) (*Order, bool) {
	return defaultManager.GetOrder(orderID)
}

// 更新订单阶段状态
func UpdateStageStatus(orderID string, stage OrderStage, status StageStatus, progress int) {
	defaultManager.UpdateStageStatus(orderID, stage, status, progress)
}
