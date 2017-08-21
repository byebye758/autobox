package main

import (
	"autobox/modules/command"
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
	namespace      = app.Flag("namespace", `eg: --namespace="default"`).Default("default").String()

	replace  = app.Flag("replace", "Set kubernetes pod nubmber").Default("0").Int32()
	image    = app.Flag("image", "Set docker image name").Default("IMAGE").String()
	autoscal = app.Flag("autoscal", `Set autoscal  eg:  "autoscal=min=1,max=20,cpuload=1"  cpuload 1-100`).Default("min=1,max=20,cpuload=0").Strings()
	//serviceport   = app.Flag("serviceport", "Set up the kubernetes internal service port").Default("0").Int32()
	//containerport = app.Flag("containerport", "Set up the kubernetes internal container port").Default("0").Int32()
	//cmd  = app.Flag("cmd", "container exec cmd").Default("abc:bcd,aaa:bbb").Strings()
	//env  = app.Flag("env", "container exec cmd").Default("abc,bcd,aaa,bbb").Strings()
	port = app.Flag("port", `eg: --port="21:21,22:22"`).Default("0:0,0:0").String()
	http = app.Flag("http", `url 方式发布您的应用 eg：--http="url=http://www.baidu.com,serverport=21,https=false,cafile=$cafilepath,keyfile=$keyfilepath,path=$Path";--http="url=https://www.163.com,serverport=22,https=true,cafile=$cafilepath,keyfile=$keyfilepath,path=$Path"`).Default(`url=http://www.baidu.com,serviceport=0,https=false,cafile=./cafilepath,keyfile=./keyfilepath`, `url=https://www.163.com,servicerport=0,https=true,cafile=./cafilepath1,keyfile=./keyfilepath1`).Strings()
)

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))
	aa, _ := tools.ArgToToolsStruct(*kubectlpath, *kubeconfigpath, *projectname, *namespace, *image, *port, *replace, *http, *autoscal)
	d, _ := aa.DeployToJson()
	s, _ := aa.Servicejson()

	sjson, _ := aa.StoJsons()
	ingjson, _ := aa.IngressToJson()
	auto, _ := aa.AutoscalToJson()
	fmt.Println(aa)
	//fmt.Println(aa)
	//fmt.Println(string(s))
	// fmt.Println(strconv.Quote(string(d)))

	if aa.Image != "IMAGE" {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(d)))

	}
	err := servicecmd(aa.Port)
	if err == nil {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(s)))
	}

	err = secretcmd(aa.Ingress)
	fmt.Println(err)
	if err == nil {
		for _, v := range sjson {
			command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(v)))
		}
	}

	err = ingresscmd(aa.Ingress)
	fmt.Println(err)
	if err == nil {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(ingjson)))

	}
	err = autoscalcmd(aa.K8sAutoScal)
	if err == nil {
		command.Kubectlapply(*kubectlpath, *kubeconfigpath, strconv.Quote(string(auto)))
		fmt.Println(string(auto))
	}
	fmt.Println(string(auto))
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

func autoscalcmd(auto kubernetes.K8sAutoScal) error {
	if auto.Cpuload == 0 {
		return errors.New("no autoscal ")
	}
	return nil
}
