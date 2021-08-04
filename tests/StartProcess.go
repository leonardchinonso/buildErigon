package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

		if p, err := os.StartProcess(args[0], args, &procAttr); err == nil {
			return p, nil
		}

		fmt.Print("error is", err)
	}

	return nil, err
}

func main() {
	//fmt.Println(os.Environ())
	//fmt.Println(os.Getenv("PATH"))
	//if proc, err := Start(`C:\Users\USER\AppData\Local\Microsoft\WindowsApps\PythonSoftwareFoundation.Python.3.8_qbz5n2kfra8p0\python.exe`,
	//	`testing.py`); err == nil {
	//	proc.Wait()
	//} else {
	//	fmt.Println("err is: ", err)
	//}
	//if proc, err := Start(`C:\Users\USER\AppData\Local\Microsoft\WindowsApps\python3.exe`,
	//	`C:\Users\USER\Documents\Learning\testingGo\testing.py`); err == nil {
	//	proc.Wait()
	//} else {
	//	fmt.Println("err is: ", err)
	//}

	// With StartProcess itself
	//args := []string {"C:\\Users\\USER\\AppData\\Local\\Microsoft\\WindowsApps\\PythonSoftwareFoundation.Python.3.8_qbz5n2kfra8p0\\python.exe",
	//	"testing.py"}

	//args := []string {"C:\\Program Files\\Git\\cmd\\git.exe", "commit", "-m", "testing commit"}

	//args := []string {"C:\\Program Files\\Git\\cmd\\git.exe", "push", "--set-upstream", "https://github.com/leonardchinonso/testingGoWithGit.git", "master"}

	//args := []string {"C:\\Program Files\\Git\\cmd\\git.exe", "branch"}

	//args := []string {"C:\\Program Files\\Git\\cmd\\git.exe", "checkout", "-b", "testing-check-new-b"}

	args := []string{"C:\\Program Files\\Git\\cmd\\git.exe", "rev-parse", "--verify", "develop"}

	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	p, err := os.StartProcess(args[0], args, &procAttr)
	if err != nil {
		fmt.Println("Error is: ", err)
	} else {
		fmt.Println("P is: ", p)
	}
}

//args := []string{"C:\\Users\\USER\\AppData\\Local\\Microsoft\\WindowsApps\\python.exe", "C:\\Users\\USER\\Documents\\Learning\\testingGo\\testing.py"}
