// +build solaris

package process

import (
	"context"
	"fmt"
	"io/ioutil"
	"syscall"
	"unsafe"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/internal/common"
	"github.com/shirou/gopsutil/net"
)

type MemoryMapsStat struct {
	Path         string `json:"path"`
	Rss          uint64 `json:"rss"`
	Size         uint64 `json:"size"`
	Pss          uint64 `json:"pss"`
	SharedClean  uint64 `json:"sharedClean"`
	SharedDirty  uint64 `json:"sharedDirty"`
	PrivateClean uint64 `json:"privateClean"`
	PrivateDirty uint64 `json:"privateDirty"`
	Referenced   uint64 `json:"referenced"`
	Anonymous    uint64 `json:"anonymous"`
	Swap         uint64 `json:"swap"`
}

type MemoryInfoExStat struct {
}

func Pids() ([]int32, error) {
	return PidsWithContext(context.Background())
}

func PidsWithContext(ctx context.Context) ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}

// NewProcess creates a new Process instance, it only stores the pid and
// checks that the process exists. Other method on Process can be used
// to get more information about the process. An error will be returned
// if the process does not exist.
func NewProcess(pid int32) (*Process, error) {
	p := &Process{
		Pid: int32(pid),
	}
	err := p.readPsinfo(context.Background())
	return p, err
}

func (p *Process) Ppid() (int32, error) {
	return p.parent, nil
}

func (p *Process) PpidWithContext(ctx context.Context) (int32, error) {
	return 0, common.ErrNotImplementedError
}

func readInt32(b []byte) int32 {
	return *(*int32)(unsafe.Pointer(&b[0]))
}

// From /usr/include/procfs.h
// #define PRARGSZ         80      /* number of chars of arguments */
// typedef struct psinfo {
//         int     pr_flag;        /* process flags (DEPRECATED; do not use) */
//         int     pr_nlwp;        /* number of active lwps in the process */
//         pid_t   pr_pid;         /* unique process id */
//         pid_t   pr_ppid;        /* process id of parent */
//         pid_t   pr_pgid;        /* pid of process group leader */
//         pid_t   pr_sid;         /* session id */
//         uid_t   pr_uid;         /* real user id */
//         uid_t   pr_euid;        /* effective user id */
//         gid_t   pr_gid;         /* real group id */
//         gid_t   pr_egid;        /* effective group id */
//         uintptr_t pr_addr;      /* address of process */
//         size_t  pr_size;        /* size of process image in Kbytes */
//         size_t  pr_rssize;      /* resident set size in Kbytes */
//         size_t  pr_pad1;
//         dev_t   pr_ttydev;      /* controlling tty device (or PRNODEV) */
//                         /* The following percent numbers are 16-bit binary */
//                         /* fractions [0 .. 1] with the binary point to the */
//                         /* right of the high-order bit (1.0 == 0x8000) */
//         ushort_t pr_pctcpu;     /* % of recent cpu time used by all lwps */
//         ushort_t pr_pctmem;     /* % of system memory used by process */
//         timestruc_t pr_start;   /* process start time, from the epoch */
//         timestruc_t pr_time;    /* usr+sys cpu time for this process */
//         timestruc_t pr_ctime;   /* usr+sys cpu time for reaped children */
//         char    pr_fname[PRFNSZ];       /* name of execed file */
//         char    pr_psargs[PRARGSZ];     /* initial characters of arg list */
//         int     pr_wstat;       /* if zombie, the wait() status */
//         int     pr_argc;        /* initial argument count */
//         uintptr_t pr_argv;      /* address of initial argument vector */
//         uintptr_t pr_envp;      /* address of initial environment vector */
//         char    pr_dmodel;      /* data model of the process */
//         char    pr_pad2[3];
//         taskid_t pr_taskid;     /* task id */
//         projid_t pr_projid;     /* project id */
//         int     pr_nzomb;       /* number of zombie lwps in the process */
//         poolid_t pr_poolid;     /* pool id */
//         zoneid_t pr_zoneid;     /* zone id */
//         id_t    pr_contract;    /* process contract */
//         int     pr_filler[1];   /* reserved for future use */
//         lwpsinfo_t pr_lwp;      /* information for representative lwp */
// } psinfo_t;

