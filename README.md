# gwp
go web programming


## Contents
```sh
|---| v1  # 单表的CRUD
    | v2  # 两张关联表的CRUD
```

## Concepts

关系数据库：基于数据的关系模型构建的数据库。可以在表与表之间建立关系，从而使数据能互相进行关联。

有4中方式可以把一项记录与其他记录关联起来：
- 一对一关联：has one (有一个的关系)
- 一对多关联：has many (有多个的关系)
- 多对一关联：belongs to (属于的关系)
- 多对多关联

ORM: 对象-关系映射器(Object-Relational Mapper)，可以将关系数据库中的表映射为编程语言中的对象。
> 如Java中的Hibernate, Ruby中的ActiveRecord

Golang中也有类似的`关系映射器`：`Sqlx`, `Gorm`

- Sqlx: 它为database/sql包提供了一系列非常有用的扩展功能，使用与database/sql相同的接口，但是它支持的特性不多。
- Gorm: 它提供了一个完整而强大的ORM机制来替代database/sql包。它遵循的是数据映射器模式(Data-Mapper pattern),该模式是通过提供映射器来将数据库中的数据映射为结构。

```
go get github.com/jmoiron/sqlx
go get github.com/jinzhu/gorm
```

