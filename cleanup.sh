docker-compose down && \
docker rmi -f alphatree && \
docker image prune -f && \
docker volume prune -f && \
docker network prune -f && \
docker-compose up