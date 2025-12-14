APP_NAME=hareta
DEPLOY_CONNECT=root@159.223.51.24

docker rm -f ${APP_NAME}
docker image rm ${APP_NAME}

echo "Docker building..."

GOOS="linux" GOARCH="amd64" go build -o app
docker build -t ${APP_NAME} .
echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}

echo "Deploying..."
scp -o StrictHostKeyChecking=no ./${APP_NAME}.tar ${DEPLOY_CONNECT}:~
ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/stg.sh
rm -f ./${APP_NAME}.tar
echo "Done"