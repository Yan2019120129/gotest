# Go Runtime GC Excalidraw Library

该图标库用于绘制 Go Runtime 垃圾回收、标记、清扫和 STW 相关流程。

## 文件

- `golang_runtime_gc_library.excalidrawlib`：可直接导入 Excalidraw 的 GC 图标库。
- `golang_runtime_gc_icons.excalidraw`：GC 图标浏览图板。

## 图标

| 图标 | 填充色 | 用途说明 |
| --- | --- | --- |
| GC | `#d3f9d8` | 垃圾回收、标记清扫和 STW 相关流程。 |

在 Excalidraw 的 Library 面板中导入 `developer_library/golang/runtime/gc/golang_runtime_gc_library.excalidrawlib`。

## 图标规范

遵循 [Go Runtime 图标总设计规范](../ICON_DESIGN_SPEC.md)。

- GC 使用约 `44 × 45 px` 的圆形主体，填充色固定为 `#d3f9d8`，标签为居中的 `GC`、`12 px`。
- 仅使用一条内部箭头表达 GC 周期或回收流向；箭头不得延伸出图标边界，也不得替代使用图中的流程连线。
- 新增 GC 相关变体保持圆形与箭头语义；只有需要表达不同回收阶段时才增加白色内部细节。
