# Developer Excalidraw Library

这是一套面向软件开发、系统架构、AI Agent、数据库和云原生场景的 Excalidraw 手绘风格图标库。

## 文件

- `developer_library.excalidrawlib`: 可直接导入 Excalidraw 的 Library 文件。
- `developer_library_catalog.excalidraw`: 全量图标浏览图板，包含名称和用途说明。

Go Runtime 图标按功能拆分为独立子库：[`GC`](golang/runtime/gc/README.md)、[`GMP`](golang/runtime/gmp/README.md)、[`Memory`](golang/runtime/memory/README.md) 和 [`Scheduler`](golang/runtime/scheduler/README.md)。

## 风格规范

- 图标画布约 80x80。
- 统一深色描边 `#1f2937`，线宽 2，Excalidraw roughness 为 1。
- 使用低饱和分类底色，避免在架构图中过度抢占视觉焦点。
- Library 图标内只保留必要缩写，完整名称和用途在 catalog 与本清单中维护。

## 布局建议

- 架构图从左到右组织：人物/客户端 -> 网络入口 -> 服务与 Kubernetes -> 消息队列/数据库 -> 监控。
- AI Agent 图建议按闭环组织：Prompt -> Planner -> Tool/MCP/Skill -> Executor -> Observation -> Reflection -> Memory/RAG/VectorDB。
- 图标间距建议 32-48，连线使用同色 2px 手绘箭头，业务主链路用实线，异步或观测链路用虚线。

## 图标清单

### 人物

| 图标 | 用途说明 |
| --- | --- |
| User | 普通终端用户或业务使用者。 |
| Admin | 平台管理员、运维管理员或权限管理角色。 |
| Developer | 研发人员、代码贡献者或调试操作者。 |
| Client | 桌面客户端、CLI 或浏览器访问端。 |
| Mobile User | 移动端 App 用户或手机访问入口。 |

### 服务器

| 图标 | 用途说明 |
| --- | --- |
| Server | 通用物理服务器或服务节点。 |
| Application Server | 承载业务应用进程的应用服务器。 |
| Linux Host | Linux 主机、裸机或云主机运行环境。 |
| VM | 虚拟机实例或虚拟化运行单元。 |
| Container | 容器运行实例或隔离执行环境。 |
| Docker | Docker 容器、镜像或 Docker Engine。 |

### 网络

| 图标 | 用途说明 |
| --- | --- |
| Internet | 公网、外部网络或用户访问来源。 |
| Cloud | 云平台、云资源池或外部云服务。 |
| CDN | 内容分发网络、边缘缓存入口。 |
| API Gateway | API 网关、统一入口、鉴权与路由。 |
| Nginx | Nginx 反向代理、静态入口或边缘代理。 |
| Load Balancer | 负载均衡器或流量分发节点。 |
| DNS | 域名解析、服务发现或 DNS 入口。 |
| Firewall | 防火墙、安全边界或访问控制层。 |
| Router | 路由器、三层转发或网络出口。 |
| Switch | 交换机、二层网络或机房交换设备。 |

### 数据库

| 图标 | 用途说明 |
| --- | --- |
| MySQL | 关系型数据库 MySQL。 |
| PostgreSQL | 关系型数据库 PostgreSQL。 |
| Redis | 缓存、KV 存储、计数器或轻量队列。 |
| MongoDB | 文档数据库 MongoDB。 |
| ClickHouse | OLAP 分析数据库 ClickHouse。 |
| Elasticsearch | 搜索引擎与倒排索引存储。 |
| Milvus | 向量数据库 Milvus。 |
| pgvector | PostgreSQL 向量扩展或向量检索表。 |

### 消息队列

| 图标 | 用途说明 |
| --- | --- |
| Kafka | 高吞吐日志流、事件总线或消息队列。 |
| RabbitMQ | AMQP 消息队列、交换机与路由。 |
| RocketMQ | RocketMQ 事务消息、顺序消息或事件流。 |
| NATS | 轻量 Pub/Sub、JetStream 或服务消息。 |
| Redis Stream | Redis Stream 消息流和消费组。 |

### AI Agent

| 图标 | 用途说明 |
| --- | --- |
| LLM | 大语言模型推理节点或模型服务。 |
| Prompt | 提示词、系统指令或用户输入。 |
| Memory | 短期/长期记忆、会话状态或上下文存储。 |
| Tool | Agent 可调用工具、函数或外部动作。 |
| MCP | Model Context Protocol 服务或工具连接层。 |
| Skill | Codex/Agent Skill、可复用能力模块。 |
| Workflow | 多步骤工作流、编排链路或 DAG。 |
| Planner | 规划器、任务分解和步骤选择。 |
| Executor | 执行器、动作运行和工具调用。 |
| Reflection | 反思、自检、复盘和改进回路。 |
| Observation | 环境观察、工具返回和状态感知。 |
| VectorDB | 向量数据库、语义检索或 embedding 存储。 |
| Embedding | 文本向量化、语义表示或向量生成。 |
| RAG | 检索增强生成链路。 |

### Kubernetes

| 图标 | 用途说明 |
| --- | --- |
| Pod | Kubernetes Pod，最小调度单元。 |
| Service | K8s Service，稳定访问入口与负载均衡。 |
| Deployment | Deployment 控制器与副本滚动更新。 |
| Node | Kubernetes 工作节点或宿主机。 |
| Cluster | Kubernetes 集群边界。 |
| Namespace | 命名空间、资源隔离或环境边界。 |
| ConfigMap | 配置项、环境变量和非敏感配置。 |
| Secret | 密钥、证书和敏感配置。 |
| PVC | 持久卷声明和存储挂载。 |

### 监控

| 图标 | 用途说明 |
| --- | --- |
| Prometheus | 指标采集、PromQL 查询和告警规则。 |
| Grafana | 可视化看板、图表和运维大盘。 |
| Jaeger | 分布式追踪查询和调用链分析。 |
| OpenTelemetry | 可观测性 SDK、采集器和统一数据模型。 |
| Log | 日志记录、日志聚合或审计文本。 |
| Metrics | 指标序列、计数器、直方图和趋势图。 |
| Trace | 调用链、span 和跨服务请求路径。 |
