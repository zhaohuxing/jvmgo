## 类型转换文档

- 170608:简单的json数据转换, 代码jsonSwap.go


## json的简单的使用

json作为数据传输的标配,自然而然需要掌握如下两点:
- struct 实例转换为 json格式的数据
- json格式的数据解析到struct实例

如上这些操作,Go中`encoding/json`都搞定了!

#### 实现二者转换的基础:struct和json数据
定义struct时,字段的类型后面可以定义tag, 如下所示:
```
type Admin struct {
	Id       int //`json:"-"`
	Username string
	Password string
}
```
在json解析时是区分大小的,大写才会被解析,所以想导出导入就需要首字母大写!

Id后面被`json:"-"`标记后,就不会被解析到json数据中.使用`struct-tag`注意事项:
- 字段的tag是"-"，那么这个字段不会输出到JSON
- tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
- tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
- 如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

#### 将Struct实例转换为json格式数据
Go自带库`encoding/json`中的方法如下:
```
func Unmarshal(data []byte, v interface{}) error
```
使用上面定义的Admin Struct进行Unmarshal操作:
```
func (self *Admin) Otoj() {
	bytes, err := json.Marshal(self)
	if err != nil {
		fmt.Println("解析失败: struct to json")
		return
	}
	fmt.Println(string(bytes))
}
```

#### 将json数据解析到Struct实例
Go自带库`encoding/json`中的方法如下:
```
func Marshal(v interface{}) ([]byte, error)
```
使用上面定义的Admin Struct进行Marshal操作:
```
//Jtoo: json --- > struct

func (self *Admin) Jtoo(values string) {
	fmt.Println("将json数据解析到Admin前:")
	fmt.Println("ID:", self.Id)
	fmt.Println("Username:", self.Username)
	fmt.Println("Password:", self.Password)
	json.Unmarshal([]byte(values), self)
	fmt.Println("将json数据解析到Admin后:")
	fmt.Println("ID:", self.Id)
	fmt.Println("Username:", self.Username)
	fmt.Println("Password:", self.Password)
}
```

#### 测试程序:
```
func main() {
	admin := &Admin{}
	str :=
		`{
			"id": 1,
			"Username": "Root",
			"Password": "123456"
		 }`
	admin.Jtoo(str)
	admin.Otoj()
}
```
