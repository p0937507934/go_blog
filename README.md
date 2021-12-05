# go_blog

## Quick Start For Deploy

```sh
# clone repo
git clone https://github.com/p0937507934/go_blog.git

# install dependency
cd go_blog

# deploy server on 8080 port
sudo docker-compose up
```
## API

User

POST /user/login

| Description |
| ------------|
|User Login|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| username    | string  |  true             |
| password    | string  |  true             |


POST /user/register

| Description |
| ------------|
|User Register|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| username    | string  |  true             |
| password    | string  |  true             |

---
Post 

GET /post

| Description |
| ------------|
|Get all post data|

GET /post/:id

| Description |
| ------------|
|Get post data by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | int  |  true             |

POST /post

| Description |
| ------------|
|Add a post|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| title    | string  |  true             |
| content    | string  |  true             |

PUT /post/:id

| Description |
| ------------|
|Update a post by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | string  |  true             |
| title    | string  |  true             |
| content    | string  |  true             |

DELETE /post/:id

| Description |
| ------------|
|delete a post by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | string  |  true             |

PUT /post/lock/:id

| Description |
| ------------|
|lock a post,just for admin|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | string  |  true             |

---
Comment

GET /comment


| Description |
| ------------|
|Get all comment data|

GET /comment/:id

| Description |
| ------------|
|Get comment data by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | int  |  true             |


POST /comment

| Description |
| ------------|
|Add a comment|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| pid    | int  |  true             |
| content    | string  |  true             |

PUT /post/:id

| Description |
| ------------|
|Update a comment by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | string  |  true             |
| content    | string  |  true             |
|status|int|false

DELETE /comment/:id

| Description |
| ------------|
|delete a comment by id|

| Parameter |         type            |     required |
| ------- | ----------------------| -----------------|
| id    | string  |  true             |



## Project structure
- controller
  - Validate all parameters and call services what you need
- services
  - Implement business logic
- router
  - allocate all the routes
- repository
  - call database's manipulation
- middleware
  - implement middleware
- utils
  - encapsulate some common function



  
