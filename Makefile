build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.support . 

run:
	docker service create --name support-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6028:6028 github.com/turistikrota/service.support:latest

remove:
	docker service rm support-api-turistikrota-com

stop:
	docker service scale support-api-turistikrota-com=0

start:
	docker service scale support-api-turistikrota-com=1

restart: remove build run
	