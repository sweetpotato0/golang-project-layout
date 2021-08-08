# golang项目结构

本项目主要参考了 DDD 以及 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

启动

使用 `wire` 生成项目依赖
```sh
cd cmd/article
wire
```

构建运行
```sh
go build -o ./bin/ ./...
./bin/article
```

configs中的 `toml` 文件进行 Mysql、Redis 配置，然后进行数据库创建：
```sql
create table articles( id int, title varchar(200), content varchar(200));
insert into articles(id, title, content) values(1, 'title', 'content');

create table user( id int, name varchar(200));
insert into user(id, name) values(1, 'jackiezhang');
```

浏览器输入 http://127.0.0.1:8080/articles 可以看到对应的效果

# TODO

- [x] 增加 grpc
- [x] 增加服务注册、服务发现
- [ ] http server 增加中间件
- [ ] config app 独立开
- [ ] 接入sentry
- [ ] 接入jaeger