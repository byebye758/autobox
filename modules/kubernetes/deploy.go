package kubernetes

import (
	//"fmt"
	//"k8s.io/api/apps/v1"
	"encoding/json"
	"k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"strings"
	//"github.com/astaxie/beego/logs"
	"k8s.io/api/core/v1"
)

type K8sport struct {
	Name          string
	Containerport int32
	Serviceport   int32
}
type K8s struct {
	Projectname string
	Replace     int32
	Namespace   string
	Image       string
	//Cmd         []string
	Port []K8sport
}

func (k *K8s) Deployjson() ([]byte, error) {

	matedata := &metav1.ObjectMeta{
		Name:      k.Projectname,
		Namespace: k.Namespace,
		Labels: map[string]string{
			"app": k.Projectname,
		},
	}

	typemate := &metav1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "extensions/v1beta1",
	}

	containerport := new(v1.ContainerPort)
	containerports := make([]v1.ContainerPort, 0)
	for _, v := range k.Port {
		containerport.Name = v.Name
		containerport.ContainerPort = v.Containerport
		containerports = append(containerports, *containerport)
	}

	container := &v1.Container{
		Command: []string{},
		Image:   k.Image,
		Name:    k.Projectname,
		Ports:   containerports,
	}
	containers := make([]v1.Container, 0)
	containers = append(containers, *container)
	template := &v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": k.Projectname,
			},
		},
		Spec: v1.PodSpec{
			Containers: containers,
		},
	}

	sp := &v1beta1.DeploymentSpec{
		Replicas: &k.Replace,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": k.Projectname,
			},
		},
		Template: *template,
	}

	deploy := &v1beta1.Deployment{
		TypeMeta:   *typemate,
		ObjectMeta: *matedata,
		Spec:       *sp,
	}

	data, err := json.Marshal(deploy)

	return data, err
}
