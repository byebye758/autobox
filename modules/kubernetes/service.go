package kubernetes

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

func (k *K8s) Servicejson() ([]byte, error) {

	typemeta := &metav1.TypeMeta{
		Kind:       k.Projectname,
		APIVersion: "v1",
	}
	matedata := &metav1.ObjectMeta{
		Name:      k.Projectname,
		Namespace: k.Namespace,
		Labels: map[string]string{
			"app": k.Projectname,
		},
	}
	serviceport := new(v1.ServicePort)
	serviceports := make([]v1.ServicePort, 0)

	for k1, v := range k.Port {

		serviceport.Name = "port" + strconv.FormatInt(int64(k1), 10)
		serviceport.Port = v.Containerport
		serviceport.TargetPort = intstr.IntOrString{
			IntVal: v.Serviceport,
		}
		serviceports = append(serviceports, *serviceport)
	}
	spec := &v1.ServiceSpec{
		Ports: serviceports,
		Selector: map[string]string{
			"app": k.Projectname,
		},
	}

	service := v1.Service{
		TypeMeta:   *typemeta,
		ObjectMeta: *matedata,
		Spec:       *spec,
	}

	data, err := json.Marshal(service)
	return data, err
}
