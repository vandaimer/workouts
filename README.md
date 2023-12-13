# Workouts

This service is responsible to manage Workouts.

![enter image description here](https://www.happycoders.eu/wp-content/uploads/2023/01/hexagonal-architecture-ddd-domain-driven-design-600x484.png)(https://www.happycoders.eu/software-craftsmanship/hexagonal-architecture/)

Inside internal, we have the whole source code.

- api: Where we implement the external world entry point for the service. For example, REST API, GraphQL others...
- domain: It contains the core of the project. The models/entities, api-ports, spi-ports, services, domain rules.
	- api-ports: Responsible to define how external world can call the domain
		- Exemple: REST API,  gRPC, GraphQL... They need create api-ports to talk to the domain
	- spi-ports: Responsible to define how the model will communicate with the external world. We define the interfaces/ports saying how the external services needs to adapt to let the domain talk to them
		- Database (PostgreSQL), Cache (Redis), Message Broker (Kafka, RabbitMQ)
	- model: The place we define the entities, for this case Workout model.
	- services: The place where we implement the business logic, usually called by use-cases on clean architecture (or similar)
	- business rules: Where, we put the pure functions that can be implemented as a very simple function, receiving input, returning output or raise an error, without taking with the external world (database, for example).
- spi: Everything that isn't part of this project. Systems, platforms, external micro services. Examples: Auth0, PostgreSQL, another internal micro service. Those resource belong to spi because if we change it from PostgreSQL to MySQL, for instance, the domain **won't and shouldn't change**.

# How to run locally

Having docker and docker-compose installed, you should run:
- `cp .env.example .env`
- `docker-compose up`

And for testing, the following curl will perform the HTTP request.

```bash
curl --request POST \
  --url 'http://localhost:3000/api/v1/analyse?nweeks=1' \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --data '[
	{
		"distance": 10000,
		"time": 3600,
		"timestamp": "2022-09-04T13:43:28.073909Z"
	}
]'
```

and getting the following response:

```bash
{
	"medium_distance": 8000,
	"medium_time": 3000,
	"max_distance": 22000,
	"max_time": 7000,
	"medium_weekly_distance": 30000,
	"medium_weekly_time": 9000,
	"max_weekly_distance": 30000,
	"max_weekly_time": 9000
}
```

## Improvements

- Implement authentication/authorization
- Implement unit tests
- Remove the hard coded response and make it dynamic
- Implement db migration
	- Removing the creation of the db table via docker-compose (workaround)
- Make available open an OpenAPI 3.0 specification
- Improve input body validation
- Create a Makefile to improve the development experience
