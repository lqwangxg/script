# script
script of shell, vb, js and so on.

## 2020/09/08 Dockerfile, docker-compose.yml included.

## git memo command:
git config --global user.email "lqwangxg@gmail.com" <br>
git config --global user.name "lqwangxg"

git clone https://github.com/lqwangxg/script
cd script

git add .
git commit -a -m "description of commit"

## maven build command under docker 
$ docker run -it --rm --name myproject  \\ <br>-v "$(pwd)":/usr/src/mymaven  \\ <br>  -w /usr/src/mymaven lqwangxg/maven-java11 mvn --version

### mvn.sh
$ docker run -it --rm --name myproject  \\ <br>-v "$(pwd)":/usr/src/mymaven   \\ <br>-w /usr/src/mymaven lqwangxg/maven-java11 mvn $@
 <br>
$ ./mvn.sh clear package <br>
$ ./mvn.sh deploy 
<br>
docker pull openjdk:8-jre-alpine # sprintboot and tomcat included
