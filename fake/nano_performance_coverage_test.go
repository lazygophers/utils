package fake

import (
	"testing"
)

// TestNanoPerformanceFunctions 测试纳秒级性能优化函数
func TestNanoPerformanceFunctions(t *testing.T) {
	// 测试NewNanoPerformance和NanoName方法
	nano := NewNanoPerformance()
	name := nano.NanoName()
	if name == "" {
		t.Error("NanoName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		name := nano.NanoName()
		if name == "" {
			t.Errorf("NanoName() should not return empty string on call %d", i)
		}
	}

	// 测试NewAtomic和AtomicName方法
	atomic := NewAtomic()
	atomicName := atomic.AtomicName()
	if atomicName == "" {
		t.Error("AtomicName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		atomicName := atomic.AtomicName()
		if atomicName == "" {
			t.Errorf("AtomicName() should not return empty string on call %d", i)
		}
	}

	// 测试NewConstant和ConstantName方法
	constant := NewConstant()
	constantName := constant.ConstantName()
	if constantName == "" {
		t.Error("ConstantName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		constantName := constant.ConstantName()
		if constantName == "" {
			t.Errorf("ConstantName() should not return empty string on call %d", i)
		}
	}

	// 测试NewIncrementOnly和IncrementOnlyName方法
	incrementOnly := NewIncrementOnly()
	incrementName := incrementOnly.IncrementOnlyName()
	if incrementName == "" {
		t.Error("IncrementOnlyName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		incrementName := incrementOnly.IncrementOnlyName()
		if incrementName == "" {
			t.Errorf("IncrementOnlyName() should not return empty string on call %d", i)
		}
	}

	// 测试NewStatic和StaticName方法
	static := NewStatic()
	staticName := static.StaticName()
	if staticName == "" {
		t.Error("StaticName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		staticName := static.StaticName()
		if staticName == "" {
			t.Errorf("StaticName() should not return empty string on call %d", i)
		}
	}

	// 测试NewCPUOptimized和CPUOptimizedName方法
	cpu := NewCPUOptimized()
	cpuName := cpu.CPUOptimizedName()
	if cpuName == "" {
		t.Error("CPUOptimizedName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		cpuName := cpu.CPUOptimizedName()
		if cpuName == "" {
			t.Errorf("CPUOptimizedName() should not return empty string on call %d", i)
		}
	}

	// 测试NewBranchless和BranchlessName方法
	branchless := NewBranchless()
	branchlessName := branchless.BranchlessName()
	if branchlessName == "" {
		t.Error("BranchlessName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		branchlessName := branchless.BranchlessName()
		if branchlessName == "" {
			t.Errorf("BranchlessName() should not return empty string on call %d", i)
		}
	}
}

// TestGlobalNanoFunctions 测试全局纳秒级函数
func TestGlobalNanoFunctions(t *testing.T) {
	// 测试NanoName全局函数
	name := NanoName()
	if name == "" {
		t.Error("Global NanoName() should not return empty string")
	}

	// 测试多次调用
	for i := 0; i < 10; i++ {
		name := NanoName()
		if name == "" {
			t.Errorf("Global NanoName() should not return empty string on call %d", i)
		}
	}

	// 测试AtomicName全局函数
	atomicName := AtomicName()
	if atomicName == "" {
		t.Error("Global AtomicName() should not return empty string")
	}

	// 测试ConstantName全局函数
	constantName := ConstantName()
	if constantName == "" {
		t.Error("Global ConstantName() should not return empty string")
	}

	// 测试IncrementOnlyName全局函数
	incrementName := IncrementOnlyName()
	if incrementName == "" {
		t.Error("Global IncrementOnlyName() should not return empty string")
	}

	// 测试StaticName全局函数
	staticName := StaticName()
	if staticName == "" {
		t.Error("Global StaticName() should not return empty string")
	}

	// 测试CPUOptimizedName全局函数
	cpuName := CPUOptimizedName()
	if cpuName == "" {
		t.Error("Global CPUOptimizedName() should not return empty string")
	}

	// 测试BranchlessName全局函数
	branchlessName := BranchlessName()
	if branchlessName == "" {
		t.Error("Global BranchlessName() should not return empty string")
	}

	// 测试UltimatePerformanceName全局函数
	ultimateName := UltimatePerformanceName()
	if ultimateName == "" {
		t.Error("Global UltimatePerformanceName() should not return empty string")
	}

	// 测试FastAtomicAdd函数
	atomicAdd := FastAtomicAdd()
	if atomicAdd == 0 {
		t.Error("FastAtomicAdd() should not return 0")
	}

	// 测试FastAtomicLoad函数
	atomicLoad := FastAtomicLoad()
	if atomicLoad == 0 {
		t.Error("FastAtomicLoad() should not return 0")
	}

	// 测试FastAtomicCAS函数
	atomicCAS := FastAtomicCAS()
	if !atomicCAS {
		t.Error("FastAtomicCAS() should return true")
	}
}
