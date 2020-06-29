# Docker image from DockerHub
- [Clique aqui](https://hub.docker.com/repository/docker/diegoclair/challenge04-03)

# Challenge description
- Create a webserver using port 8000 that returns a bold string ```Code.education Rocks.```
- The return of this string should be based on function greeting. This function receive a string parameter and return this string with <b></b> tags
- Create a test for this greeting function.
- Create a optimized Dockerfile image.
- Create a cloudbuild file to CI process.
- Create a trigger in GCP to read the cloudbuild file always when a push is made to any branch.
- Create kubernetes type LoadBalance service to access on the browser

## Challenge Continuous Delivery(CD) description
- Create kubernetes files to do the deploy to it
- Use the CI process (cloudbuild.yaml) for all branchs, except for master
- Use CI&CD (cloudbuild.prod.yaml) for branch master
- Configure this rules on GCP Cloud

# Start application
- Run the command: <b>```docker-compose up```</b>
- Go to your browser and put http://localhost:8000