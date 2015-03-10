# Ready to run app
This is a ready to run application. It is a restful image upload built on martini.
You have possibility to configure your upload endpoints

	{
		"Host": "127.0.0.1",
		"Port": 8080,
		"Endpoints": [{
			"Endpoint": "/storage",
			"FileField": "image",
			"Storage": {
				"Type": "file",
				"Configuration": "./data"
			}
		}]
	}
	
Each endpoint has its own storage. Example above will generate and endpoint where you can do

POST at http://127.0.0.1:8080/storage/ for upload a new image
 
	curl -i -X POST -F image=@data/kth.jpg http://127.0.0.1:8080/storage/
	=> {"filename":"0f4c13e0-05f9-49ae-5fc9-f2dea9e3e16b.jpg","status":"OK"}

PUT http://127.0.0.1:8080/storage/{filename} to update filename with new content
	
	curl -i -X PUT -F image=@data/kth.jpg  http://127.0.0.1:8080/storage/0f4c13e0-05f9-49ae-5fc9-f2dea9e3e16b.jpg
	=> {"filename":"0f4c13e0-05f9-49ae-5fc9-f2dea9e3e16b.jpg","status":"OK"}
	
GET http://127.0.0.1:8080/storage/{filename} to retrieve image content

	http://127.0.0.1:8080/storage/0f4c13e0-05f9-49ae-5fc9-f2dea9e3e16b.jpg

DELETE http://127.0.0.1:8080/storage/{filename} to remove it
	
	curl -i -X DELETE http://127.0.0.1:8080/storage/0f4c13e0-05f9-49ae-5fc9-f2dea9e3e16b.jpg
	=> {"status":"OK"}


## How to run

	go run imageupload.go -h
	
-c="./config/app.json": Path to configurations

-h="": Host to listen on

-p=0: Port number to listen on

You can override host and port be specify -h and -p