func (p *Process) readPsinfo(ctx context.Context) error {
	pid := p.Pid
	var statPath string

	statPath = fmt.Sprintf("/proc/%d/psinfo", pid)

	contents, err := ioutil.ReadFile(statPath)
	if err != nil {
		return err
	}

	// XXX assume 64-bit Solaris
	p.parent = readInt32(contents[8:12])

	return nil
}

func (p *Process) Name() (string, error) {
	return p.NameWithContext(context.Background())
}

func (p *Process) NameWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Tgid() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Exe() (string, error) {
	return p.ExeWithContext(context.Background())
}

func (p *Process) ExeWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Cmdline() (string, error) {
	return p.CmdlineWithContext(context.Background())
}

func (p *Process) CmdlineWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) CmdlineSlice() ([]string, error) {
	return p.CmdlineSliceWithContext(context.Background())
}

func (p *Process) CmdlineSliceWithContext(ctx context.Context) ([]string, error) {
	return []string{}, common.ErrNotImplementedError
}
func (p *Process) CreateTime() (int64, error) {
	return p.CreateTimeWithContext(context.Background())
}

func (p *Process) CreateTimeWithContext(ctx context.Context) (int64, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Cwd() (string, error) {
	return p.CwdWithContext(context.Background())
}

func (p *Process) CwdWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Parent() (*Process, error) {
	return p.ParentWithContext(context.Background())
}

func (p *Process) ParentWithContext(ctx context.Context) (*Process, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Status() (string, error) {
	return p.StatusWithContext(context.Background())
}

func (p *Process) StatusWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Uids() ([]int32, error) {
	return p.UidsWithContext(context.Background())
}

func (p *Process) UidsWithContext(ctx context.Context) ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}
func (p *Process) Gids() ([]int32, error) {
	return p.GidsWithContext(context.Background())
}

func (p *Process) GidsWithContext(ctx context.Context) ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}
func (p *Process) Terminal() (string, error) {
	return p.TerminalWithContext(context.Background())
}

func (p *Process) TerminalWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Nice() (int32, error) {
	return p.NiceWithContext(context.Background())
}

func (p *Process) NiceWithContext(ctx context.Context) (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) IOnice() (int32, error) {
	return p.IOniceWithContext(context.Background())
}

