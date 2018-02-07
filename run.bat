docker rm imgapi-app-run
docker rmi imgapi-app
docker build -t imgapi-app .
docker run --publish 8080:8000 --rm --name imgapi-app-run imgapi-app
explorer "http://localhost:8080"
pause