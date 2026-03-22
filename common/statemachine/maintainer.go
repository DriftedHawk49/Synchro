/*
This is a state maintainer which saves the state of machine,
keeps track of current state, avoids any other operation acceptance until current operation is done
*/

package statemachine

import (
	"sync"
	"time"
)

type State struct {
	ServiceName  string
	State        string
	JobId        string
	JobStartedAt time.Time
	m            sync.RWMutex
}

func New(svcName string) *State {
	return &State{
		ServiceName: svcName,
		m:           sync.RWMutex{},
	}
}

func (s *State) Set(jobId string, state string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.State = state
	s.JobId = jobId
	s.JobStartedAt = time.Now()
}

// Will change state to IDLE
func (s *State) UnSet() {
	s.m.Lock()
	defer s.m.Unlock()
	s.State = "IDLE"
	s.JobId = ""
	s.JobStartedAt = time.Now()

}

func (s *State) Busy() bool {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.State != "IDLE"

}

// Returns State, JobId & JobStartedAt
func (s *State) CurrentState() (string, string, time.Time) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.State, s.JobId, s.JobStartedAt
}
