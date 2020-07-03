TinyURL service is standalone service which creates tiny URLs and stores them in SQL DB.</p> There are many online services available online I created one.</p>
This package comes along with a Docker file which can package an application and create an image. Here are the steps which will be required to set this service lcoally</p>

1. git clone https://github.com/sushiljacksparrow/tinyurl.git
2. docker build -t tinyurl .
3. docker tag tinyurl <DOCKER_REPOSITORY>/tinyurl:1.0
4. docker push <DOCKER_REPOSITORY>/tinyurl:1.0
6. minikube start
7. kubectl apply -f deployment.yml
8. kubectl apply service.yml
9. kubectl port-forward service/tinyurl-service 8080:80

Once it is set up then you can use PostMan and hit URL
- Create Short URL

   POST http://localhost:8080/url/tiny </p>
`
    {
      "original_url": "http://random-very-log-url.com",
      "user": "random-user"
    }
`
<br>
Response 
`
    {
        "tiny_url": "<SHORT_SIX_CHAR_URL>"
    }
`

- Get Original URL

    GET http://localhost:8080/url/long?TinyUrl=<SHORT_SIX_CHAR_URL>&User=<ranom-user>



 
 
