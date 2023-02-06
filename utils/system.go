package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
)

var AppStartTime time.Time
var SysStatus struct {
	Uptime       int64
	NumGoroutine int //当前 Goroutines 数量

	// General statistics.
	MemAllocated string // bytes allocated and still in use 当前内存使用量
	MemTotal     string // bytes allocated (even if freed) 所有被分配的内存
	MemSys       string // bytes obtained from system (sum of XxxSys below) 内存占用量
	Lookups      uint64 // number of pointer lookups 指针查找次数
	MemMallocs   uint64 // number of mallocs 内存分配次数
	MemFrees     uint64 // number of frees 内存释放次数

	// Main allocation heap statistics.
	HeapAlloc    string // bytes allocated and still in use 当前 Heap 内存使用量
	HeapSys      string // bytes obtained from system Heap 内存占用量
	HeapIdle     string // bytes in idle spans Heap 内存空闲量
	HeapInuse    string // bytes in non-idle span 正在使用的 Heap 内存
	HeapReleased string // bytes released to the OS 被释放的 Heap 内存
	HeapObjects  uint64 // total number of allocated objects Heap 对象数量

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // bootstrap stacks 启动 Stack 使用量
	StackSys    string // 被分配的 Stack 内存
	MSpanInuse  string // mspan structures MSpan 结构内存使用量
	MSpanSys    string // 被分配的 MSpan 结构内存
	MCacheInuse string // mcache structures MCache 结构内存使用量
	MCacheSys   string // 被分配的 MCache 结构内存
	BuckHashSys string // profiling bucket hash table 被分配的剖析哈希表内存
	GCSys       string // GC metadata 被分配的 GC 元数据内存
	OtherSys    string // other system allocations 其它被分配的系统内存

	// Garbage collector statistics.
	NextGC       string // next run in HeapAlloc time (bytes) 下次 GC 内存回收量
	LastGC       string // last run in absolute time (ns) 距离上次 GC 时间
	PauseTotalNs string // GC 暂停时间总量
	PauseNs      string // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256] 上次 GC 暂停时间
	NumGC        uint32 //GC 执行次数
}

func TimeSincePro(then time.Time) int64 {
	diff := time.Now().Unix() - then.Unix()
	return diff
}

func UpdateSystemStatus() {
	SysStatus.Uptime = TimeSincePro(AppStartTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	SysStatus.NumGoroutine = runtime.NumGoroutine()

	SysStatus.MemAllocated = FileSize(int64(m.Alloc))
	SysStatus.MemTotal = FileSize(int64(m.TotalAlloc))
	SysStatus.MemSys = FileSize(int64(m.Sys))
	SysStatus.Lookups = m.Lookups
	SysStatus.MemMallocs = m.Mallocs
	SysStatus.MemFrees = m.Frees

	SysStatus.HeapAlloc = FileSize(int64(m.HeapAlloc))
	SysStatus.HeapSys = FileSize(int64(m.HeapSys))
	SysStatus.HeapIdle = FileSize(int64(m.HeapIdle))
	SysStatus.HeapInuse = FileSize(int64(m.HeapInuse))
	SysStatus.HeapReleased = FileSize(int64(m.HeapReleased))
	SysStatus.HeapObjects = m.HeapObjects

	SysStatus.StackInuse = FileSize(int64(m.StackInuse))
	SysStatus.StackSys = FileSize(int64(m.StackSys))
	SysStatus.MSpanInuse = FileSize(int64(m.MSpanInuse))
	SysStatus.MSpanSys = FileSize(int64(m.MSpanSys))
	SysStatus.MCacheInuse = FileSize(int64(m.MCacheInuse))
	SysStatus.MCacheSys = FileSize(int64(m.MCacheSys))
	SysStatus.BuckHashSys = FileSize(int64(m.BuckHashSys))
	SysStatus.GCSys = FileSize(int64(m.GCSys))
	SysStatus.OtherSys = FileSize(int64(m.OtherSys))

	SysStatus.NextGC = FileSize(int64(m.NextGC))
	SysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	SysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	SysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	SysStatus.NumGC = m.NumGC
}

func FileSize(s int64) string {
	return humanize.IBytes(uint64(s))
}
