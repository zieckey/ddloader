package ddl

func init() {
    SetAlarm(DefaultAlarmCallback)
    SetLogger(new(GLog))
}

type Main struct {
    C *Config
}


// 几个注意事项：
// 1. 全程注意参考PHP代码
// 2. 注意记录各种日志数据
// 3. 注意将日志记录、报警等事项留下callback，供业务方自定义


func (m* Main) Init() error {
    // 0. 从命令行参数中获取 配置文件 路径
    // 1. 解析配置文件得到 Config 对象
    // 2. 执行 Config.OnInitShellCommand
    return nil
}

func (m* Main) Run() error {
    // 3. 遍历 Config.DataItems
    //      a. 挨个读取 DataItem
    //      b. 根据 DataItem 数据成员 DataSourceAddress 选择合适的 Fetcher，并下载数据，并放置到 DestinationPath（注意：先写临时文件，然后用move指令）
    //      c. 下载成功后，执行指定的shell命令 DataItem.ShellCommand
    // 4. 等所有的 DataItem 都下载完成后，执行全局shell命令 Config.OnUpdatedShellCommand
    return nil
}

func (m* Main) Uninit() error {
    // 5. 进程退出时，执行 Config.OnExitShellCommand
    return nil
}


