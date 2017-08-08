package main

import (
	"autobox/modules/command"
	"autobox/modules/kubernetes"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strconv"
	"strings"
)

var (
	app            = kingpin.New("autobox", "auto build k8s cd for k8s api")
	kubectlpath    = app.Flag("kubectlpath", "Set kubectl Path").Default("/opt/kubernetes/bin/kubectl").String()
	kubeconfigpath = app.Flag("kubeconfigpath", "Set kubeconfig Path").Default("/etc/kubernetes/admin.conf").String()
	projectname    = app.Flag("Projectname", "Set  App  Project  Name. The Jenkins project name is recommended").Default("JOB_NAME").String()
	namespace      = app.Flag("namespace", "").Default("default").String()

	replace = app.Flag("replace", "Set kubernetes pod nubmber").Default("1").Int32()
	image   = app.Flag("image", "Set docker image name").Default("IMAGE").String()
	//serviceport   = app.Flag("serviceport", "Set up the kubernetes internal service port").Default("0").Int32()
	//containerport = app.Flag("containerport", "Set up the kubernetes internal container port").Default("0").Int32()
	cmd  = app.Flag("cmd", "container exec cmd").Default("abc:bcd,aaa:bbb").Strings()
	env  = app.Flag("env", "container exec cmd").Default("abc,bcd,aaa,bbb").Strings()
	port = app.Flag("port", "").Default("21:21,22:22").String()
)

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))
	//fmt.Println(*kubectlpath, *cmd, *port)
	pp := make([]kubernetes.K8sport, 0)
	p := strings.SplitN(*port, ",", -1)
	kport := new(kubernetes.K8sport)
	for k, v := range p {
		v := (strings.SplitN(v, ":", -1))
		kport.Name = "port" + string(k+100)

		kport.Serviceport = stoint32(v[0])
		kport.Containerport = stoint32(v[1])
		pp = append(pp, *kport)

	}
	aa := kubernetes.K8s{
		Projectname: *projectname,
		Replace:     *replace,
		Namespace:   *namespace,
		Image:       *image,
		Port:        pp,
	}
	//fmt.Println(aa)
	d, _ := aa.Deployjson()
	fmt.Println(string(d))
	command.Kubectlapply(*kubectlpath, *kubeconfigpath, string(d))

}
func stoint32(s string) (i int32) {
	ii, _ := strconv.ParseInt(s, 10, 64)
	i = int32(ii)
	return i

}
