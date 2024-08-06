# Cloudinary-API-GO

<img alt="Pict" src="https://i.imgur.com/WWuSUQb.png" align="center"/>

https://github.com/Yunnie-pin/cloudinary-cdn-api

Cloudinary-API-GO is a simple API that allows users to upload images to Cloudinary and store the image details in a PostgreSQL database. The API is built using Go and the Gin framework. The API also includes authentication.

## Endpoints

### Uploader

| Method | Endpoint                   | Description                       |
| ------ | -------------------------- | --------------------------------- |
| POST   | `/api/uploader` | Upload an image to Cloudinary     |


### Files

| Method | Endpoint                     | Description                  |
| ------ | ---------------------------- | ---------------------------- |
| GET   | `/api/buckets`            | Get list buckets              |
| POST    | `/api/path` | Create a new path in a bucket |
| GET    | `/api/files/{{ bucket_id }}/{{ path }}` | Get list files in a bucket    |
| DELETE | `/api/files/{{ bucket_id }}/{{ path }}` | Delete a file in a bucket     |


## Access Documentation

- Postman: https://github.com/Yunnie-pin/cloudinary-cdn-api/blob/main/CDN.postman_collection.json

## Tech Stack

- **Server:** [Golang](https://golang.org/)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **ORM:** [GORM](https://gorm.io/)
- **Cloudinary:** [Cloudinary](https://cloudinary.com/)
- **Router:** [Gin](https://gin-gonic.com/docs/)


## Init

#### Inisiasi Package

```
go mod tidy
```

### Migrasi

```
go mod main.go --migrate
```

#### development

```
go run main.go
```

#### build production

```
go build
```

## Deployment

link : [https://cloudinary-cdn-api-production.up.railway.app/](https://cloudinary-cdn-api-production.up.railway.app/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you find this project useful, please consider giving it a ⭐️ on [GitHub](https://github.com/Yunnie-pin/cloudinary-cdn-api). Your support is greatly appreciated! 😄




