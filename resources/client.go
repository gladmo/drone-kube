package resources

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// UpdateInfo 更新 deploy 所需信息
type UpdateInfo struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`

	ContainerName string `json:"container_name"` // 为空修改第一个容器
	Image         string `json:"update_image"`   // 为空根据版本修改
	ImageVersion  string `json:"image_version"`  // 与image二选一
}

// UpdateResult 更新结果
type UpdateResult struct {
	OldImage     string
	CurrentImage string
}

// getClientSet 获取 k8s client
func getClientSet() (client *kubernetes.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return
	}

	return kubernetes.NewForConfig(config)
}
