





#Then build the image:

docker build -t server_new:latest -f Dockerfile.server_new .
docker build -t client_new:latest -f Dockerfile.client_new .


# Push to DockerHub or your container registry

docker push debolek/client_new:latest
docker push debolek/server_new:latest

