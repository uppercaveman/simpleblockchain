package main

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"simpleblockchain/blockchain"
	"simpleblockchain/modules/log"
	pow "simpleblockchain/proofofwork"
	"simpleblockchain/setting"
)

var configPath string

func init() {
	pwd, _ := os.Getwd()
	flag.StringVar(&configPath, "c", filepath.Join(pwd, "app.toml"), "-c /path/to/app.toml config gile")
}

func bootstrap() {
	if err := setting.InitConf(configPath); err != nil {
		panic(err.Error())
	}

	// 初始化日志系统
	log.InitLogService(setting.Conf.Logs)
}

func main() {
	bootstrap()

	// 创建区块链
	bc := blockchain.NewBlockchain()
	// 创建区块
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		log.Info("PrevBlockHash. hash: %x", block.PrevBlockHash)
		log.Info("Data: %s", block.Data)
		log.Info("Hash: %x", block.Hash)
		log.Info("POW Validate : %s\n", strconv.FormatBool(pow.NewProofOfWork(block).Validate()))
	}

	initSignal()
}

func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		log.Info("market_shell signal %s \n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			continue
		default:
			return
		}
	}
}
