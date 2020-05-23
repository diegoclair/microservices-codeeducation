#!/bin/bash

cp ./.docker/app/.env-example ./.docker/app/.env
cp .env-example .env
#chmod 777 -R storage

#set our .env dinamically by templates using dockerize
#with this template, we can use dinamically values with {{some value}} in .env as we can see in .docker/app/.env
#We just set the env in the compose and it will set the values to .env
dockerize -template ./.docker/app/.env:.env

#dockerize will try to connect to our service db for 40s
dockerize -wait tcp://db:3306 -timeout 40s


#after connect to db above, we will install the composer and run the migrations
#we did the composer install in dockerfile, but we need to do here too because when we install the composer from dockerfile
#it will create the vendor folder, but after that, the docker-compose will share the volume that we have in our machine, and if
#our machine don't have the vendor folder, we will lose the vendor folder in the container.
#Its is necessary to let the code in the dockerfile for generate a image working, but we need to do here too.
composer install 
php artisan key:generate
php artisan config:cache
php artisan migrate

#start the app
php-fpm