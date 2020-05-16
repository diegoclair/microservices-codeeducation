#!/bin/bash

#dockerize will try to connect to our service db for 40s
dockerize -wait tcp://db:3306 -timeout 40s

#after connect to db above, we will run the migrations
php artisan migrate
php-fpm