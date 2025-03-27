package main

import (
	"fmt"
	"os"
	"sync"
)

// 配置参数
const (
	sourceFilePath = "/path/to/source/file"
	destFilePath   = "/path/to/destination/file"
	chunkSize      = 64 * 1024 * 1024 // 64MB 每个块的大小
	numThreads     = 2                // 使用 2 个线程
)

// copyChunk 负责复制文件的一部分（块）
func copyChunk(wg *sync.WaitGroup, start, end int64, sourceFile *os.File, destFile *os.File) {
	defer wg.Done()

	// 使用文件指针定位
	_, err := sourceFile.Seek(start, 0)
	if err != nil {
		fmt.Printf("Error seeking source file: %v\n", err)
		return
	}

	_, err = destFile.Seek(start, 0)
	if err != nil {
		fmt.Printf("Error seeking destination file: %v\n", err)
		return
	}

	// 读取源文件块并写入目标文件
	buffer := make([]byte, chunkSize)
	for start < end {
		// 计算当前读取的块大小
		readSize := chunkSize
		if end-start < chunkSize {
			readSize = end - start
		}

		// 从源文件读取数据
		n, err := sourceFile.Read(buffer[:readSize])
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("Error reading source file: %v\n", err)
			return
		}

		// 写入目标文件
		_, err = destFile.Write(buffer[:n])
		if err != nil {
			fmt.Printf("Error writing to destination file: %v\n", err)
			return
		}

		// 更新读取的偏移量
		start += int64(n)
	}
}

// copyLargeFile 负责管理整个文件的复制过程
func copyLargeFile() error {
	// 获取源文件的大小
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destFilePath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %v", err)
	}

	fileSize := sourceFileInfo.Size()
	chunkCount := int(fileSize / chunkSize)
	if fileSize%chunkSize != 0 {
		chunkCount++ // 需要额外的块来处理剩余数据
	}

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		// 计算每个 goroutine 负责的文件块范围
		start := int64(i * int(fileSize/int64(numThreads)))
		end := start + int64(fileSize/int64(numThreads))
		if i == numThreads-1 {
			end = fileSize
		}

		// 启动 goroutine 复制文件块
		wg.Add(1)
		go copyChunk(&wg, start, end, sourceFile, destFile)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	return nil
}

func main() {
	err := copyLargeFile()
	if err != nil {
		fmt.Printf("Error during file copy: %v\n", err)
	} else {
		fmt.Println("File copy completed successfully.")
	}
}
