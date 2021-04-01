package runner

import "GoBackupscan/pkg/options"

type Runner struct {
	Options *options.Options
}


func NewRunner(options *options.Options)(*Runner){
	runner := &Runner{
		Options: options,
	}
	return runner
}
