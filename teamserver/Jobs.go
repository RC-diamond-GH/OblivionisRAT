package main

const (
	SHELL   = 0x10000001
	COMMAND = 0x10000002
)

const (
	DOWNLOAD = 0x20000001
)

type Job struct {
	job_type uint8
	command  uint8
	shell    string
}

func is_jobs_null(listener *Listener, i int) bool {
	if len(listener.Beacons[i].jobs) == 0 {
		return false
	} else {
		return true
	}
}
