package main

import (
	//"autobox/modules/command"
	"autobox/modules/kubernetes"
	"autobox/modules/tools"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strconv"
	//"strings"
	"errors"
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
	//cmd  = app.Flag("cmd", "container exec cmd").Default("abc:bcd,aaa:bbb").Strings()
	//env  = app.Flag("env", "container exec cmd").Default("abc,bcd,aaa,bbb").Strings()
	port = app.Flag("port", "21:21,22:22").Default("0:0,0:0").String()
	http = app.Flag("http", `url 方式发布您的应用 eg：--http="url=http://www.baidu.com,serverport=21,https=false,cafile=$cafilepath,keyfile=$keyfilepath,path=$Path";--http="url=https://www.163.com,serverport=22,https=true,cafile=$cafilepath,keyfile=$keyfilepath,path=$Path"`).Default(`url=http://www.baidu.com,serviceport=0,https=false,cafile=./cafilepath,keyfile=./keyfilepath`, `url=https://www.163.com,servicerport=0,https=true,cafile=./cafilepath1,keyfile=./keyfilepath1`).Strings()
)

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))
	aa, _ := tools.ArgToToolsStruct(*kubectlpath, *kubeconfigpath, *projectname, *namespace, *image, *port, *replace, *http)

	d, _ := aa.DeployToJson()
	s, _ := aa.Servicejson()

	sjson, _ := aa.StoJsons()
	ingjson, _ := aa.IngressToJson()

	//fmt.Println(aa)
	//fmt.Println(string(s))
	// fmt.Println(strconv.Quote(string(d)))

	fmt.Println(string(d), string(s), string(ingjson))

	// command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(d)))
	// command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(s)))
	for _, v := range sjson {
		//command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(v)))
		fmt.Println(string(v))
	}
	// command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(ingjson)))

	if aa.Image != "IMAGE" {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(d)))

	}
	err := servicecmd(aa.Port)
	if err == nil {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(s)))
	}

	err = secretcmd(aa.Ingress)
	if err == nil {
		for _, v := range sjson {
			command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(v)))
		}
	}

	err = ingresscmd(aa.Ingress)
	if err == nil {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(ingjson)))

	}
}

func stoint32(s string) (i int32) {
	ii, _ := strconv.ParseInt(s, 10, 64)
	i = int32(ii)
	return i

}

func servicecmd(ports []kubernetes.K8sport) error {
	for _, v := range ports {
		if v.ContainerPort == 0 {
			return errors.New("no service port")

		}
	}
	return nil
}

func ingresscmd(ingress []kubernetes.K8sIngress) error {
	for _, v := range ingress {
		if v.ServicePort == 0 {
			return errors.New("no ingress port")

		}
	}
	return nil
}

func secretcmd(ingress []kubernetes.K8sIngress) error {
	for _, v := range ingress {
		if v.Secret.CafilePath == `./cafilepath` {
			return errors.New("no secretcafile ")
		}
	}
	return nil
}
