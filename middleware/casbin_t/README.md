# Casbin + JWT + Fiber 示例

本目录是一个独立入口，复用仓库中的 `common/module/gorm/database` 和
`common/module/cache`。先按需要修改 `common/config/config.yml` 的 MySQL、Redis
连接信息。目标数据库必须已存在；首次启动会由 GORM 创建三张表，再导入初始化
数据：

```bash
CASBIN_T_JWT_SECRET='replace-with-a-long-random-secret' go run ./middleware/casbin_t
mysql -u root -p go_test < middleware/casbin_t/db.sql
```

默认监听 `:3001`（可用 `CASBIN_T_ADDR` 覆盖）。三个演示账号的密码都是
`Admin@2026!`：`admin`、`operate`、`user`。
部署时也可以用 `GOTEST_CONFIG_PATH=/path/to/config.yml` 指向独立的数据库与
Redis 配置文件。

```bash
# 登录，取响应中的 access_token
curl -X POST http://127.0.0.1:3001/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"user","password":"Admin@2026!"}'

curl http://127.0.0.1:3001/api/tests -H "Authorization: Bearer $TOKEN"
curl -X POST http://127.0.0.1:3001/api/tests -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"name":"first","content":"demo"}'
curl -X PUT http://127.0.0.1:3001/api/tests/1 -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"name":"updated","content":"demo"}'
curl -X DELETE http://127.0.0.1:3001/api/tests/1 -H "Authorization: Bearer $TOKEN"
```

权限由 `casbin_rule` 保存：`admin` 可增删查改，`operate` 可增查改，`user`
只能查询。登录时角色会以 `casbin_t:role:<username>` 缓存在 Redis 5 分钟；Redis
不可用时会回退到 MySQL，不影响登录。
