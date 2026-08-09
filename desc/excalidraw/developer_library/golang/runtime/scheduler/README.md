# Go Runtime Scheduler Excalidraw Library

该图标库用于绘制 Go Runtime 调度、运行队列、网络轮询、定时器和等待节点关系。

## 文件

- `golang_runtime_scheduler_library.excalidrawlib`：可直接导入 Excalidraw 的调度图标库。
- `golang_runtime_scheduler_icons.excalidraw`：调度图标浏览图板。

## 图标

| 图标 | 填充色 | 用途说明 |
| --- | --- | --- |
| Scheduler | `#ebfbee` | GMP 调度器或调度循环。 |
| Run Queue | `#d0bfff` | P 的本地可运行 G 队列。 |
| Global Queue | `#ffdeeb` | 全局可运行 G 队列。 |
| Netpoll | `#a5d8ff` | 网络轮询器与异步 IO 唤醒。 |
| Timer | `#ffec99` | 定时器堆、计时任务和超时唤醒。 |
| pollDesc | `#ffc9c9` | 网络描述符、等待队列和 netpoll 状态。 |
| sudog | `#b2f2bb` | channel/同步等待队列中的 G 包装节点。 |
| schedt Global Scheduler | `#d8f5a2` | 保存全局 runq、空闲 P/M 等调度状态。 |
| sysmon | `#d0f4de` | 周期检查抢占、syscall、timer 与 netpoll 的系统监控器。 |
| GOMAXPROCS | `#fff4e6` | 决定可执行 Go 代码的 P 数量上限。 |
| Syscall / Kernel Boundary | `#e9ecef` | 表示 entersyscall、内核阻塞与 exitsyscall 边界。 |

在 Excalidraw 的 Library 面板中导入 `developer_library/golang/runtime/scheduler/golang_runtime_scheduler_library.excalidrawlib`。

## 图标规范

遵循 [Go Runtime 图标总设计规范](../ICON_DESIGN_SPEC.md)。

- Scheduler 使用约 `60 × 60 px` 的圆形和内部箭头，`10 px` 标签，填充 `#ebfbee`；箭头只表示调度循环。
- Run Queue、Global Queue 使用约 `60 × 34 px` 矩形和 `10 px` 标签，填充分别为 `#d0bfff`、`#ffdeeb`。
- Netpoll、sudog 使用约 `48 × 48 px` 圆形和 `13 px` 短标签，填充分别为 `#a5d8ff`、`#b2f2bb`；Timer 使用约 `48 × 61 px` 圆形加内部刻度线、`9 px` 标签、填充 `#ffec99`。
- pollDesc 使用约 `40 × 56 px` 分层矩形、`9 px` 标签和填充 `#ffc9c9`；内部线只表达描述符字段分区。
- schedt 使用约 `60 × 56 px` 的浅绿色全局调度面板，带全局 runq 槽位和资源线，区别于仅表示调度循环的 Scheduler。
- sysmon 使用约 `54 × 54 px` 的浅绿色监控环，带时钟指针与中心枢纽；只表示系统监控职责，不表示特定 goroutine。
- GOMAXPROCS 使用约 `64 × 48 px` 的浅橙配置卡，顶部为配置名、底部为三个 P 容量槽；槽位表达可用处理器数量，不表示实际固定数量。
- Syscall / Kernel Boundary 使用约 `60 × 56 px` 的灰色上下分区卡，箭头从 syscall 指向 kernel，表示运行时与内核的系统调用边界。
- 队列连接、netpoll 唤醒、定时器到期和 park/unpark 关系均用使用图的外部箭头表达，不在节点图标内重复绘制。
