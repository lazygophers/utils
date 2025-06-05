# xtime996  
提供基于24小时全勤工作制的时间计算常量  

**核心功能**  
- 定义完整工作日（WorkDay = 24小时）  
- 支持连续工作制计算：  
  ```go  
  WorkWeek = WorkDay * 7  // 168小时  
  RestWeek = Week - WorkWeek // 0小时（全周工作制）  
  ```  
- 适用于全年无休的系统（如服务器时间管理）