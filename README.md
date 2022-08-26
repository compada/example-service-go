# Compada Person Service

Federated graphQL sub-graph concerning persons.

## Getting Started

This service makes heavy use of docker and docker compose. Setting up docker is beyond the scope of this readme.

```shell
cp .env.sample .env
docker compose up -d
```

## Usage

```graphql
query fetchPeople {
  persons {
    id
    name
  }
}
```

## Contributing

```sh
docker compose build
docker compose run --rm go mod tidy
docker compose run --rm go run github.com/99designs/gqlgen init
docker compose run --rm gen
```

## Deployment

This should be done for you via CI/CD. See <https://github.com/compada/cd>.

## Reading

Learn about the various tech powering this application:

- [GraphQL](https://graphql.org)
- [GraphQL Federation](https://www.apollographql.com/docs/federation)
- [Docker](https://docs.docker.com/compose/gettingstarted)
