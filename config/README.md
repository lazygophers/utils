# config  

配置管理模块，支持多格式配置加载与验证  

**核心功能**  
- 支持JSON/TOML/YAML格式解析  
- 配置加载流程：  
  1. 优先使用指定路径  
  2. 坠回环境变量LAZYGOPHERS_CONFIG  
  3. 自动搜索当前目录及执行目录  
- 内置验证机制：  
  ```go  
  func LoadConfig(c any) error {  
      // 加载并验证配置  
  }