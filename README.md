# Thor

项目根目录运行如下命令启动项目

```
go run main.go  -a web -c ./conf/app.toml
```

## 相关文档

* [ldap用户管理](http://dev.flyaha.top:8004/cmd.php?server_id=1&redirect=true)
* [ldap用户创建wiki](http://wiki.flyaha.top/zh/%E6%8A%80%E6%9C%AF%E5%88%86%E4%BA%AB/%E6%9D%A8%E7%A5%96%E8%B1%AA/Other/LDAP%E6%90%AD%E5%BB%BA%E5%8F%8A%E7%94%A8%E6%88%B7%E5%88%9B%E5%BB%BA)

## 第一版部署

```yaml
---

kind: pipeline
type: ssh
name: thor_serv
server:
  host:
    from_secret: apps-host
  user:
    from_secret: apps-user
  password:
    from_secret: apps-pwd
steps:
  - name: thor go deploy
    commands:
      - chmod 755 ./build.sh
      - sh build.sh
  when:
    branch: master
    event: push
    status: success
  branches:
    include: [ master ]

```
