This is a Go micro service project as detailed in the guides by Ewan Valentine https://ewanvalentine.io/microservices-in-golang-part-1/



## Commands for the API

### Create user
``` json
{
	"service": "shippy.auth",
	"method": "Auth.Create", 
	"request": 
		{ 
			"name": "",
			"email": "", 
			"password": "",
			"company" : ""
			
		}
}
```

### Authenticate user
The result will be a JWT. Use this for future requests

``` json
{ 
	"service": "shippy.auth",
	"method": "Auth.Auth",
	"request":
		{ 
			"email": "", 
			"password": "" 
			
		}
}
```

### Create Consignment
Add in a heading 'token' and put in the JWT

{
      "service": "shippy.consignment",
      "method": "ConsignmentService.Create",
      "request": {
        "description": "",
        "weight": 0,
        "containers": []
      }
}

### Get consignments
Add in a heading 'token' and put in the JWT

{
      "service": "shippy.consignment",
      "method": "ConsignmentService.Get",
      "request": {
        
      }
}

