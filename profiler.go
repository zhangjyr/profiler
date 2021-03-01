package profiler

import (
	"time"
)

type Profile interface {
	GetStart() time.Time
	GetEnd() time.Time
	GetDuration() time.Duration
	IsReady() bool
	String() string
	Parent() Profile
	Children() []Profile
}

type Profiler interface {
	Start(string) Profile
	End(Profile) Profile
	Close()
}

type ReplayProfiler interface {
	StartAt(string, time.Time) Profile
	EndAt(Profile, time.Time) Profile
	Close()
}

type BaseProfile struct {
	parent   Profile
	children []Profile
}

func (p *BaseProfile) Parent() Profile {
	return p.parent
}

func (p *BaseProfile) Children() []Profile {
	return p.children
}
