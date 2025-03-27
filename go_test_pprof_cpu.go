package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	Analysis()
}

func Analysis() {
	// 创建 CPU 分析文件
	cpuProfile, err := os.Create("./profile/cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuProfile.Close()

	// 开始 CPU 分析
	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// 模拟一些 CPU 密集型工作
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}

	// 创建内存分析文件
	memProfile, err := os.Create("./profile/mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer memProfile.Close()

	// 强制进行垃圾回收，以便获取准确的内存分析数据
	runtime.GC()

	// 开始内存分析
	if err := pprof.WriteHeapProfile(memProfile); err != nil {
		log.Fatal(err)
	}

	// 模拟一些内存使用
	data := make([]byte, 1024*1024)
	_ = data

	time.Sleep(100 * time.Second) // 等待一段时间以便生成分析数据

	log.Println("完成性能分析")
}
