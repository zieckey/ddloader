package ddl

import (
    "github.com/zieckey/goini"
    "fmt"
    "errors"
)

// 数据项是最基本的配置单元，代表了一个数据文件的远端数据源地址、本地目标路径、更新后的执行动作等信息
type DataItem struct {
    // 该数据项的名称. 一般而言对应INI配置文件的节的名称
    Name string

    // 数据文件的下载源地址, 支持多种形式的地址：
    //   1. 本地文件目录
    //      - /data/download/project/rule/*.dat
    //   2. HTTP
    //      - http://10.12.12.3:8080/rule.dat
    DataSourceAddress string

    // 是否优先下载远端md5文件来检查是否有必要更新
    CheckMD5BeforeDownload bool

    // 下载到文件后，将最新的数据文件存放的目标位置
    // 例如：/var/ddl/project/rule.dat
    DestinationPath string

    // 下载成功后，执行的shell命令
    // 支持管道
    // 例如：ls -l | wc
    ShellCommand string

    // 下载超时时间，单位秒
    Timeout int

    // 失败重试次数
    Retry int

    // 失败重试间隔，单位秒
    RetryInterval int
}

type Config struct {
    INI* goini.INI

    // 启动时执行的命令, 只在启动时执行一次
    OnInitShellCommand string

    // 程序退出时执行的命令，只在退出时执行一次
    OnExitShellCommand string

    // 当前这次主循环逻辑完成后，如果任意一个数据项有更新的话就会执行该命令
    OnUpdatedShellCommand string

    // 数据项
    DataItems map[string]DataItem
}

func NewConfig() *Config {
    c := &Config{}
    c.DataItems = make(map[string]DataItem)
    return c
}

func (c *Config) ParseDataItem(section string) (DataItem, error) {
    var d DataItem
    var found bool
    d.Name = section
    d.DataSourceAddress, found = c.INI.SectionGet(section, "DataSourceAddress")
    if !found {
        s := fmt.Sprintf("Cannot find %s.%s", section, "DataSourceAddress")
        Errorf(s)
        return d, errors.New(s)
    }

    d.DestinationPath, found = c.INI.SectionGet(section, "DestinationPath")
    if !found {
        s := fmt.Sprintf("Cannot find %s.%s", section, "DestinationPath")
        Errorf(s)
        return d, errors.New(s)
    }

    d.CheckMD5BeforeDownload, _ = c.INI.SectionGetBool(section, "DataSourceAddress")
    d.ShellCommand, _ = c.INI.SectionGet(section, "ShellCommand")
    d.Timeout, _ = c.INI.SectionGetInt(section, "Timeout")
    d.Retry, _ = c.INI.SectionGetInt(section, "Retry")
    d.RetryInterval, _ = c.INI.SectionGetInt(section, "RetryInterval")

    return d, nil
}

func (c *Config) Parse(conf string) error {
    ini := goini.New()
    ini.SetTrimQuotes(true)
    err := ini.ParseFile(conf)
    if err != nil {
        Errorf("Parse INI config file %v failed : %v", conf, err.Error())
        return err
    }

    c.INI = ini
    c.OnInitShellCommand, _ = ini.Get("OnInitShellCommand")
    c.OnExitShellCommand, _ = ini.Get("OnExitShellCommand")
    c.OnUpdatedShellCommand, _ = ini.Get("OnUpdatedShellCommand")

    m := ini.GetAll()
    for sec := range m {
        if sec == "" {
            continue
        }

        d, err := c.ParseDataItem(sec)
        if err != nil {
            Errorf("%s", err.Error())
            return err
        }

        c.DataItems[d.Name] = d
    }

    return nil
}
