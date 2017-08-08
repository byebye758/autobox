package command

import (
	"fmt"
	"os/exec"
)

func Kubectlapply(kubectlpath, kubeconfigpath, json string) {
	args := "KUBECONFIG=" + kubeconfigpath + "&& echo  " + json + " | " + kubectlpath + " apply -f -"
	cmd := exec.Command("/bin/sh", "-c", args)
	a, err := cmd.CombinedOutput()
	fmt.Println(string(a), err)

}
