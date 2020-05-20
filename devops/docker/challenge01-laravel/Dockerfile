FROM php:7.3.6-fpm-alpine3.9

RUN apk add --no-cache openssl bash mysql-client
RUN docker-php-ext-install pdo pdo_mysql

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /var/www
RUN rm -rf /var/www/html

RUN curl -s https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

COPY . /var/www
RUN ln -s public html
RUN chmod 755 -R /var/www/storage

RUN composer install && \
    php artisan key:generate && \
    php artisan config:cache

EXPOSE 9000

#entrypoint is the command that run when the container is up
#but the entrypoint needs to keep the container up, otherwise the container will down
ENTRYPOINT ["php-fpm"]
