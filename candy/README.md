# candy  

提供糖果语法解析与转换工具集  

**核心功能**  
- 支持Go语言中的糖果表达式解析  
- 转换逻辑：  
  ```go  
  func ParseCandy(expr string) (ast.Expr, error) {  
      // 解析糖果表达式  
  }  
  ```  
- 可扩展的转换规则引擎