# GKEでの高負荷テスト
```shell
REGION=asia-northeast1
ZONE=${REGION}-b
PROJECT=$(gcloud config get-value project)
CLUSTER=gke-load-test
SCOPE="https://www.googleapis.com/auth/cloud-platform"

gcloud config set compute/zone ${ZONE}
gcloud config set project ${PROJECT}

```

# コンテナのビルド
```shell
gcloud builds submit \
    --tag gcr.io/$PROJECT/locust-tasks:latest docker/loadtest
```

# GKE クラスタの作成
```shell
gcloud container clusters create $CLUSTER \
   --preemptible \
   --zone $ZONE \
   --scopes $SCOPE \
   --enable-autoscaling --min-nodes "10" --max-nodes "30" \
   --scopes=logging-write,storage-ro \
   --addons HorizontalPodAutoscaling,HttpLoadBalancing
```

# GKE クラスタに接続
```shell
gcloud container clusters get-credentials $CLUSTER \
   --zone $ZONE \
   --project $PROJECT
```

# Locust のマスターノードとワーカーノードをデプロイします。
```shell
# yml内のcontainers/imageについて、上記でビルド/レジストリアップロードしたプロジェクト/イメージ名に併せて変更
gedit docker/loadtest/gke/config/*.yml

# 起動中のクラスタに設定を適用
kubectl apply -f docker/loadtest/gke/config/locust-master-controller.yml
kubectl apply -f docker/loadtest/gke/config/locust-master-service.yml
kubectl apply -f docker/loadtest/gke/config/locust-worker-controller.yml


# confirm
kubectl get pods -o wide
kubectl get services

# copy ip
kubectl get svc locust-master --watch
EXTERNAL_IP=$(kubectl get svc locust-master -o jsonpath="{.status.loadBalancer.ingress[0].ip}")
echo $EXTERNAL_IP
```

# GKE クラスタの削除
```shell
gcloud container clusters delete $CLUSTER --zone $ZONE
```
