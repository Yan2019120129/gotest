# PostgreSQL Physical Storage Excalidraw Library

这套图标库用于绘制 PostgreSQL 物理存储结构图，重点覆盖 `Table -> Relation File -> Segment File -> Page / Block -> Tuple / Row -> Column Value`，也覆盖 MVCC、TOAST、Index、VACUUM、FSM/VM、OS Storage 等教学和架构分析场景。

## 文件

- `postgresql_storage_library.excalidrawlib`: 可直接导入 Excalidraw 的 Library 文件，每个图标是独立 item，并包含 `name + elements`。
- `postgresql_storage_icons.excalidraw`: 平铺浏览图板，按分类分区并标注图标名称。
- `icons_index.md`: 图标名称、中文解释、使用场景、推荐颜色和连接关系索引。
- `README.md`: 当前说明文档。

## 如何导入 Excalidraw

1. 打开 Excalidraw。
2. 打开 Library 面板。
3. 选择导入库文件。
4. 选择 `developer_library/postgresql_storage/postgresql_storage_library.excalidrawlib`。

## 使用规范

- 图标默认画布约 80x80，建议图标间距 32-48 px。
- 主链路使用实线箭头：`Table -> Relation File -> Segment File -> Page / Block -> Heap Tuple -> Column Value`。
- 引用关系、可见性辅助关系、索引回表关系建议使用虚线箭头。
- 不使用真实 PostgreSQL Logo，保持教材级、工程化、可复用表达。
- 颜色约定：蓝色逻辑层，橙色物理文件层，紫色 Page/Block，绿色 Tuple/Row，黄色 Column/Data，红色 MVCC/Vacuum，灰色 OS/Storage。

## PostgreSQL Storage 推荐布局

```text
Table
  |
Relation File
  |
Segment File
  |
Page / Block
  |
Heap Tuple
  |
Column Value
```

## 扩展方式

- 新增图标时，在 `scripts/generate_postgresql_storage_library.go` 的 `pgIcons()` 中增加一项。
- 复用现有 `Template` 可以保持风格统一；需要新结构时新增一个模板函数。
- 修改颜色请优先调整 `pgIcons()` 中的分类色，避免单个图标漂移。
- 重新生成时脚本会把已存在目标文件先备份为 `.bak`，再写入新文件。

## 图标数量

共 46 个图标。
