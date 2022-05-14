## Mode声明


## 方法
`First`调用时，自动会添加`order by`，可以替换成`Find`或者`Scan`避免这个情况。

## 外键关联

#### belong to
通过`foreignKey`中指定的字段，查找目标对象。
`foreignKey`: 当前结构体中指定的字段，一般这个字段的值是目标字段的主键。

#### has one
`foreignKey`: 目标结构体中的字段，一般这个字段的值是当前结构体的主键。

#### has many
`foreignKey`: 目标结构体中的字段。
`references`: 当前结构体中某一个字段。如果`foreignKey`指定的字段不是当前结构体的主键时，需要用`references`明确指定。

#### Many To Many
`many2many`:  指定关联表的名称
`foreignKey`:
`joinForeignKey`:
`References`:
`joinReferences`:

## 常用的方法

### 判断是否存在
```go
var exists bool
err = db.Model(model).Select("count(*) > 0").
		Limit(1).
		Where("id = ?", id).
        Find(&exists).
        Error
```