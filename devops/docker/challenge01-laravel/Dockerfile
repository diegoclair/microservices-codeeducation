FROM php:7.3.6-fpm-alpine3.9

RUN apk add bash mysql-client
RUN docker-php-ext-install pdo pdo_mysql

WORKDIR /var/www
RUN rm -rf /var/www/html

RUN curl -s https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

COPY . /var/www
RUN ln -s public html
RUN chmod 755 -R /var/www/storage

RUN composer install && \
     cp .env.example .env && \
     php artisan key:generate && \
     php artisan config:cache

EXPOSE 9000
ENTRYPOINT ["php-fpm"]
