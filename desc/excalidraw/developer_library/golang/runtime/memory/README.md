# Go Runtime Memory Excalidraw Library

该图标库用于绘制 Go Runtime 的栈、堆、缓存、span 与页分配层级。

## 文件

- `golang_runtime_memory_library.excalidrawlib`：可直接导入 Excalidraw 的内存图标库。
- `golang_runtime_memory_icons.excalidraw`：内存图标浏览图板。

## 图标

| 图标 | 填充色 | 用途说明 |
| --- | --- | --- |
| Heap | `#ffe8cc` | Go 堆内存和对象分配区域。 |
| Stack | `#ffd8a8` | goroutine 栈与栈增长。 |
| mcache | `#fff0f6` | P 本地小对象分配缓存。 |
| mcentral | `#f3d9fa` | 中心 span 缓存与补给层。 |
| mheap | `#e7f5ff` | 运行时堆管理和 span 分配来源。 |
| Arena | `#c3fae8` | heap arena 内存映射区域。 |
| Span | `#fff9db` | `mspan`，连续页组成的分配单元。 |
| Page | `#f1f3f5` | runtime 内存页或 page allocator 页。 |

在 Excalidraw 的 Library 面板中导入 `developer_library/golang/runtime/memory/golang_runtime_memory_library.excalidrawlib`。

## 图标规范

遵循 [Go Runtime 图标总设计规范](../ICON_DESIGN_SPEC.md)。

- Heap、Stack、mcache、mcentral、mheap、Arena、Page 使用约 `50 × 56 px` 的竖向分层矩形，标签为居中的 `11 px`；内部横线仅表示容器层级或页分隔。
- Span 使用约 `60 × 34 px` 的横向矩形与 `10 px` 标签，不添加分层线，突出连续分配单元的语义。
- 填充色固定为：Heap `#ffe8cc`、Stack `#ffd8a8`、mcache `#fff0f6`、mcentral `#f3d9fa`、mheap `#e7f5ff`、Arena `#c3fae8`、Span `#fff9db`、Page `#f1f3f5`。
- 内存层级与分配路径使用外部连线表达；图标内部不得加入对象数量、容量或版本等动态数据。
