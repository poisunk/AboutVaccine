package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"vax/internal/config"
)

var (
	rootCmd = &cobra.Command{
		Use: "app",
	}

	initCmd = &cobra.Command{
		Use: "init",
		Run: func(_ *cobra.Command, _ []string) {
			InitConfigs()
		},
	}

	runCmd = &cobra.Command{
		Use: "run",
		Run: func(_ *cobra.Command, _ []string) {
			RunApp()
		},
	}
)

func init() {
	for _, cmd := range []*cobra.Command{initCmd, runCmd} {
		rootCmd.AddCommand(cmd)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func InitConfigs() {
	conf := &config.AllConfig{}
	// 数据库配置
	driver := os.Getenv("DATABASE_DRIVER")
	connection := os.Getenv("DATABASE_CONNECTION")
	if len(driver) == 0 || len(connection) == 0 {
		fmt.Println("使用默认数据库配置")
		err := InstallSqlite()
		if err != nil {
			os.Exit(1)
		}
		driver = "sqlite"
		connection = "/data/vax.db"
	}
	conf.Database = &config.Database{
		Driver:     driver,
		Connection: connection,
	}
	// 服务器配置
	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		fmt.Println("使用默认服务器配置")
		fmt.Println("默认端口: 8080")
		port = "8080"
	}
	conf.Server = &config.Server{
		Port: port,
	}
	// 保存配置
	err := config.SaveConfig(conf)
	if err != nil {
		fmt.Println("保存配置失败")
		os.Exit(1)
	}
}

func InstallSqlite() error {
	cmd := exec.Command("apk", "add", "--no-cache", "sqlite")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("SQLite 安装失败: %s", err)
	}
	fmt.Println("SQLite 安装成功")
	return nil
}
