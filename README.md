# Imger
[![Go Report Card](https://goreportcard.com/badge/github.com/RexterR/imger)](https://goreportcard.com/report/github.com/RexterR/imger)<br>
IMGER it's an HTTP service based on Go to process Image with profile store functionality.
### Features
- Image Processing and Manipulation
- Image caching (Redis)
- Predefined Profiles (MongoDB)
-  HTTP API Documentation (Swagger Specification) and MD
- Robust Error Handling
- Dockerized
- Makefile
- Testing using TDD approach
- Graceful shutdown
- Healthcheck
### Docker Support
- Dockerfile (Development)
- Dockerfile.CI (Production)
- docker-compose.yaml (Run application)
- docker-compose.test.yaml (Testing Manually)
### Setup Project
- Clone repository: `https://github.com/RexterR/imger.git`
- Install dependencies: `make deps`
- Run using docker: `docker-compose up` or `make docker`
- Open application: Application Listening on`http:localhost:4005`
## Effects

The engine behind image manipulation is the fabulous library: [github.com/disintegration/imaging](github.com/disintegration/imaging)<br>
**Available Effects**
|Effect Name|Eg Spec 
|-----------|-------------------------------------------------------------------------|
|resize     |`{"id":"resize","parameters":{"width":25,"height":50,"filter":"linear"}}`|
| crop      |`{"id":"crop","parameters":{"rectangle":[0,0,202,150]}}`                 |
|blur       | `{"id":"blur","parameters":{"sigma":0.9}`                               |
|brightness | `{"id":"brightness","parameters":{"percentage":-50}}`                   |
|contrast   |`{"id":"contrast","parameters":{"percentage":100}}`                      |
|gamma      |`{"id":"gamma","parameters":{"gamma":0.2}}`                              |

It's possible to combine multiple effects:
```
[{"id":"overlay","parameters":{"position":[25,75],"url":"https://goo.gl/UBrXeo"}},{"id":"overlay","parameters":{"position":[22,-35],"url":"https://goo.gl/aEkkDh"}}, {"id":"crop","parameters":{"rectangle":[0,0,202,150]}}]
```
## Profiles

If you don't want to specify filters in URL, you can create a profile with all pre configured filters and then use it in query parameters `&profile={profile-id}`.
## HTTP Docs
[**HTTP Docs**](https://github.com/sumandas0/imger/blob/main/HTTP.md)
