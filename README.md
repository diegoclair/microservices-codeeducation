<h2 align="center"> 
	Microservices by 
  <a href="https://code.education/">Code Education</a>
</h2>
<h1 align="center">
    <img alt="CodeEducationTitle" title="#CodeEducationTitle" src=".github/title.png" width="650px" />
</h1>

<h4 align="center"> 
	Developing modern and scalable applications with microservices <br/>
  The digital transformation begin at the developer 
</h4>
<p align="center">
  <a href="https://github.com/diegoclair/microservices-codeeducation/commits/master">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/diegoclair/microservices-codeeducation?label=Last%20commit">
  </a>

  <img alt="GitHub languages count" src="https://img.shields.io/github/languages/count/diegoclair/microservices-codeeducation" style="margin-left:3px;">

  <a href="https://www.linkedin.com/in/diegoclair/">
    <img alt="Made by Diego Clair" src="https://img.shields.io/badge/Made%20by-Diego%20Clair-informational" style="margin-left:3px;">
  <img src="https://img.icons8.com/color/2x/linkedin.png" width="24px" height="25.1px" style="margin-left:-3px; margin-bottom: -2.9px">  
  </a>
</p>

## 📝 Project 
In this training we'll develop an application like Netflix, where the user will be able to create an account, to make a subscription, look for a movie and watch them.
### 💻 Features 
We are going to develop the following microservices:
* Authentication Single Sign On: <br/>
  <img alt="Single-Sign-On" title="#Single-Sign-On" src=".github/single-sign-on.png" width="350px" />

* User Subscription: <br/>
  <img alt="User-Subscription" title="#User-Subscription" src=".github/single-sign-on.png" width="350px" />

* Search: <br/>
  <img alt="Search" title="#Search" src=".github/search.png" width="350px" />

* Video Catalog: <br/>
  <img alt="Video-catalog" title="#Video-catalog" src=".github/video-catalog.png" width="350px" />

* Encoding videos: <br/>
  <img alt="Encoding-video" title="#Encoding-video" src=".github/encoding-video.png" width="350px" />

## Status
In development 🚧

## 👨‍💻 Technologies and content 
* Software Architecture 🚧
  - Introduction ✅
  - Monolitic architecture ✅
  - Microservice architecture ✅
  - Microservice communication types ✅
  - SOLID ✅
  - DDD ✅

* Devops 🚧
  - [Docker](https://www.docker.com/) ✅
  - [GitFlow](https://datasift.github.io/gitflow/IntroducingGitFlow.html) ✅
  - [CI](https://codeship.com/continuous-integration-essentials) (Continuous Integration) ✅
     - CI using Google CloudBuild ✅
  - [Kubernetes](https://kubernetes.io/) 🚧
     - Kubernetes using GCP (Google Cloud Plataform)
  - [CD](https://codeship.com/continuous-integration-essentials)
* Programming languages 🕒
  - [Golang](https://golang.org/)
  - [Nodejs](https://nodejs.org/en/about/)
  - [PHP](https://laravel.com/) (Laravel)
  - [React](https://reactjs.org/)
* Messaging Queues 🕒
  - [RabbitMQ](https://www.rabbitmq.com/)
  - [Kafka](https://kafka.apache.org/)
* Databases 🕒
  - [Redis](https://redis.io/)
  - [MySQL](https://www.mysql.com/)
  - [Elasticsearch + kibana](https://www.elastic.co/)

## 🚧 Extras 
In this repo, you'll find the challenges that I did during the course.  

* [Repo challenge 01](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge01-laravel) - Docker 🐳✅:
  - Challenges:
    - Publishing a Laravel image to dockerhub
    - Using Dockerize template to define environment variables for nginx image in docker-compose.yaml
    
* [Repo challenge 02](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge02-golang) - Docker 🐳✅:
  - Challenges:
    - Publishing a Golang image to dockerhub that print the text: 'Code.education Rocks!'
    
* (CI) Continuous Integration challenge:
  - Challenges:
    - [Repo](https://github.com/diegoclair/ci-gcp): Create a docker-compose image and push to GCP Container Register.
    - Implement CI process at [Laravel Repo](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge01-laravel) to each pull request created and to any branch.
    
* [Repo challenge 03](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge03-CI-Go) -  CI an Docker 🐳✅:
  - Challenges:
    - Create a project in Go with a function <b>soma</b> and add a test for it.
    - Implement the CI process that will execute the unit test, generate and push the image to GCP Container Register and with trigger with github repo to start CI process when execute any push to any branch.
    
* [Repo challenge 04](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge04-ks8) - Kubernetes ✅: 
  - [Challenge 01](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge04-ks8/challenge01-nginx) Create a kubernets service with nginx image that returns the message ```Code.education Rocks.```
  - [Challenge 02](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge04-ks8/challenge02-mysql) - Kubernetes with mysql, using persistent volume and secret with environment variables.
  - [Challenge 03](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge04-ks8/golang) Create a webserver that returns a bold string ```Code.education Rocks.``` with test, CI process and kubernetes.
* [Repo challenge 05](https://github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge05-Go-hpa) - Kubernetes hpa (Horizontal Pod Autoscaler)
