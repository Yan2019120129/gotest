# 大文件日志生成与读取示例

该示例提供日志生成和日志读取两个子命令。生成的数据保存在 `practice/file_t/read_t` 目录下的 `logs<大小>` 文件夹中，例如默认的 `logs20`。

以下命令均从 `gotest` 模块根目录执行。

## 生成日志

默认生成 20 GiB 日志，并写入 `practice/file_t/logs20`：

```bash
go run ./practice/file_t/read_t generate
```

生成较小的示例日志：

```bash
go run ./practice/file_t/read_t generate -size-gb 0.5
```

已存在同名日志时默认拒绝覆盖；确认需要重建时使用：

```bash
go run ./practice/file_t/read_t generate -size-gb 0.5 -overwrite
```

## 读取日志

不带子命令时保持读取兼容行为，也可以显式使用 `read`。三种方法使用相同的参数并返回指定时间范围内出现次数最多的节点。

- `custom`：单个读取协程逐行读取，多个 worker 并发解析。
- `custom_optimized`：将文件按字节分成 5 个对齐到换行符的分片，多个读取协程流式读取，再交给 worker 池处理。
- `traverse`：单协程顺序遍历读取。

```bash
go run ./practice/file_t/read_t read \
  -method custom_optimized \
  -file practice/file_t/logs20/access.log \
  -start 2026-07-29T13:32:00Z \
  -end 2026-07-29T13:55:00Z \
  -k 10
```

默认读取路径为 `practice/file_t/logs20/access.log`。省略 `read` 子命令时，以下命令等价：

```bash
go run ./practice/file_t/read_t -method traverse -file practice/file_t/logs20/access.log
```

## 参数

| 子命令 | 参数 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `generate` | `-size-gb` | `20` | 日志总大小（GiB） |
| `generate` | `-output-dir` | `practice/file_t/logs<大小>` | 日志输出目录 |
| `generate` | `-overwrite` | `false` | 允许覆盖已有日志 |
| `read` | `-method` | `custom_optimized` | `custom`、`custom_optimized` 或 `traverse` |
| `read` | `-file` | `practice/file_t/logs20/access.log` | 访问日志文件路径 |
| `read` | `-start`、`-end` | 当前年月 29 日的默认时间范围 | RFC3339 时间范围 |
| `read` | `-k` | `10` | 返回的节点数量 |
