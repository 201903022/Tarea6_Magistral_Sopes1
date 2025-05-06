# Despliegue de una App Go en GKE usando NodePort

## ✨ Objetivo

Desplegar una aplicación web escrita en Go en Google Kubernetes Engine (GKE), exponiéndola al exterior usando un servicio `NodePort`.

---

## 🔧 Tecnologías Utilizadas

* Google Kubernetes Engine (GKE)
* Kubernetes (`kubectl`)
* Docker
* Lenguaje de programación Go
* Helm (opcional)

---

## 📂 Estructura del Proyecto

```
T6/
├── main.go              # Aplicación Go que responde "Hola Mundo"
├── Dockerfile           # Imagen Docker para compilar y ejecutar Go
├── Deployment.yaml      # Despliegue en Kubernetes
├── nodePort.yaml        # Servicio tipo NodePort
```

---

## 💡 main.go

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola Mundo desde Go en Kubernetes!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Servidor iniciado en puerto 8080...")
    http.ListenAndServe(":8080", nil)
}
```

---

## 📦 Dockerfile

```Dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod init T6 && go build -o server

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

---

## 📦 deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: helloworld-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: helloworld
  template:
    metadata:
      labels:
        app: helloworld
    spec:
      containers:
      - name: helloworld
        image: <REGISTRO>/<IMAGEN>:<TAG>
        ports:
        - containerPort: 8080
```

---

## 🏠 nodePort.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: helloworld-nodeport
  labels:
    app: helloworld
spec:
  type: NodePort
  selector:
    app: helloworld
  ports:
  - port: 80
    targetPort: 8080
    nodePort: 30080
```

---

## ⚡ Crear regla de firewall en GCP

```bash
gcloud compute firewall-rules create allow-nodeport-30080 \
  --allow tcp:30080 \
  --source-ranges=0.0.0.0/0 \
  --direction=INGRESS \
  --priority=1000 \
  --network=default \
  --description="Permitir acceso al servicio NodePort 30080 para cualquier IP"
```

---

## 🚀 Acceso al servicio desde navegador

Obtener IP externa del nodo:

```bash
kubectl get nodes -o wide
```

Ejemplo de URL de acceso:

```
http://34.58.100.243:30080
```

---

## 👍 Verificación desde dentro del clúster

```bash
kubectl run curlpod --image=radial/busyboxplus:curl -i --tty --rm
```

Dentro del pod:

```sh
curl http://helloworld-nodeport.default.svc.cluster.local:80
```

---

## 📅 Notas finales

* No se requiere `Ingress` para acceder vía `NodePort`
* La solución es ideal para pruebas, demostraciones y desarrollo
* En producción se recomienda usar `Ingress` con dominio + TLS

---

**Autor:** Jonathan
**Fecha:** Mayo 2025
