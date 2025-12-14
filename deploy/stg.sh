#!/usr/bin/env bash

APP_NAME=hareta
docker rm -f ${APP_NAME}
docker image rm -f ${APP_NAME}
docker load -i ${APP_NAME}.tar
rm -f ${APP_NAME}.tar

docker run -d --name ${APP_NAME} \
    --network my-net \
    -e REDIS_GO_REDIS_URI="redis://:ngogiahan7203@redis:6379/0" \
    -e MYSQL_GORM_DB_URI="root:master_root_password@tcp(159.223.51.24:32804)/my_database?charset=utf8mb4&parseTime=True&loc=Local" \
    -e MAIL_USERNAME="leductoan3082004@gmail.com" \
    -e MAIL_PORT="587" \
    -e MAIL_PASSWORD="pyvufdtbjociioiz" \
    -e MAIL_HOST="smtp.gmail.com" \
    -e MAIL_ADDRESS="smtp.gmail.com:587" \
    -e JWT_SECRET_KEY="giahanismylover" \
    -e AWS_REGION="ap-southeast-1" \
    -e AWS_DOMAIN="d2csq352pki9k7.cloudfront.net" \
    -e AWS_BUCKET="hareta-bucket" \
    -p 3000:3000 \
    -e VIRTUAL_HOST="api.hareta.me" \
    -e LETSENCRYPT_HOST="api.hareta.me" \
    ${APP_NAME}