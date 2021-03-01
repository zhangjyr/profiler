package profiler

import (
  "time"
)

type TimeProfile struct {
  BaseProfile
  start    time.Time
  end      time.Time
}

func (p *TimeProfile) GetStart() time.Time {
  return p.start
}

func (p *TimeProfile) GetEnd() time.Time {
  return p.end
}

func (p *TimeProfile) GetDuration() time.Duration {
  return p.end.Sub(p.start)
}

func (p *TimeProfile) IsReady() bool {
  return p.end != time.Time{}
}

func (p *TimeProfile) String() string {
  return p.GetDuration().String()
}


type TimeProfiler struct {

}

func (p *TimeProfiler) Start(name string) Profile {
  return &TimeProfile{ start: time.Now() }
}

func (p *TimeProfiler) End(profile Profile) Profile {
  tp := profile.(*TimeProfile)
  tp.end = time.Now()
  return tp
}

func (p *TimeProfiler) Close() {

}
