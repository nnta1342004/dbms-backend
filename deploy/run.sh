APP_NAME=hareta
VERSION=latest

docker run -d --name ${APP_NAME} \
    --network my-net \
    -e REDIS_GO_REDIS_URI="redis://:ngogiahan7203@redis:6379/0" \
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
    ${APP_NAME}:${VERSION}