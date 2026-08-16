package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// ProgressWriter 进度条统计器
type ProgressWriter struct {
	Total      int64
	Downloaded int64
	StartTime  time.Time
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Downloaded += int64(n)
	pw.render()
	return n, nil
}

func (pw *ProgressWriter) render() {
	const barWidth = 30
	var percent float64
	var bar string

	if pw.Total > 0 {
		percent = float64(pw.Downloaded) / float64(pw.Total) * 100
		completed := int(float64(barWidth) * (float64(pw.Downloaded) / float64(pw.Total)))
		if completed > barWidth {
			completed = barWidth
		}

		if completed == barWidth {
			bar = strings.Repeat("=", barWidth)
		} else {
			bar = strings.Repeat("=", completed) + ">" + strings.Repeat(" ", barWidth-completed-1)
		}
	} else {
		bar = "大小未知..."
	}

	// 计算瞬时下载速度
	elapsed := time.Since(pw.StartTime).Seconds()
	var speed float64
	if elapsed > 0 {
		speed = float64(pw.Downloaded) / 1024 / 1024 / elapsed // MB/s
	}

	// \r 返回行首覆盖输出
	fmt.Printf("\r[%s] %6.2f%% | %7.2f MB / %7.2f MB | %5.2f MB/s",
		bar,
		percent,
		float64(pw.Downloaded)/(1024*1024),
		float64(pw.Total)/(1024*1024),
		speed,
	)
}

func DownloadWithProgressBar(url, savePath string) error {
	// 1. 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，状态码: %s", resp.Status)
	}

	// 2. 创建本地文件
	out, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	// 3. 初始化进度监听器 (resp.ContentLength 获取总大小)
	pw := &ProgressWriter{
		Total:     resp.ContentLength,
		StartTime: time.Now(),
	}

	// 4. 使用 io.TeeReader 同时向文件和进度条分发数据流
	_, err = io.Copy(out, io.TeeReader(resp.Body, pw))
	if err != nil {
		return fmt.Errorf("写入失败: %w", err)
	}

	fmt.Println("\n下载完成！")
	return nil
}

func main() {
	url := "https://speed.hetzner.de/100MB.bin"
	if err := DownloadWithProgressBar(url, "download.bin"); err != nil {
		fmt.Printf("\n错误: %v\n", err)
	}
}
