# Dev Container Manager (dcm)

dcm makes ephemeral development environments easy. 

## How does it work?

Given a docker compose file, dcm will manage the start -> stop -> rebuild -> start lifecycle of your development containers.

This is managed by a unix-cron string format to perform the rebuild on your desired schedule.

Annotations in your docker compose file will be used to determine whether a service is ephemeral or persistent

**Ephemeral** - All services are assumed to be ephemeral - meaning that `dcm` will execute on the schedule of your choosing and perform either a pull or build of your development containers.

**Persistent** - Services with the `dcm: persistent` annotation will be started and managed by `dcm` but will not be rebuilt or modified on the established schedule.

## Why?
Use dcm to control ensuring you always have the latest packages installed in your development containers. It also is a forcing-function to ephemeral development - a best practice in current development whereby the only tooling that exists at the start are those declared in your image build process.



