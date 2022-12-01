package main

import "os/exec"

//step define a pipeline step; has step name;
//exe represent executable name of the external tool we want to execute;
//args contains the arguments for the executable; message
//which is the output message in case of success;
//proj represent the target project on which to execute the task
type step struct {
	name    string
	exe     string
	args    []string
	message string
	proj    string
}

//newStep instantiates and returns a new step
func newStep(name, exe, message, proj string, args []string) step {
	return step{
		name:    name,
		exe:     exe,
		args:    args,
		message: message,
		proj:    proj,
	}
}

//execute the external program
func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}
	return s.message, nil
}
