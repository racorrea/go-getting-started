$ docker build -t my-golang-app .
$ docker run -p 10000:10000 -it --rm --name my-running-app my-golang-app