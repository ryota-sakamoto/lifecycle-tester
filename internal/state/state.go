package state

import (
	"sync"
)

type State struct {
	IsFailedHealthz      bool  `json:"is_failed_healthz"`
	ShutdownDelaySeconds int64 `json:"shutdown_delay_seconds"`
}

type StateManager struct {
	state State
	mu    sync.RWMutex
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) GetState() State {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.state
}

func (sm *StateManager) SetState(state State) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.state = state
}
