## Links

https://go.dev/play/
https://go.dev/doc/effective_go
https://github.com/golang/go/wiki/CodeReviewComments
https://bmuschko.com/blog/go-testing-frameworks
https://medium.com/agrea-technogies/basic-testing-patterns-in-go-d8501e360197
https://levelup.gitconnected.com/go-error-best-practice-f0864c5c2385
https://github.com/swaggo/http-swagger
https://gorm.io/docs/models.html
https://github.com/pkg/errors

to learn
- context

project: 
- endpoint to facilitate rest request
- somewhere to persist data
- REST - POST/GET/DELETE/PATCH

want to experiment with
- json encoding/decoding
- web framework
  - https://github.com/diyan/go-web-framework-comparison
  - https://github.com/go-kit/kit
  - https://github.com/labstack/echo
  - https://www.gorillatoolkit.org/
  - https://github.com/goji/goji
  - https://github.com/gin-gonic/gin
  - https://github.com/gocraft/work
  - https://gobuffalo.io/en/
  - https://github.com/paulbellamy/mango
  - https://goa.design/
  - https://github.com/zeromicro/go-zero
  - https://gofiber.io/
  - https://revel.github.io/
- testing
  - component tests (use http client to invoke rest calls)
  - core tests (tests directly calling the service, no http)
- observability
  - tracing
  - metrics
  - logging
- messaging
- swagger
- clients
  - https://github.com/go-resty/resty
  - generation
    - https://github.com/Stratoscale/swagger
  - resilience
    - retry + circuit breaker
- config
- dependency injection
  - https://github.com/google/wire
  - https://blog.drewolson.org/go-dependency-injection-with-wire
  - https://github.com/uber-go/dig
  - https://pkg.go.dev/go.uber.org/fx
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
- how to handle optional data - input (json) and storage (sql, etc)
  - https://github.com/golang/go/issues/11939
  - https://github.com/guregu/null
- database migrations
  - https://github.com/golang-migrate/migrate
  - https://github.com/pressly/goose

struct copy options
https://github.com/jinzhu/copier
https://github.com/ulule/deepcopier
