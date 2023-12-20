package globals

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var Client map[string]*kubernetes.Clientset
var TestClient *kubernetes.Clientset

func InitClient() {
	Client = make(map[string]*kubernetes.Clientset)
	kubeconfigPath := "/root/.kube/config"
	
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		// kubeconfig 파일이 없는 경우 InClusterConfig 사용
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.Errorln(err)
		}
	}

	// Kubernetes 클라이언트 생성
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Errorln(err)
	}

	Client["1"] = clientset
	TestClient = clientset
}
