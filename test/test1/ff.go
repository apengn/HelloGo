package main

import (
    "strings"
    "fmt"
    "github.com/ghodss/yaml"
    "k8s.io/api/core/v1"
)
const k8sYaml1  = `
apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    labels:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
    name: ttttt-nginx
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: ttttt-nginx
        io.wise2c.service: nginx
        io.wise2c.stack: ttttt
    template:
      metadata:
        labels:
          app: ttttt-nginx
          io.wise2c.service: nginx
          io.wise2c.stack: ttttt
      spec:
        containers:
        - env:
          - name: MY_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: MY_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          image: nginx
          name: nginx
          resources: {}
          securityContext:
            capabilities: {}
            privileged: false
          volumeMounts:
          - mountPath: /etc/resolv.conf
            name: nginx
            subPath: resolv.conf
        volumes:
        - configMap:
            name: ttttt-nginx
          name: nginx
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
    name: ttttt-nginx
  spec:
    clusterIP: None
    ports:
    - name: http-default-port
      port: 8080
    selector:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
- apiVersion: v1
  data:
    resolv.conf: "\nnameserver 10.96.0.10 \n\nsearch ttttt.ns-team-1-env-1.svc.cluster.local\
      \ ns-team-1-env-1.svc.cluster.local svc.cluster.local cluster.local\noptions\
      \ ndots:6"
  kind: ConfigMap
  metadata:
    labels:
      io.wise2c.stack: ttttt
    name: ttttt-nginx
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: ttttt-nginx1
        io.wise2c.service: nginx1
        io.wise2c.stack: ttttt
    template:
      metadata:
        labels:
          app: ttttt-nginx1
          io.wise2c.service: nginx1
          io.wise2c.stack: ttttt
      spec:
        containers:
        - env:
          - name: MY_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: MY_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          image: nginx
          name: nginx1
          resources: {}
          securityContext:
            capabilities: {}
            privileged: false
          volumeMounts:
          - mountPath: /pppppppp
            name: pppppppppp
          - mountPath: /etc/resolv.conf
            name: nginx1
            subPath: resolv.conf
        volumes:
        - name: pppppppppp
          persistentVolumeClaim:
            claimName: ttttt-pppppppppp
        - configMap:
            name: ttttt-nginx1
          name: nginx1
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
  spec:
    clusterIP: None
    ports:
    - name: http-default-port
      port: 8080
    selector:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
- apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    annotations:
      name: pppppppppp
      volume.beta.kubernetes.io/storage-class: sc
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-ppppppppppfffffffffff
  spec:
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 2Gi
- apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    annotations:
      name: pppppppppp
      volume.beta.kubernetes.io/storage-class: sc
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-pppppppppp
  spec:
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 2Gi
- apiVersion: v1
  data:
    resolv.conf: "\nnameserver 10.96.0.10 \n\nsearch ttttt.ns-team-1-env-1.svc.cluster.local\
      \ ns-team-1-env-1.svc.cluster.local svc.cluster.local cluster.local\noptions\
      \ ndots:6"
  kind: ConfigMap
  metadata:
    labels:
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
kind: List
metadata: {}
`


const k8sYaml  = `
apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    labels:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
    name: ttttt-nginx
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: ttttt-nginx
        io.wise2c.service: nginx
        io.wise2c.stack: ttttt
    template:
      metadata:
        labels:
          app: ttttt-nginx
          io.wise2c.service: nginx
          io.wise2c.stack: ttttt
      spec:
        containers:
        - env:
          - name: MY_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: MY_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          image: nginx
          name: nginx
          resources: {}
          securityContext:
            capabilities: {}
            privileged: false
          volumeMounts:
          - mountPath: /etc/resolv.conf
            name: nginx
            subPath: resolv.conf
        volumes:
        - configMap:
            name: ttttt-nginx
          name: nginx
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
    name: ttttt-nginx
  spec:
    clusterIP: None
    ports:
    - name: http-default-port
      port: 8080
    selector:
      app: ttttt-nginx
      io.wise2c.service: nginx
      io.wise2c.stack: ttttt
- apiVersion: v1
  data:
    resolv.conf: "\nnameserver 10.96.0.10 \n\nsearch ttttt.ns-team-1-env-1.svc.cluster.local\
      \ ns-team-1-env-1.svc.cluster.local svc.cluster.local cluster.local\noptions\
      \ ndots:6"
  kind: ConfigMap
  metadata:
    labels:
      io.wise2c.stack: ttttt
    name: ttttt-nginx
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: ttttt-nginx1
        io.wise2c.service: nginx1
        io.wise2c.stack: ttttt
    template:
      metadata:
        labels:
          app: ttttt-nginx1
          io.wise2c.service: nginx1
          io.wise2c.stack: ttttt
      spec:
        containers:
        - env:
          - name: MY_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: MY_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          image: nginx
          name: nginx1
          resources: {}
          securityContext:
            capabilities: {}
            privileged: false
          volumeMounts:
          - mountPath: /pppppppp
            name: pppppppppp
          - mountPath: /etc/resolv.conf
            name: nginx1
            subPath: resolv.conf
        volumes:
        - name: pppppppppp
          persistentVolumeClaim:
            claimName: ttttt-pppppppppp
        - configMap:
            name: ttttt-nginx1
          name: nginx1
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
  spec:
    clusterIP: None
    ports:
    - name: http-default-port
      port: 8080
    selector:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
- apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    annotations:
      name: pppppppppp
      volume.beta.kubernetes.io/storage-class: sc
    labels:
      app: ttttt-nginx1
      io.wise2c.service: nginx1
      io.wise2c.stack: ttttt
    name: ttttt-ppppppppppfffffffffff
  spec:
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 2Gi
- apiVersion: v1
  data:
    resolv.conf: "\nnameserver 10.96.0.10 \n\nsearch ttttt.ns-team-1-env-1.svc.cluster.local\
      \ ns-team-1-env-1.svc.cluster.local svc.cluster.local cluster.local\noptions\
      \ ndots:6"
  kind: ConfigMap
  metadata:
    labels:
      io.wise2c.stack: ttttt
    name: ttttt-nginx1
kind: List
metadata: {}
`

func main() {

    s := strings.Split(k8sYaml, "- ")
    var xx v1.PersistentVolumeClaim
    comparemapNew:=make(map[string]string)
    for _, v := range s {
        if strings.Contains(v, "PersistentVolumeClaim") {
            pvc := "  "+v
            fmt.Println(pvc)
            err := yaml.Unmarshal([]byte(pvc), &xx)
            if err != nil {
                fmt.Println(err)
                return
            }
            comparemapNew[xx.Name]=xx.Name
        }
    }

    s2 := strings.Split(k8sYaml1, "- ")
    var xx2 v1.PersistentVolumeClaim
    comparemapOld:=make(map[string]string)
    for _, v := range s2 {
        if strings.Contains(v, "PersistentVolumeClaim") {
            pvc := "  "+v
            fmt.Println(pvc)
            err := yaml.Unmarshal([]byte(pvc), &xx2)
            if err != nil {
                fmt.Println(err)
                return
            }
            comparemapOld[xx.Name]=xx.Name
        }
    }


    fmt.Println("new",comparemapNew)

    fmt.Println("old",comparemapOld)
}

