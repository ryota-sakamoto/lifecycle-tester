package handler

type State struct {
	IsFailedHealthz bool `json:"is_failed_healthz"`
}

var state State
