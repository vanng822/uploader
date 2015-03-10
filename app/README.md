# Ready to run app
This is a ready to run application. It is build on martini.
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
	
Each endpoint has its own storage.

## How to run

	go run imageupload.go -h
	
-c="./config/app.json": Path to configurations

-h="": Host to listen on

-p=0: Port number to listen on

You can override host and port be specify -h and -p