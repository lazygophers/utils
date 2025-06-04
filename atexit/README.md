# atexit  
程序退出时的安全清理机制  

## 核心功能  
1. 注册退出时执行的清理函数（如资源释放、日志记录）  
2. 支持跨平台实现（当前包含Linux系统优化）  

## 使用示例  
```go  
import "github.com/lazygophers/utils/atexit"  

func main() {  
    atexit.Register(func() {  
        fmt.Println("程序退出时执行清理操作")  
        cleanupResources()  
    })  
}  
```  

## 注意事项  
- 注册函数会在程序正常退出（os.Exit）或panic时执行  
- 执行顺序与注册顺序相反（后注册的先执行）