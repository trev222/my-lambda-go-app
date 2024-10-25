


# install These 
golang - https://go.dev/doc/install
docker - https://docs.docker.com/engine/install/
nodejs - https://nodejs.org/en/download/package-manager
npm - https://docs.npmjs.com/downloading-and-installing-node-js-and-npm
nodemon - npm install -g nodemon

Optional
iTerm 2 - https://iterm2.com/
VsCode - https://code.visualstudio.com/download


# scripts

# run 
npm start

# stop
npm stop

# build (to manually build)
npm build

aws ecr get-login-password --region ap-northeast-1 | \
docker login --username AWS --password-stdin 339712943371.dkr.ecr.ap-northeast-1.amazonaws.com


aws ecr create-repository --repository-name my-lambda-go-app
docker tag my-lambda-go-app:latest 339712943371.dkr.ecr.ap-northeast-1.amazonaws.com/my-lambda-go-app:latest
docker push 339712943371.dkr.ecr.ap-northeast-1.amazonaws.com/my-lambda-go-app:latest

sam build
