# superstar-db
简单的球星库

依赖的包有：
go-sql-driver/mysql
go-xorm/xorm
gorilla/securecookie
kataras/iris

注意设置go build main.go中的 Edit Configuration中的WorkDictionary
让其指向WEB目录下，从而可以使用简洁路径方式访问到前端静态文件