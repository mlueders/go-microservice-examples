## Links

https://go.dev/play/
https://bmuschko.com/blog/go-testing-frameworks
https://medium.com/agrea-technogies/basic-testing-patterns-in-go-d8501e360197
https://levelup.gitconnected.com/go-error-best-practice-f0864c5c2385
https://github.com/swaggo/http-swagger
https://gorm.io/docs/models.html


to learn
- context

project: 
- endpoint to facilitate rest request
- somewhere to persist data
- component tests (use http client to invoke rest calls)
- core tests (tests directly calling the service, no http)
- REST - POST/GET/DELETE/PATCH

want to experiment with
- json encoding/decoding
- web framework
- testing
- observability
  - tracing
  - metrics
  - logging
- messaging
- swagger
- config
- dependency injection
- building - what options are there to hook into the build process.. could a container be started as part of build
- service to service correlation, over rest and over messaging
- authentication
  - jwt tokens
  - ui to service auth
  - service to service auth
  - unauthenticated calls
- optimistic locking
- struct copying - FooResponse vs AddFooRequest... FooEntity vs FooApi, need mechanism to copy most of a struct to a different struct with minimal manual intervention
- paging


struct copy options
https://github.com/jinzhu/copier
https://github.com/ulule/deepcopier
