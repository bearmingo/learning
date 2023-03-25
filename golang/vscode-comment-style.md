# 代码注释

换行需要添加空格缩紧

函数或者变量使用 \`\`包起来

```go
// removeMember 删除活动成员。
//   不支持删除活动的owner，owner需要调用 `LeaveActivity`
func removeMember(tx gorm.DB, userIDs []uint64, activityID uint64) error {
...
}
```