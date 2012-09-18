package gpu

import (
	"github.com/barnex/cuda4/cu"
	"github.com/barnex/cuda4/safe"
	"nimble-cube/core"
)

type Uploader struct {
	hostdata []float32
	rlock    *core.RMutex
	devdata  safe.Float32s
	wlock    *core.RWMutex
	bsize    int
	stream   cu.Stream
}

func NewUploader(hostdata []float32, hostlock *core.RMutex, devdata safe.Float32s, devlock *core.RWMutex) *Uploader {
	u := new(Uploader)
	u.hostdata = hostdata
	u.rlock = hostlock
	u.devdata = devdata
	u.wlock = devlock
	u.bsize = 16 // TODO !! Always lock max
	return u
}

func (u *Uploader) Run() {
	core.Debug("uploader: run")
	LockCudaThread()
	u.stream = cu.StreamCreate()
	MemHostRegister(u.hostdata)
	bsize := u.bsize

	for {
		for i := 0; i < len(u.hostdata); i += bsize {
			j := i + bsize
			u.rlock.RLock(i, i)
			u.wlock.WLock(i, i)
			u.rlock.RLock(i, j)
			u.wlock.WLock(i, j)
			core.Debug("upload", i, j)
			u.devdata.CopyHtoDAsync(u.hostdata, u.stream)
			u.stream.Synchronize()
			u.rlock.RLock(j, j)
			u.wlock.WLock(j, j)
		}
	}
}

// _____________________________________

type Downloader struct {
	devdata  safe.Float32s
	rlock    *core.RMutex
	hostdata []float32
	wlock    *core.RWMutex
	bsize    int
	stream   cu.Stream
}

func NewDownloader(devdata safe.Float32s, devlock *core.RMutex, hostdata []float32, hostlock *core.RWMutex) *Downloader {
	u := new(Downloader)
	u.devdata = devdata
	u.rlock = devlock
	u.hostdata = hostdata
	u.wlock = hostlock
	u.bsize = 16 // TODO !! Always lock max
	return u
}

func (u *Downloader) Run() {
	core.Debug("downloader: run")
	LockCudaThread()
	u.stream = cu.StreamCreate()
	MemHostRegister(u.hostdata)
	bsize := u.bsize

	for {
		for i := 0; i < len(u.hostdata); i += bsize {
			j := i + bsize
			u.rlock.RLock(i, i)
			u.wlock.WLock(i, i)
			u.wlock.WLock(i, j)
			u.rlock.RLock(i, j)
			core.Debug("download", i, j)
			u.devdata.CopyDtoHAsync(u.hostdata, u.stream)
			u.stream.Synchronize()
			u.rlock.RLock(j, j)
			u.wlock.WLock(j, j)
		}
	}
}
