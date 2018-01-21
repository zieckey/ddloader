package ddl

import (
    "testing"
    "github.com/bmizerany/assert"
)

func TestConfig_Parse(t *testing.T) {
    c := NewConfig()
    assert.NotEqual(t, c, nil)

    err := c.Parse("../etc/example.ini")
    assert.Equal(t, err, nil)

    assert.Equal(t, len(c.DataItems), 2)

    assert.Equal(t, c.OnUpdatedShellCommand, "reload_nginx; ls -a;")
    assert.Equal(t, c.OnInitShellCommand, "ls -l; uname -n")
    assert.Equal(t, c.OnExitShellCommand, "uname -n; ls -l;")

    d, ok := c.DataItems["remote_http_file_without_md5"]
    assert.Equal(t, ok, true)
    assert.Equal(t, d.Name, "remote_http_file_without_md5")
    assert.Equal(t, d.DataSourceAddress, "https://www.toutiao.com/robots.txt")
    assert.Equal(t, d.DestinationPath, "/tmp/toutiao.robots.txt")
    assert.Equal(t, d.ShellCommand, "rm -f /tmp/toutiao.robots.txt")
    assert.Equal(t, d.CheckMD5BeforeDownload, false)

    d, ok = c.DataItems["local_file"]
    assert.Equal(t, ok, true)
    assert.Equal(t, d.Name, "local_file")
    assert.Equal(t, d.DataSourceAddress, "/bin/ls")
    assert.Equal(t, d.DestinationPath, "/tmp/ddloader.ls.bin.tmp")
    assert.Equal(t, d.ShellCommand, "mv /tmp/ddloader.ls.bin.tmp /tmp/ddloader.ls.bin; rm -rf /tmp/ddloader.ls.bin")
    assert.Equal(t, d.CheckMD5BeforeDownload, false)
}
