# Go Runtime GMP 初始化流程图标映射

本文档说明 [`go_runtime_gmp_init.excalidraw`](../../../go_runtime_gmp_init.excalidraw) 中的实体应如何复用 `developer_library/golang/runtime` 图标。函数名和状态转换继续使用流程框与文字，不单独制作图标。

## 已有图标即可覆盖的流程实体

| 流程片段 | 使用图标 |
| --- | --- |
| G、M、P 的核心关系与 `execute(G)` | GMP：G、M、P、Goroutine |
| `p.runnext`、`p.runq`、`sched.runq`、work stealing | Scheduler：Run Queue、Global Queue；使用箭头表达转移与偷取。 |
| `netpoll`、定时唤醒、等待队列 | Scheduler：Netpoll、Timer、pollDesc、sudog。 |
| 栈检查、GC 辅助与回收 | Memory：Stack；GC：GC。 |

## 新增图标的使用位置

| 流程片段 | 新图标 | 使用方式 |
| --- | --- | --- |
| `_rt0_go` 初始化 `m0` / `g0` | m0 Bootstrap Thread、g0 System Goroutine | 在启动链路中分别替换“第一个 OS Thread”和“调度栈”概念框。 |
| `schedinit` 的全局 sched 初始化 | schedt Global Scheduler | 表示 `runtime.schedt`，不要与 `schedule()` 循环图标混用。 |
| 读取 `GOMAXPROCS`、创建 P 列表 | GOMAXPROCS | 放在配置输入到 `procresize` 的链路上。 |
| `_Grunning` | GMP：_Grunning | 绿色具名状态图标，表示 G 已绑定 M 与 P、正在执行 Go 代码。 |
| `_Grunnable`、`_Gwaiting`、`_Gsyscall`、`_Gdead`、`_Gpreempted` | GMP：对应具名 G 状态图标 | 使用与状态语义对应的蓝、黄、紫、灰、橙色状态卡；状态名称仍保留文字。 |
| `sysmon` 抢占、retake P、timer/GC/netpoll 辅助 | sysmon | 放在后台监控支路，箭头指向受影响的调度或唤醒路径。 |
| `entersyscall`、内核阻塞、`exitsyscall` | Syscall / Kernel Boundary | 放在 syscall 分支上，G/M/P 的解绑与重新获取 P 仍使用箭头和既有 GMP 图标。 |

## 不应新增为图标的项目

`rt0_go`、`runtime.args`、`runtime.osinit`、`runtime.schedinit`、`newproc`、`procresize`、`runtime.mstart`、`schedule`、`findRunnable`、`execute`、`goready`、`goexit` 与 work stealing 都是函数、操作或流程步骤，使用流程框、文字和箭头表达即可。
