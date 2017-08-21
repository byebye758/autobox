package tools

import (
	"autobox/modules/kubernetes"
	//"fmt"
	//"gopkg.in/alecthomas/kingpin.v2"
	//"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ArgToToolsStruct(kubectlpath, kubeconfigpath, projectname, namespace, image, port string, replace int32, http []string, autoscal string) (k8s kubernetes.K8s, err error) {

	pp := make([]kubernetes.K8sport, 0)
	p := strings.SplitN(port, ",", -1)
	kport := new(kubernetes.K8sport)
	for k, v := range p {
		v := (strings.SplitN(v, ":", -1))
		kport.Name = "port" + strconv.FormatInt(int64(k), 10)

		kport.ServicePort = stoint32(v[0])
		kport.ContainerPort = stoint32(v[1])
		pp = append(pp, *kport)

	}

	ingresss, _ := HttpParser(http, projectname, namespace)
	auto, _ := AutoscalParser(autoscal)
	//fmt.Println(auto)
	//fmt.Println(ingresss, err)
	k8s = kubernetes.K8s{
		ProjectName: projectname,
		NameSpace:   namespace,
		Image:       image,
		Port:        pp,
		Ingress:     ingresss,
		Replace:     replace,
		K8sAutoScal: auto,
	}
	return k8s, nil

}

func HttpParser(http []string, projectname, namespace string) (ingresss []kubernetes.K8sIngress, err error) {

	secret := new(kubernetes.IngressSecret)
	//secrets := make([]kubernetes.IngressSecret, 0)
	ingress := new(kubernetes.K8sIngress)
	for _, v := range http {
		p := strings.SplitN(v, ",", -1)
		//fmt.Println(p)
		ddd := make(map[string][]byte)
		for k, v := range p {
			//fmt.Println(v)
			pp := strings.SplitN(v, "=", -1)
			//fmt.Println(pp, pp[0])

			switch pp[0] {
			case "url":
				ingress.Host = pp[1]
			case "https":
				if strings.EqualFold(pp[1], "true") {
					ingress.Https = true
				} else {
					ingress.Https = false
				}
			case "cafile":
				secret.CafilePath = pp[1]

				_, d1 := Read3(secret.CafilePath)

				ddd["tls.crt"] = d1

				secret.Name = projectname + strconv.FormatInt(int64(k), 10)
				secret.NameSpace = namespace

			case "keyfile":
				secret.KeyfilePath = pp[1]
				_, d2 := Read3(secret.KeyfilePath)
				ddd["tls.key"] = d2
				//secret.Data
			case "path":
				ingress.Path = pp[1]
			case "serviceport":
				ingress.ServicePort = stoint32(pp[1])

			}
			//secrets = append(secrets, *secret)

		}
		secret.Data = ddd
		ingress.Secret = *secret
		ingresss = append(ingresss, *ingress)

	}
	return ingresss, nil

}

func AutoscalParser(autoscal string) (auto kubernetes.K8sAutoScal, err error) {
	p := strings.SplitN(autoscal, ",", -1)
	for _, v := range p {

		pp := strings.SplitN(v, "=", -1)
		switch pp[0] {
		case "min":
			auto.Min = stoint32(pp[1])
		case "max":
			auto.Max = stoint32(pp[1])
		case "cpuload":
			auto.Cpuload = stoint32(pp[1])

		}
	}
	err = nil
	//fmt.Println(auto)
	return auto, err
}

func stoint32(s string) (i int32) {
	ii, _ := strconv.ParseInt(s, 10, 64)
	i = int32(ii)
	return i

}

func Read3(path string) (string, []byte) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)

	return string(fd), fd

}
