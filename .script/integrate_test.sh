#!/usr/bin/env bash

set -e
docker compose -p webook -f .script/integration_test_compose.yml down -v
docker compose -p webook -f .script/integration_test_compose.yml up -d
go test -race -shuffle=on -failfast -tags=e2e -count=1 ./...
docker compose -p webook -f .script/integration_test_compose.yml down -v
