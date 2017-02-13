#!/bin/sh
set -e

#Chown ???
if [[ ! -z "$APPLICATION_TYPE" ]]
then
  echo "No application type defined"
else
  echo "Lets start $APPLICATION_TYPE logging"
fi

cd /opt/datamgmt/filebeat

./filebeat -c conf.d/$APPLICATION_TYPE.yml -e -v
