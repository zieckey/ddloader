package main

import (
    "github.com/zieckey/ddloader/ddl"
    "os"
)

func main() {
    m := new(ddl.Main)
    err := m.Init()
    if err != nil {
        ddl.Fatalf("Init failed. %v", err.Error())
        os.Exit(-1)
    }

    err = m.Run()
    if err != nil {
        ddl.Fatalf("Run failed. %v", err.Error())
        os.Exit(-2)
    }

    err = m.Uninit()
    if err != nil {
        ddl.Fatalf("Uninit failed. %v", err.Error())
        os.Exit(-3)
    }

    os.Exit(0)
}
