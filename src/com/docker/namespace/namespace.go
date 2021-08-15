package namespace

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

/**
configuration配置
run kind选择directory
directory选择G:\workspace\golang\learn\src\com\lyh 即main文件的目录
out directory 选择G:\workspace\golang\learn\bin
working directory选择G:\workspace\golang\learn\src\com\lyh 即main文件的目录
environments GOOS=windows;   linux GOOS=linux;GOARCH=amd64
不能识别linux的配置时，setting-->GO-->vendor & building tags修改 OS 选择为linux
 */
func Command(){
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,

	}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid: uint32(1), Gid: uint32(1)}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil{
		log.Fatal(err)
	}
}
