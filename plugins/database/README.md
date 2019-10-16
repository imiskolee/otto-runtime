# Database


# Examples

```go
database.RegisterDatabaseResolver(func(conn string) *gorm.DB {
//do here........
})
```

```js
var connection = NewDatabaseConnection("<Your Database Name here>")
var lst = connection.Query("SELECT * FROM `<Your Table>`")
console.log(lst.length)
```