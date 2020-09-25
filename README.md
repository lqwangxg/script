# script
script of shell, vb, js and so on.

## 2020/09/08 Dockerfile, docker-compose.yml included.

## git memo command:
git config --global user.email "lqwangxg@gmail.com" <br>
git config --global user.name "lqwangxg"

## git unset config 
git config --global  --unset http.proxy <br>
git config --global  --unset https.proxy <br>
git config --unset http.proxy <br>
git config --unset https.proxy <br>

## git clone for download
git clone https://github.com/lqwangxg/script 

## git commit and push for upload
git commit -a -m "description of commit" <br>
git push -u origin master 

## maven build command under docker 
$ docker run -it --rm --name myproject  \\ <br>-v "$(pwd)":/usr/src/mymaven  \\ <br>  -w /usr/src/mymaven lqwangxg/maven-java11 mvn --version

### mvn.sh
$ docker run -it --rm --name myproject  \\ <br>-v "$(pwd)":/usr/src/mymaven   \\ <br>-w /usr/src/mymaven lqwangxg/maven-java11 mvn $@
 <br>
$ ./mvn.sh clear package <br>
$ ./mvn.sh deploy 
<br>
docker pull openjdk:8-jre-alpine # sprintboot and tomcat included

## docker-stack-mongo.yml
mongo is a innerdb for MongoExpress and other microservice will access it by innerport,
it no need open outer port 27017 now.
