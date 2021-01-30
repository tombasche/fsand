package process

import "os/exec"

func Execute(script string) error {
	cmd := exec.Command(script)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil

}
