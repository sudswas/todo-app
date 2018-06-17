# The todo-app

This project has been written in less 10 hours time.
The project aims at demonstrating the following:
  - A sample to-do app written in golang, using mariadb as a database.
  - Deployment in a kubernetes cluster and exposing the service using NodePort.
  - A commandline utility written in python, that could be used to interact with the service deployed on kubernetes
  - A sample Makefile (untested), which should help someone bootstrap fast.
  
 
 # What this project IS NOT
 
 This project is just a fun work, and does not gurantee anything about it's working.
 This project has been rapidly prototyped with very less thought on design and more stress on functionality.
 There are many things that can be better and should be fixed. That's for another iteration :)
 
 # Design
 
 Following are the sequence of events that happen:
 
 1. mariadb is bootstrapped in a kuberentes POD alongside the app container. So basically this deployment POD consists of two containers - one for mariadb and one for the app.
 2. The app code in `web/webserver.go` creates the table after sleeping for 6 seconds. This is a simple hack to wait till the mariadb container comes up. There are n different ways of handling this better.
 3. A simple to-do table is created in mariadb. The database uses ephemeral storage and hence if you kill the container/deployment, all your to-do configs would be lost. 
 4. The to-do table consists of 4 fields - `ID, task, duedate, status`.
 5. `db/db_utils.go` is responsible for connecting to the DB. However, we can definitely do better if we use some connection pool instead of connecting to the db everytime a call is made.
 6. `web/web_server.go` uses mux and creates a simple REST based web service in GO that exposes few functionalities (discussed below).
 7. There's a commandline tool inside the `cli` folder that should help anyone use this tool using a python based cli.
 
 # Installation pre-requisites
 
 1. Please have go installed and `GOPATH` properly configured.
 2. Please have a kubernetes deployment ready that you can deploy this app to.
 3. Please have python installed.
 4. Please have `make` installed.
 5. Please have `docker` installed and in working condition in the machine you plan to build the code.
 
 # Installation
 
 The `Makefile` in the root directory gives a easy way to install the code. However, I ran out of battery and could not test it fully. Here's what the steps manually should be like:
 
 1. Clone the code inside `$GOPATH/src` folder.
 2. Perform go gets (mentioned in the make file as well:
    - go get -u github.com/go-sql-driver/mysql
    - go get -u github.com/gorilla/mux
 3. Perform the following: `go build -o main myapp/web` This will generate the `main` binary in your `src` folder.
 4. Perform : docker build -t myapp:latest .
 5. cd kubernetes/ ; create -f todo-deployment.yaml 
 6. Stay in the same directory and do: create -f todo-svc.yaml
    Note: NodePort is used to expose the service. If you have an ingress controller configured, please create a ingress yaml       and you should be able to use it on top of this.
 7. Ensure that the kube deployment is successful and the pods are up and running. Sample out:
    #### kubectl get pods
    `todo-deployment-4134519886-smjmr                                 2/2       Running   1          36m`
 
 # Sample flow
 
 The easiest way to use the app is by the `commandline.py` file inside the `cli` folder. Please edit the `config.in` file with the IP:port number of the service as obtained from your kubernetes cluster. Sample outputs:
 
    # python commandline.py --add "Buy milk" "2018-06-18"
    Request successfully sent
    # python commandline.py --get all
    [{"Task":"Grab milk","Dueby":"2018-08-09","Status":"Upcoming"},{"Task":"Pay bill","Dueby":"2018-04-   12","Status":"Incomplete"}]
    # python commandline.py --get today
    [{"Task":"Discuss kube","Dueby":"2018-06-17","Status":"Upcoming"}]
    # python commandline.py --getbydate "2018-06-18"
    [{"Task":"Buy milk","Dueby":"2018-06-18","Status":"Upcoming"}]
    # python commandline.py --getbydate "2018-04-12"
    [{"Task":"Newday","Dueby":"2018-04-12","Status":"Incomplete"}]
    
 The status of a given to-do is calculated automatically by the code once the date expires. A way to mark a given to-do as complete will be an enhancement to this app.


# Limitations

Because of the rapid nature of development, there are many good practices which have got overseen.
There are plenty of new upcoming features that could be done like:

1. Marking a to-do as complete.
2. Fetching to-dos with week etc.
3. Making the data visible via rendering of an html page.
4. Kubernetes secrets could be used to pass the passwords to the apps.
5. Connection pools should be implemented to improve DB access.
6. Tests should be written to check the webserver functionalities.
7. The mariadb container should have persistent storage by using to a PersistentVolume in Kubernetes.







 
 
