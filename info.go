package goinfo

// Xminfo info finder.
type Xminfo struct {
}

var gTimeFarmat = "2006-01-02 15:04:05"

var (
	gCPUFile     = "/proc/stat"
	gMemFile     = "/proc/meminfo"
	gNetFile     = "/proc/net/dev"
	gHostFile    = "/proc/uptime"
	gDiskFile    = "/proc/diskstats"
	gLoadavgFile = "/proc/loadavg"
	gSnmpFile    = "/proc/net/snmp"
)

// NewXminfo new a info finder.
func NewXminfo() (xminfo *Xminfo) {
	return &Xminfo{}
}

var defaultInfoFinder = NewXminfo()

// CPU return the default info finder's CPU info.
func CPU() (info []*CPUInfo, err error) {
	return defaultInfoFinder.CPU()
}

// Disk return the default info finder's Disk info.
func Disk() (disk []*DiskStat, err error) {
	return defaultInfoFinder.Disk()
}

// Host return the default info finder's Host info.
func Host() (host *HostName, err error) {
	return defaultInfoFinder.Host()
}

// Load return the default info finder's Load info.
func Load() (load *Loadavg, err error) {
	return defaultInfoFinder.Load()
}

// Memory return the default info finder's Memory info.
func Memory() (free *Free, err error) {
	return defaultInfoFinder.Memory()
}

// Net return the default info finder's Net info.
func Net() (traffic []*Traffic, err error) {
	return defaultInfoFinder.Net()
}

// TCP return the default info finder's tcp info.
func TCP() (tcp *Tcp, err error) {
	return defaultInfoFinder.TCP()
}

// UDP return the default info finder's udp info.
func UDP() (udp *Udp, err error) {
	return defaultInfoFinder.UDP()
}
