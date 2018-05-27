#!/usr/bin/env bash

if [ ! -d apache-maven-3.5.3 ]; then
    wget http://apache.mirror.amaze.com.au/maven/maven-3/3.5.3/binaries/apache-maven-3.5.3-bin.tar.gz
    tar xvf apache-maven-3.5.3-bin.tar.gz
    ln -s apache-maven-3.5.3/bin/mvn mvn
    rm apache-maven-3.5.3-bin.tar.gz
fi

./dev/change_scala_version.sh 2.11
./mvn clean package -B -Pbuild-distr -DskipTests -Pspark-2.2 -Phadoop-2.7 -Pscala-2.11
