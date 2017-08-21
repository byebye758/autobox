package kubernetes

import (
	"encoding/json"
	"k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8s) AutoscalToJson() ([]byte, error) {

	matedata := &metav1.ObjectMeta{
		Name:      k.ProjectName,
		Namespace: k.NameSpace,
		Labels: map[string]string{
			"app": k.ProjectName,
		},
	}
	typemate := &metav1.TypeMeta{
		Kind:       "HorizontalPodAutoscaler",
		APIVersion: "autoscaling/v1",
	}

	spec := &v1.HorizontalPodAutoscalerSpec{
		ScaleTargetRef: v1.CrossVersionObjectReference{
			Kind:       "Deployment",
			Name:       k.ProjectName,
			APIVersion: "extensions/v1beta1",
		},
		MinReplicas:                    &k.K8sAutoScal.Min,
		MaxReplicas:                    k.K8sAutoScal.Max,
		TargetCPUUtilizationPercentage: &k.K8sAutoScal.Cpuload,
	}

	autoscal := v1.HorizontalPodAutoscaler{
		ObjectMeta: *matedata,
		TypeMeta:   *typemate,
		Spec:       *spec,
	}

	data, err := json.Marshal(autoscal)

	return data, err
}
