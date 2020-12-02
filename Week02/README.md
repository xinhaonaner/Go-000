学习笔记

#### 题目

> 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
>

思路：
dao层只是用于获取数据，对于数据的处理应该放在业务层去做，所以在dao层应该如实反映操作数据的情况。 因此，当dao层遇到一个sql.ErrNoRows时，应该wrap起来抛给上层。但是为了业务层兼容多种数据库存储，我们一般不直接对sql.ErrNoRows进行wrap，而是wrap一个由业务指定的NOT FOUND的错误给上层，对于 dao 层，不同 DB 的 lib 库可能有区别

```
// Go 避免野协程 panic
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover:", err)
			}
		}()
		x()
	}()
}
```


