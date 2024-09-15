# RabbitMQ Go Learning Project

This repository contains examples and exercises for learning RabbitMQ with Go. It covers various messaging patterns and RabbitMQ concepts, providing practical implementations for each.

## Table of Contents

1. [Simple Sending and Receiving](#simple-sending-and-receiving)
2. [Work Queues](#work-queues)
3. [Publish/Subscribe](#publishsubscribe)
4. [Routing](#routing)
5. [Topics](#topics)

## Simple Sending and Receiving

Basic example of sending and receiving a "Hello World" message using RabbitMQ and Go.

## Work Queues

Demonstrates the concept of work queues (task queues) to distribute time-consuming tasks among multiple workers.

Key concepts covered:

-   Round-robin dispatching
-   Message acknowledgment
-   Message durability
-   Fair dispatch

## Publish/Subscribe

Implements a simple logging system to broadcast log messages to multiple receivers.

Key concepts covered:

-   Exchanges (specifically fanout exchanges)
-   Temporary queues
-   Bindings

## Routing

Extends the logging system to subscribe to a subset of messages based on severity.

Key concepts covered:

-   Direct exchanges
-   Multiple bindings

## Topics

Implements a more complex routing system based on multiple criteria.

Key concepts covered:

-   Topic exchanges
-   Wildcards in routing keys

## Running the Examples

Each section contains instructions on how to run the examples. Typically, you'll need to run separate processes for producers and consumers.

## Prerequisites

-   Go (1.15 or later recommended)
-   RabbitMQ server installed and running
