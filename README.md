# script
script of shell, vb, js and so on.

## 2020/09/08 Dockerfile, docker-compose.yml included.

## git memo command:
git config --global user.email "lqwangxg@gmail.com"
git config --global user.name "lqwangxg"

git clone https://github.com/lqwangxg/script
cd script

git add .
git commit -m "description of commit"

## maven build command under docker 
$ docker run -it --rm --name myproject -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven lqwangxg/maven-java11 mvn --version

### mvn.sh
$ docker run -it --rm --name myproject -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven lqwangxg/maven-java11 mvn $@

$ ./mvn.sh clear package
$ ./mvn.sh deploy 

