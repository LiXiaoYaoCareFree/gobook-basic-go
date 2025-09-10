```bash
$$env:GOOS = "linux"
```
```bash
$$env:GOARCH = "arm"
```
```bash
go build -tags=k8s -o webook .
```
```bash
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
```


## 构建和推送Docker镜像
项目中的Makefile已经提供了构建Docker镜像的命令：
```bash
make docker
```
## 按顺序部署K8s资源
按照依赖关系，应该按以下顺序部署：

### 1. 部署存储资源
```bash
kubectl apply -f webook-record-mysql-pv.yaml
kubectl apply -f webook-record-mysql-pvc.yaml
```
### 2. 部署MySQL和Redis
```bash
kubectl apply -f webook-record-mysql-deployment.yaml
kubectl apply -f webook-record-mysql-service.yaml
kubectl apply -f webook-record-redis-deployment.yaml
kubectl apply -f webook-record-redis-service.yaml
```
### 3. 部署应用服务
```bash
kubectl apply -f webook-deployment.yaml
kubectl apply -f webook-service.yaml
```
### 4. 部署Ingress配置
```bash
kubectl apply -f webook-record-ingress.yaml
```
## 验证部署
部署完成后，可以使用以下命令检查各个资源的状态：
```bash
kubectl get pods       # 查看所有Pod状态
kubectl get services   # 查看所有Service
kubectl get ingress    # 查看Ingress配置
```

## 访问应用
根据Ingress配置，应用可以通过以下方式访问：

- 如果在本地Kubernetes集群上运行：在浏览器中访问 http://localhost/
- 需确保hosts  文件中有 127.0.0.1 localhost 的映射

应用使用了持久卷存储MySQL数据，确保主机上存在 /mnt/data 目录
如果需要删除所有资源，可以按照相反的顺序使用 kubectl delete -f <文件名> 命令。

```bash
kubectl get all 
kubectl delete pv webook-mysql-pvc 
kubectl delete pvc webook-mysql-pvc 
kubectl delete ingress webook-record-ingress 
kubectl delete service webook-record webook-record-mysql webook-record-redis 
kubectl delete deployment webook-record-service webook-record-mysql webook-record-redis
```