func (p *Process) IOniceWithContext(ctx context.Context) (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Rlimit() ([]RlimitStat, error) {
	return p.RlimitWithContext(context.Background())
}

func (p *Process) RlimitWithContext(ctx context.Context) ([]RlimitStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) RlimitUsage(gatherUsed bool) ([]RlimitStat, error) {
	return p.RlimitUsageWithContext(context.Background(), gatherUsed)
}

func (p *Process) RlimitUsageWithContext(ctx context.Context, gatherUsed bool) ([]RlimitStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) IOCounters() (*IOCountersStat, error) {
	return p.IOCountersWithContext(context.Background())
}

func (p *Process) IOCountersWithContext(ctx context.Context) (*IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) NumCtxSwitches() (*NumCtxSwitchesStat, error) {
	return p.NumCtxSwitchesWithContext(context.Background())
}

func (p *Process) NumCtxSwitchesWithContext(ctx context.Context) (*NumCtxSwitchesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) NumFDs() (int32, error) {
	return p.NumFDsWithContext(context.Background())
}

func (p *Process) NumFDsWithContext(ctx context.Context) (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) NumThreads() (int32, error) {
	return p.NumThreadsWithContext(context.Background())
}

func (p *Process) NumThreadsWithContext(ctx context.Context) (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Threads() (map[int32]*cpu.TimesStat, error) {
	return p.ThreadsWithContext(context.Background())
}

func (p *Process) ThreadsWithContext(ctx context.Context) (map[int32]*cpu.TimesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Times() (*cpu.TimesStat, error) {
	return p.TimesWithContext(context.Background())
}

func (p *Process) TimesWithContext(ctx context.Context) (*cpu.TimesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) CPUAffinity() ([]int32, error) {
	return p.CPUAffinityWithContext(context.Background())
}

func (p *Process) CPUAffinityWithContext(ctx context.Context) ([]int32, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) MemoryInfo() (*MemoryInfoStat, error) {
	return p.MemoryInfoWithContext(context.Background())
}

func (p *Process) MemoryInfoWithContext(ctx context.Context) (*MemoryInfoStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) MemoryInfoEx() (*MemoryInfoExStat, error) {
	return p.MemoryInfoExWithContext(context.Background())
}

func (p *Process) MemoryInfoExWithContext(ctx context.Context) (*MemoryInfoExStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Children() ([]*Process, error) {
	return p.ChildrenWithContext(context.Background())
}

func (p *Process) ChildrenWithContext(ctx context.Context) ([]*Process, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) OpenFiles() ([]OpenFilesStat, error) {
	return p.OpenFilesWithContext(context.Background())
}

func (p *Process) OpenFilesWithContext(ctx context.Context) ([]OpenFilesStat, error) {
	return []OpenFilesStat{}, common.ErrNotImplementedError
}
func (p *Process) Connections() ([]net.ConnectionStat, error) {
	return p.ConnectionsWithContext(context.Background())
}

func (p *Process) ConnectionsWithContext(ctx context.Context) ([]net.ConnectionStat, error) {
	return []net.ConnectionStat{}, common.ErrNotImplementedError
}
func (p *Process) NetIOCounters(pernic bool) ([]net.IOCountersStat, error) {
	return p.NetIOCountersWithContext(context.Background(), pernic)
}

func (p *Process) NetIOCountersWithContext(ctx context.Context, pernic bool) ([]net.IOCountersStat, error) {
	return []net.IOCountersStat{}, common.ErrNotImplementedError
}
func (p *Process) IsRunning() (bool, error) {
	return p.IsRunningWithContext(context.Background())
}

func (p *Process) IsRunningWithContext(ctx context.Context) (bool, error) {
	return true, common.ErrNotImplementedError
}
func (p *Process) MemoryMaps(grouped bool) (*[]MemoryMapsStat, error) {
	return p.MemoryMapsWithContext(context.Background(), grouped)
}

func (p *Process) MemoryMapsWithContext(ctx context.Context, grouped bool) (*[]MemoryMapsStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) SendSignal(sig syscall.Signal) error {
	return p.SendSignalWithContext(context.Background(), sig)
}

func (p *Process) SendSignalWithContext(ctx context.Context, sig syscall.Signal) error {
	return common.ErrNotImplementedError
}
func (p *Process) Suspend() error {
	return p.SuspendWithContext(context.Background())
}

func (p *Process) SuspendWithContext(ctx context.Context) error {
	return common.ErrNotImplementedError
}
func (p *Process) Resume() error {
	return p.ResumeWithContext(context.Background())
}

func (p *Process) ResumeWithContext(ctx context.Context) error {
	return common.ErrNotImplementedError
}
func (p *Process) Terminate() error {
	return p.TerminateWithContext(context.Background())
}

func (p *Process) TerminateWithContext(ctx context.Context) error {
	return common.ErrNotImplementedError
}
func (p *Process) Kill() error {
	return p.KillWithContext(context.Background())
}

func (p *Process) KillWithContext(ctx context.Context) error {
	return common.ErrNotImplementedError
}
func (p *Process) Username() (string, error) {
	return p.UsernameWithContext(context.Background())
}

func (p *Process) UsernameWithContext(ctx context.Context) (string, error) {
	return "", common.ErrNotImplementedError
}
