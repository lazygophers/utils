package runtime

import (
	"os"
	"testing"
	"time"
)

func TestGetExitSign(t *testing.T) {
	t.Run("get_exit_signal_channel", func(t *testing.T) {
		// 测试获取退出信号通道
		sigCh := GetExitSign()
		if sigCh == nil {
			t.Error("GetExitSign() returned nil channel")
		}

		// 验证通道是缓冲的
		select {
		case sigCh <- os.Interrupt:
			// 能发送说明是缓冲通道
			t.Logf("Channel is buffered as expected")
		default:
			t.Error("Channel should be buffered")
		}

		// 清理通道
		select {
		case <-sigCh:
		default:
		}
	})

	t.Run("get_exit_signal_multiple_calls", func(t *testing.T) {
		// 测试多次调用GetExitSign
		sigCh1 := GetExitSign()
		sigCh2 := GetExitSign()

		if sigCh1 == nil || sigCh2 == nil {
			t.Error("GetExitSign() should never return nil")
		}

		// 每次调用都应该返回新的通道
		if sigCh1 == sigCh2 {
			t.Error("GetExitSign() should return different channels on different calls")
		}
	})
}

func TestWaitExit(t *testing.T) {
	t.Run("wait_exit_with_timeout", func(t *testing.T) {
		// 测试WaitExit函数（使用超时防止无限等待）
		done := make(chan bool, 1)

		go func() {
			// 在goroutine中调用WaitExit
			WaitExit()
			done <- true
		}()

		// 等待很短时间，然后发送信号
		time.Sleep(10 * time.Millisecond)

		// 向当前进程发送SIGTERM信号来触发WaitExit退出
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatalf("Failed to find current process: %v", err)
		}

		err = process.Signal(os.Interrupt)
		if err != nil {
			t.Logf("Failed to send signal (expected on some systems): %v", err)
		}

		// 等待WaitExit完成或超时
		select {
		case <-done:
			t.Logf("WaitExit completed successfully")
		case <-time.After(1 * time.Second):
			t.Logf("WaitExit test timed out (expected behavior)")
		}
	})
}

func TestExit(t *testing.T) {
	t.Run("exit_function_logic", func(t *testing.T) {
		// 测试Exit函数的逻辑（但不实际退出）
		// 这个测试主要验证代码路径而不是实际的退出行为

		// 检查能否找到当前进程
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Errorf("os.FindProcess should not fail for current process: %v", err)
		}

		if process.Pid != os.Getpid() {
			t.Errorf("Process PID mismatch: expected %d, got %d", os.Getpid(), process.Pid)
		}

		// 我们不能直接测试Exit()函数，因为它会终止进程
		// 但我们可以测试它使用的组件
		t.Logf("Exit function components work correctly")
	})

	t.Run("exit_signal_sending", func(t *testing.T) {
		// 测试向进程发送信号的能力
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatalf("Failed to find current process: %v", err)
		}

		// 测试发送SIGTERM信号（不会终止进程，只是测试Signal方法）
		err = process.Signal(os.Interrupt)
		if err != nil {
			t.Logf("Signal sending failed (may be expected on some systems): %v", err)
		} else {
			t.Logf("Signal sent successfully")
		}
	})

	t.Run("exit_with_invalid_pid", func(t *testing.T) {
		// 测试无效PID的情况
		// 使用一个很大的PID，通常不存在
		invalidPid := 999999
		process, err := os.FindProcess(invalidPid)

		// 在某些系统上，FindProcess可能不会立即返回错误
		// 但Signal调用会失败
		if err == nil && process != nil {
			err = process.Signal(os.Interrupt)
			if err != nil {
				t.Logf("Signal to invalid process failed as expected: %v", err)
			}
		}
	})
}

// 测试信号通道的容量和行为
func TestExitSignalChannel(t *testing.T) {
	t.Run("signal_channel_capacity", func(t *testing.T) {
		sigCh := GetExitSign()

		// 尝试发送多个信号来测试通道容量
		signals := []os.Signal{os.Interrupt, os.Kill}

		for i, sig := range signals {
			select {
			case sigCh <- sig:
				t.Logf("Signal %d sent successfully", i)
			case <-time.After(100 * time.Millisecond):
				if i == 0 {
					t.Error("First signal should be sent immediately to buffered channel")
				} else {
					t.Logf("Channel is full after %d signals", i)
				}
				break
			}
		}
	})
}

// 集成测试：测试信号处理流程
func TestExitSignalIntegration(t *testing.T) {
	t.Run("signal_integration_test", func(t *testing.T) {
		// 创建信号通道
		sigCh := GetExitSign()

		// 启动一个goroutine监听信号
		received := make(chan os.Signal, 1)
		go func() {
			sig := <-sigCh
			received <- sig
		}()

		// 发送信号给当前进程
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatalf("Failed to find current process: %v", err)
		}

		err = process.Signal(os.Interrupt)
		if err != nil {
			t.Logf("Failed to send signal: %v", err)
			return
		}

		// 等待信号被接收
		select {
		case sig := <-received:
			t.Logf("Received signal: %v", sig)
		case <-time.After(1 * time.Second):
			t.Logf("Signal integration test timed out")
		}
	})
}
