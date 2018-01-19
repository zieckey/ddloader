package ddl

type Porter interface {

    // 检查数据源是否有更新
    HasUpdate() bool

    // 从数据源地址获取数据到并保存到目标地址
    Fetch() bool
}

