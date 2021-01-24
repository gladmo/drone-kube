package resources

import (
	"context"
	"strings"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

// UpdateImage 更新镜像名称
func (th UpdateInfo) UpdateImage(containers []v1.Container) (ur UpdateResult, result []v1.Container) {
	var index int
	// find container and change image
	if th.ContainerName != "" {
		for idx := range containers {
			if containers[idx].Name == th.ContainerName {
				index = idx
				break
			}
		}
	}

	ur.OldImage = containers[index].Image
	if len(th.Image) > 0 {
		containers[index].Image = th.Image
	} else {
		imageName := containers[index].Image
		versionStart := strings.Index(imageName, ":")
		if versionStart != -1 {
			imageName = imageName[:versionStart+1] + th.ImageVersion
			containers[index].Image = imageName
		}
		// 未找到版本不更新镜像信息
	}

	ur.CurrentImage = containers[index].Image
	result = containers
	return
}

// UpdateDeploymentImage 更新 deploy 镜像
func (th UpdateInfo) UpdateDeploymentImage() (ur UpdateResult, err error) {
	clientSet, err := getClientSet()
	if err != nil {
		return
	}

	deploymentsClient := clientSet.AppsV1().Deployments(th.Namespace)

	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deploy, err := deploymentsClient.Get(
			context.Background(),
			th.Name,
			metaV1.GetOptions{})
		if err != nil {
			return err
		}

		// 更新镜像
		containers := deploy.Spec.Template.Spec.Containers
		ur, containers = th.UpdateImage(containers)
		deploy.Spec.Template.Spec.Containers = containers
		_, err = deploymentsClient.Update(
			context.TODO(), deploy, metaV1.UpdateOptions{})
		return err
	})
	if err != nil {
		return
	}

	return
}
