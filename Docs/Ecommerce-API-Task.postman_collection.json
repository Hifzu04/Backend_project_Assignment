{
	"info": {
		"_postman_id": "f3d892f8-780a-446f-ac26-5dd088f14180",
		"name": "Ecommerce-API-Task",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "51265066",
		"_collection_link": "https://rahmankhan372003-5636891.postman.co/workspace/MyWorkspace~eba3c6ce-f985-491b-87af-f03cd53e6a8c/collection/51265066-f3d892f8-780a-446f-ac26-5dd088f14180?action=share&source=collection_link&creator=51265066"
	},
	"item": [
		{
			"name": "Register",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"name\": \"Admin User\",\n  \"email\": \"admin@test.com\",\n  \"password\": \"password123\",\n  \"role\": \"admin\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/auth/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"auth",
						"register"
					]
				}
			},
			"response": []
		},
		{
			"name": "Login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"email\": \"admin@test.com\",\n  \"password\": \"password123\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/auth/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"auth",
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "GetProduct",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": [
						{
							"key": "token",
							"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluQHRlc3QuY29tIiwiUm9sZSI6ImFkbWluIiwiZXhwIjoxNzc0Nzk0NzQzfQ.ecVMT60ico4O7e7dXdDrDhpJakIBC8MrTLcM6wf6O_A",
							"type": "string"
						}
					]
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/v1/products/",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"products",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "CreateProduct",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": [
						{
							"key": "token",
							"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluQHRlc3QuY29tIiwiUm9sZSI6ImFkbWluIiwiZXhwIjoxNzc0Nzk0NzQzfQ.ecVMT60ico4O7e7dXdDrDhpJakIBC8MrTLcM6wf6O_A",
							"type": "string"
						}
					]
				},
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"name\": \"MacBook 2\",\n  \"price\": 1999.99,\n  \"description\": \"Powerful laptop for developers\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/products/",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"products",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "UserRegister",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"name\": \"Regular Customer\",\n  \"email\": \"user@test.com\",\n  \"password\": \"password123\",\n  \"role\": \"user\" \n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/auth/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"auth",
						"register"
					]
				}
			},
			"response": []
		}
	]
}