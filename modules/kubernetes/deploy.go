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
	"k8s.io/apimachinery/pkg/api/resource"
)

func (k *K8s) DeployToJson() ([]byte, error) {

	matedata := &metav1.ObjectMeta{
		Name:      k.ProjectName,
		Namespace: k.NameSpace,
		Labels: map[string]string{
			"app": k.ProjectName,
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
		containerport.ContainerPort = v.ContainerPort
		containerports = append(containerports, *containerport)
	}
	var cpu v1.ResourceName
	cpu = "cpu"
	cpuload := resource.Quantity{
		Format: "400m",
	}

	container := &v1.Container{
		Command: []string{},
		Image:   k.Image,
		Name:    k.ProjectName,
		Ports:   containerports,
		Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				cpu: cpuload,
			},
		},
	}
	containers := make([]v1.Container, 0)
	containers = append(containers, *container)
	template := &v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": k.ProjectName,
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
				"app": k.ProjectName,
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
