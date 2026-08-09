# Go Runtime GMP Excalidraw Library

该图标库用于绘制 Go Runtime 的 Goroutine、G、M、P 及其基础执行关系。图标通过任务、上下文、线程与处理器资源的结构差异表达 GMP 的职责边界。

## 文件

- `golang_runtime_gmp_library.excalidrawlib`：可直接导入 Excalidraw 的 GMP 图标库。
- `golang_runtime_gmp_icons.excalidraw`：GMP 图标浏览图板。

## 图标

| 图标 | 填充色 | 用途说明 |
| --- | --- | --- |
| Goroutine | `#e3fafc` | 由 `go` 语句创建的轻量执行单元。 |
| G | `#d0ebff` | 保存状态、上下文与 goroutine 栈的 `runtime.g`。 |
| M | `#e5dbff` | 绑定 P 后承载 G 的操作系统线程抽象。 |
| P | `#fff3bf` | 持有本地 runq 等调度资源的逻辑处理器。 |
| g0 System Goroutine | `#dbe4ff` | 每个 M 的调度与系统栈。 |
| m0 Bootstrap Thread | `#ffe3e3` | 启动 Go Runtime 的第一个 OS 线程。 |
| G State | `#f8f0fc` | G 生命周期状态的可复用标记。 |
| _Grunnable | `#a5d8ff` | 已在运行队列中，等待获得 P 与 M 执行。 |
| _Grunning | `#d3f9d8` | 已绑定 M 与 P、正在执行 Go 代码。 |
| _Gwaiting | `#fff3bf` | 因同步、网络或定时器事件而等待。 |
| _Gsyscall | `#e5dbff` | 正在执行系统调用，暂时脱离 P。 |
| _Gdead | `#e9ecef` | 已结束执行，可被复用或回收。 |
| _Gpreempted | `#ffe8cc` | 被调度器抢占，等待重新调度。 |
| _Gidle | `#f1f3f5` | 刚分配但尚未初始化的 goroutine。 |
| _Gcopystack | `#d0bfff` | 运行时正在移动该 G 的栈。 |
| _Gleaked | `#ffc9c9` | GC 检测到的疑似泄漏 goroutine。 |
| _Gdeadextra | `#ffd8a8` | 供 cgo 回调 extra M 使用的空闲 G。 |

在 Excalidraw 的 Library 面板中导入 `developer_library/golang/runtime/gmp/golang_runtime_gmp_library.excalidrawlib`。

## 图标规范

遵循 [Go Runtime 图标总设计规范](../ICON_DESIGN_SPEC.md)。

- Goroutine 使用约 `56 × 44 px` 的青色圆角任务卡，居中 `go` 标签配小型启动箭头，表达 `go` 语句创建的执行单元。
- G 使用约 `52 × 56 px` 的蓝色圆角上下文卡，顶部为 `G` 标签、下方三条栈帧线；栈帧线仅表达执行上下文和栈，不表示具体调用深度。
- M 使用约 `52 × 56 px` 的紫色线程承载卡，顶部为 `M` 标签、内部纵向线程轨道及两条执行位，表达 OS 线程承载关系。
- P 使用约 `52 × 52 px` 的黄色处理器芯片，带四个资源引脚、`P` 标签和三个白色本地 runq 槽位；槽位只表达“拥有本地队列”，不表示队列容量。
- g0 使用约 `52 × 56 px` 的蓝灰色系统栈卡，带白色内层栈区和两条栈帧线；仅用于调度栈和 systemstack 语义，不替代普通 G。
- m0 使用约 `52 × 56 px` 的浅红色启动线程卡，带线程轨道和启动箭头；仅用于程序启动阶段的首个 OS 线程。
- G State 使用约 `60 × 44 px` 的浅紫状态卡，左侧为 G、右侧为四个状态槽；状态名称由使用图文字指定，槽位不编码固定状态颜色。
- 十个活动 G 状态使用 `60 × 44 px` 状态卡与下方 `12.518 px` 等宽状态标签组成的图标；左侧固定为 `G`，右侧以微型语义符号表达生命周期状态，且 Library 与 catalog 保持一致。
- 身份色固定为：Goroutine `#e3fafc`、G `#d0ebff`、M `#e5dbff`、P `#fff3bf`。描边、文字和内部结构线遵循总规范；GMP 对象之间的绑定、运行和迁移关系仍使用架构图的外部连线表达。

## G 状态转换组合

每个转换组合均为独立 Library item：不使用共同外框，左侧放来源状态卡，右侧放目标状态卡，中间使用 `#1f2937`、`2 px` 实线箭头。箭头上方为无描边浅灰胶囊“状态变化”，下方为无描边浅灰胶囊，显示该迁移的触发说明；来源和目标卡直接复用各自状态的微型语义符号。

| 状态 | 右侧语义符号 |
| --- | --- |
| _Grunning | 播放与执行指向 |
| _Grunnable | 队列条目与进入方向 |
| _Gwaiting | 暂停与等待标记 |
| _Gsyscall | 向外的调用标记 |
| _Gdead | 四个叉号 |
| _Gpreempted | 被斜线打断的执行标记 |
| _Gidle | 四个空心圆点 |
| _Gcopystack | 栈层复制与迁移 |
| _Gleaked | 告警与脱离点 |
| _Gdeadextra | cgo 标记与空闲圆点 |

已收录的基础状态迁移：`_Gidle → _Gdead`、`_Gidle → _Gdeadextra`、`_Gdead → _Grunnable`、`_Gdead → _Gwaiting`、`_Grunnable → _Grunning`、`_Grunning → _Grunnable`、`_Grunning → _Gwaiting`、`_Gwaiting → _Grunnable`、`_Gwaiting → _Grunning`、`_Grunning → _Gsyscall`、`_Gsyscall → _Grunning`、`_Grunning → _Gdead`、`_Grunning → _Gcopystack`、`_Gcopystack → _Grunning`、`_Grunning → _Gpreempted`、`_Gpreempted → _Gwaiting`、`_Gwaiting → _Gleaked`、`_Gleaked → _Gwaiting`、`_Gdeadextra → _Gsyscall`、`_Gsyscall → _Gdeadextra`。

不为 `_Gscan*` 单独绘制状态卡：它是 GC 扫描叠加位，而非独立生命周期节点。`_Gmoribund_unused` 与 `_Genqueue_unused` 是未使用保留常量，也不纳入图标库。`_Gsyscall` 在未取得 P 时会短暂进入 `_Grunning` 后再变为 `_Grunnable`，因此用两个真实基础迁移表达。
