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
	serviceName  string
	state        string
	jobId        string
	idleState    string
	jobStartedAt time.Time
	m            sync.RWMutex
}

// Needs idle state to compare
func New(svcName string, idleState string) *State {
	return &State{
		serviceName: svcName,
		m:           sync.RWMutex{},
		idleState:   idleState,
	}
}

func (s *State) Set(jobId string, state string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.state = state
	s.jobId = jobId
	s.jobStartedAt = time.Now()
}

// Will change state to IDLE
func (s *State) UnSet() {
	s.m.Lock()
	defer s.m.Unlock()
	s.state = s.idleState
	s.jobId = ""
	s.jobStartedAt = time.Now()

}

func (s *State) Busy() bool {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.state != s.idleState

}

// Returns Servicename, State, JobId & JobStartedAt
func (s *State) CurrentState() (string, string, string, time.Time) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.serviceName, s.state, s.jobId, s.jobStartedAt
}
