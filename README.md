# k8s-test

A test repository to get more familiar with Docker and Kubernetes. The goal is to have a simple frontend application that can talk to a backend application within Kubernetes.

---

### Build and run locally

As long as you have a somewhat recent version of [Golang](https://golang.org/dl/) you should be able to use `go run` to build and run both the backend and the frontend application.

```bash
# Start backend
$ cd backend/
$ go run server.go
2021/05/09 16:10:42 Listening on :8000

$ curl -s http://127.0.0.1:8000/users
[{"id":1,"name":"Daniel1337","email":"daniel@mail.se"},{"id":2,"name":"AnnaPanna","email":"anna@mail.se"},{"id":3,"name":"Trollfar","email":"troll@mail.se"},{"id":4,"name":"Kakburken","email":"kaka@mail.se"}]

# Start frontend
$ cd frontend/
$ go run server.go
2021/05/09 16:11:02 Listening on :8080
```

You can also run `docker compose up` to start them as containers.
```bash
$ docker compose up
...
backend_1   | 2021/05/09 14:13:49 Listening on :8000
frontend_1  | 2021/05/09 14:13:49 Listening on :8080
```

### Run in Kubernetes

I don't have access to fancy cloud services with cool names, but this seems to work in a local Kubernetes environment (Docker for Windows). 
Images for both the frontend and backend application are built and pushed to DockerHub when a release is created in the GitHub repository.

```bash
$ kubectl apply -f k8s/
deployment.apps/backend created
service/backend created
deployment.apps/frontend created
service/frontend created

$ kubectl get deployments
NAME       READY   UP-TO-DATE   AVAILABLE   AGE
backend    1/1     1            1           42s
frontend   2/2     2            2           42s

$ kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
backend-9f889bb86-pxxsl     1/1     Running   0          56s
frontend-5b6f64b87b-jx4ql   1/1     Running   0          56s
frontend-5b6f64b87b-nbgg5   1/1     Running   0          56s

$ kubectl get services
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
backend      ClusterIP   10.107.128.212   <none>        8000/TCP         66s
frontend     NodePort    10.108.161.249   <none>        1338:30493/TCP   66s
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP          26h

$ kubectl describe service frontend | grep ^NodePort
NodePort:                 <unset>  30493/TCP

# You should be able to visit http://127.0.0.1:30493 in the browser and get a list of users.

# Clean up
$ kubectl delete service frontend backend
service "frontend" deleted
service "backend" deleted

$ kubectl delete deployment frontend backend
deployment.apps "frontend" deleted
deployment.apps "backend" deleted
```


