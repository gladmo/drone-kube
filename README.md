# drone-kube
drone-kube is drone k8s continuous delivery plugin.

## Usage
deliver one
```yaml
...
- name: delivery
  image: gladmo/drone-kube:1
  settings:
    master_url: https://192.168.0.1:6443 # you k8s master url
    resource: deployment                 # what resource you want delivery
    namespace: default                   # delivery namespace
    name: nginx                          # resource name
    container_name: nginx                # need update image container_name, default: deploy.Spec.Template.Spec.Containers[0]
...
```

deliver multiple
```yaml
...
- name: delivery
  image: gladmo/drone-kube:1
  settings:
    master_url: https://192.168.0.1:6443
    namespace: default
    multiple:
      - resource: deployment
        name: nginx
        container_name: nginx
      - resource: cronjob
        name: publish-github
        container_name: publish-github
...
```