# PostgreSQL Storage Icons Index

| 分类 | 图标名称 | 中文解释 | 使用场景 | 推荐颜色 | 可连接关系 |
| --- | --- | --- | --- | --- | --- |
| 逻辑层 | Database | 数据库实例或逻辑数据库。 | 表示 PostgreSQL database 边界。 | `#e7f5ff` | Database -> Schema |
| 逻辑层 | Schema | 命名空间和对象分组。 | 表示 public/schema 等对象容器。 | `#e7f5ff` | Schema -> Table |
| 逻辑层 | Table | 用户表的逻辑结构。 | 存储结构主链路入口。 | `#e7f5ff` | Table -> Relation |
| 逻辑层 | Relation | PostgreSQL relation 抽象。 | 连接逻辑表和物理 relfilenode。 | `#e7f5ff` | Relation -> Relation File |
| 物理文件层 | Relation File | relation 对应的物理文件。 | 表示 relfilenode 主文件。 | `#fff4e6` | Relation -> Relation File -> Segment File |
| 物理文件层 | Segment File | 超过 1GB 后拆分的段文件。 | 表示 relfilenode.1、relfilenode.2 等。 | `#fff4e6` | Relation File -> Segment File |
| 物理文件层 | Fork File | relation fork 文件抽象。 | 统一表示 main/fsm/vm/init fork。 | `#fff4e6` | Relation File -> Fork File |
| 物理文件层 | Main Fork | 存放 heap/index 主数据的 fork。 | 连接 page/block 主体。 | `#fff4e6` | Main Fork -> Page / Block (8KB) |
| 物理文件层 | FSM Fork | Free Space Map 文件。 | 表示页面可用空间索引。 | `#fff4e6` | FSM Fork -> Free Space Map |
| 物理文件层 | VM Fork | Visibility Map 文件。 | 表示 all-visible/all-frozen 位图。 | `#fff4e6` | VM Fork -> Visibility Map |
| 物理文件层 | TOAST Table | 大字段外部存储表。 | 表示 TOAST relation。 | `#fff4e6` | TOAST Pointer -> TOAST Table |
| Page / Block 层 | Page / Block (8KB) | PostgreSQL 默认 8KB 数据页。 | 主链路中的 page/block 节点。 | `#f3f0ff` | Segment File -> Page / Block (8KB) |
| Page / Block 层 | Heap Page | heap 表页面内部结构。 | 展示 header、line pointer、tuple area。 | `#f3f0ff` | Heap Page -> Page Header |
| Page / Block 层 | Page Header | 页头元数据区域。 | 表示 pd_lsn、pd_lower、pd_upper 等。 | `#f3f0ff` | Page Header -> Line Pointer Array |
| Page / Block 层 | Line Pointer Array | 行指针数组。 | 表示页面中 ItemId 列表。 | `#f3f0ff` | Line Pointer Array -> ItemId |
| Page / Block 层 | ItemId | 行指针槽位。 | 表示 tuple offset/length/state。 | `#f3f0ff` | ItemId -> Heap Tuple |
| Page / Block 层 | Free Space | 页面空闲空间区域。 | 表示 pd_lower 和 pd_upper 之间的空间。 | `#f3f0ff` | Free Space -> Free Space Map |
| Page / Block 层 | Tuple Data Area | tuple 实际存放区域。 | 表示 heap tuple 从页尾向前分配。 | `#f3f0ff` | Tuple Data Area -> Heap Tuple |
| Tuple / Row 层 | Heap Tuple | heap 表中的一行物理记录。 | 连接 page 和 tuple header/user data。 | `#e6fcf5` | ItemId -> Heap Tuple |
| Tuple / Row 层 | Tuple Header | tuple 头部元数据。 | 表示事务、ctid、标志位等。 | `#e6fcf5` | Heap Tuple -> Tuple Header |
| Tuple / Row 层 | xmin | 插入该 tuple 的事务 ID。 | 讲解 MVCC 可见性。 | `#e6fcf5` | Tuple Header -> xmin |
| Tuple / Row 层 | xmax | 删除或锁定该 tuple 的事务 ID。 | 讲解 UPDATE/DELETE 可见性。 | `#e6fcf5` | Tuple Header -> xmax |
| Tuple / Row 层 | ctid | 当前 tuple 的物理位置指针。 | 连接版本链或索引回表。 | `#e6fcf5` | Tuple Header -> ctid |
| Tuple / Row 层 | infomask | tuple 状态标志位。 | 展示 hint bits、锁状态、null 标记等。 | `#e6fcf5` | Tuple Header -> infomask |
| Tuple / Row 层 | null bitmap | NULL 字段位图。 | 表示可空列的 null 标记。 | `#e6fcf5` | Tuple Header -> null bitmap |
| Tuple / Row 层 | User Data | 用户列数据区域。 | 连接 tuple 到 column value。 | `#e6fcf5` | Heap Tuple -> User Data |
| Tuple / Row 层 | Column Value | 单个列值。 | Table -> Page -> Tuple -> Column 的终点。 | `#fff9db` | User Data -> Column Value |
| Tuple / Row 层 | TOAST Pointer | 指向 TOAST 外部值的指针。 | 展示大字段外部存储。 | `#fff9db` | Column Value -> TOAST Pointer -> TOAST Table |
| Index 层 | Index | 索引对象。 | 表示 btree/hash/gin/gist 等索引入口。 | `#e7f5ff` | Table -> Index |
| Index 层 | Index Page | 索引页。 | 表示索引内部 page/block。 | `#f3f0ff` | Index -> Index Page |
| Index 层 | Index Tuple | 索引条目。 | 表示 key + heap TID。 | `#e6fcf5` | Index Page -> Index Tuple |
| Index 层 | TID | 指向 heap tuple 的 block/offset。 | 索引回表定位。 | `#fff9db` | Index Tuple -> TID |
| Index 层 | B-Tree Page | B-Tree 节点页。 | 展示 root/internal/leaf 层次。 | `#f3f0ff` | Index -> B-Tree Page |
| Index 层 | Heap Pointer | 从索引指向 heap tuple 的指针。 | 表示 index scan 回表。 | `#fff9db` | TID -> Heap Tuple |
| MVCC / Vacuum 层 | Old Tuple | 更新前的旧版本 tuple。 | 展示版本链和可见性判断。 | `#ffe3e3` | Old Tuple -> New Tuple |
| MVCC / Vacuum 层 | New Tuple | UPDATE 后的新版本 tuple。 | 展示 HOT/non-HOT 更新。 | `#e6fcf5` | Old Tuple -> New Tuple |
| MVCC / Vacuum 层 | Dead Tuple | 不再对任何事务可见的 tuple。 | VACUUM 可清理对象。 | `#ffe3e3` | Dead Tuple -> VACUUM |
| MVCC / Vacuum 层 | Live Tuple | 当前可见或仍需保留的 tuple。 | 表示可见数据版本。 | `#e6fcf5` | Visibility Map -> Live Tuple |
| MVCC / Vacuum 层 | Update Version Chain | ctid 串联的新旧版本链。 | 展示 UPDATE 版本跳转。 | `#ffe3e3` | Old Tuple -> Update Version Chain -> New Tuple |
| MVCC / Vacuum 层 | VACUUM | 清理 dead tuple 并回收空间。 | 表示 vacuum 过程。 | `#ffe3e3` | VACUUM -> Free Space Map |
| MVCC / Vacuum 层 | Visibility Map | all-visible/all-frozen 位图。 | 优化 vacuum 和 index-only scan。 | `#ffe3e3` | VM Fork -> Visibility Map |
| MVCC / Vacuum 层 | Free Space Map | 页面可用空间索引。 | 辅助插入选择有空闲空间的页。 | `#ffe3e3` | FSM Fork -> Free Space Map |
| OS / Storage 层 | OS File System | 操作系统文件系统。 | 表示 ext4/xfs 等文件系统层。 | `#f1f3f5` | Tablespace -> OS File System |
| OS / Storage 层 | Disk | 磁盘或块设备。 | 表示最终持久化介质。 | `#f1f3f5` | OS File System -> Disk |
| OS / Storage 层 | Data Directory | PostgreSQL 数据目录 PGDATA。 | 表示 base/global/pg_wal 等目录入口。 | `#f1f3f5` | Data Directory -> Relation File |
| OS / Storage 层 | Tablespace | 表空间目录或外部存储位置。 | 表示 pg_tblspc 链接目标。 | `#f1f3f5` | Tablespace -> Data Directory |
