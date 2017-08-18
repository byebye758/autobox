package command

import (
	"fmt"
	"os/exec"
)

func Kubectlapply(kubectlpath, kubeconfigpath, json string) {
	//args := "KUBECONFIG=" + kubeconfigpath + "&& echo  " + json + " | " + kubectlpath + " apply -f -"
	//args := "echo $KUBECONFIG && echo  " + json + " | " + kubectlpath + " apply -f -"
	args := "echo $KUBECONFIG "

	cmd := exec.Command("/bin/sh", "-c", args)
	env := make([]string, 1)

	env[0] = "KUBECONFIG=" + kubeconfigpath
	fmt.Println(env)
	a, err := cmd.CombinedOutput()
	fmt.Println(string(a), err)

}
