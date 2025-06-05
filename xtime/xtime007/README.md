# xtime007  
提供基于12小时工作制的时间计算常量  

**核心功能**  
- 定义12小时工作日（WorkDay = 12小时）  
- 支持时间单位换算：  
  ```go  
  WorkWeek = WorkDay * 5  // 60小时  
  RestWeek = Week - WorkWeek // 96小时  
  ```  
- 适用于需要紧凑排班的场景（如轮班制系统）
