# golang项目结构

本项目主要参考了 DDD 以及 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

启动

使用 `wire` 生成项目依赖
> cd cmd/article
> wire

构建运行
> go build -o ./bin/ ./...
> ./bin/article

configs中的 `toml` 文件进行 Mysql、Redis 配置，然后进行数据库创建：
```sql
create table articles( id int, title varchar(200), content varchar(200));
insert into articles(id, title, content) values(1, 'title', 'content');
```

浏览器输入 http://127.0.0.1:8080/articles 可以看到对应的效果