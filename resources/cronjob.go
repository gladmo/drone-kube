package resources

import (
	"context"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

// UpdateCronJobsImage 更新 cronjob 镜像
func (th UpdateInfo) UpdateCronJobsImage() (ur UpdateResult, err error) {
	clientSet, err := getClientSet()
	if err != nil {
		return
	}

	cronJobs := clientSet.BatchV1beta1().CronJobs(th.Namespace)
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		job, err := cronJobs.Get(context.Background(), th.Name, metaV1.GetOptions{})
		if err != nil {
			return err
		}

		containers := job.Spec.JobTemplate.Spec.Template.Spec.Containers
		ur, containers = th.UpdateImage(containers)
		job.Spec.JobTemplate.Spec.Template.Spec.Containers = containers
		_, err = cronJobs.Update(
			context.TODO(), job, metaV1.UpdateOptions{})
		return err
	})

	if err != nil {
		return
	}

	return
}
