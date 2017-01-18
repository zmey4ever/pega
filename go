#!/bin/bash
docker stop jenkins;
docker stop apache2;
docker rm jenkins;
docker rm apache2;

docker run -d --name jenkins -p 8080:8080 -p 50000:50000 -p 22222:22 --volumes-from jenkins2   --add-host="master:172.31.33.9"  zmey/jenkins;
docker run --name apache2 -ti -d  -p 8888:80 -p 8822:22  zmey/apache;
docker exec apache2 bash -c "echo 'Ready to go.' > /var/www/html/index.html"
docker exec apache2 bash -c "/etc/init.d/apache2 start"
docker exec apache2 bash -c "/etc/init.d/ssh start"

