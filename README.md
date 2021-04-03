# MicroservicesPOA

##Execution
Pour exécuter le programme, il faut dans un premier temps compiler le projet, puis lancer l'exécutable  :
```
go build
.\MicroservicesPOA.exe (pour Windows)
./MicroservicesPOA     (pour Linux) 
```

Le service est maintenant lancé sur le port 8000 en local.

##Utilisation

Pour tester je conseille l'utilisation de Postman.
Voici différents moyens de tester le microservice.

###Ajout d'un post
```
curl.exe -X POST localhost:8000/posts/ -H "Content-Type: application/json" -d 
'{
    "id":"12345",
    "title":"Test",
    "author":"Valentin ROCHE",
    "content":"Ca marche",
    "published_on":"2021-03-29T18:53:58.9592333-04:00"
}'
```

###Récupération de tous les posts
```
curl.exe -X GET localhost:8000/posts/
```
Ou via http://localhost:8000/posts/

###Récupération de post par ID 
```
curl.exe -X GET localhost:8000/posts/1234
```
Ou via http://localhost:8000/posts/1234

###Récupération de tous les posts d'un auteur
```
curl.exe -X GET http://localhost:8000/posts/?author=Valentin ROCHE
```
Ou via http://localhost:8000/posts/?author=Valentin ROCHE

###Recherche de posts par titre
```
curl.exe -X GET http://localhost:8000/posts/?query=Te
```
Ou via http://localhost:8000/posts/?query=Te

###Supression d'un poste
```
curl.exe -X DELETE http://localhost:8000/posts/1234
```

