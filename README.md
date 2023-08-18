## 构建自己的 Redis

### Project Introduction

Welcome to this engaging programming project! We will be creating a simplified version of Redis, a widely used in-memory data structure store, using Golang. Our version will be capable of handling basic commands such as storing and retrieving key-value pairs, deleting key-value pairs, checking if a key exists, listing all keys in the database, and adding an expiration time to a key-value pair. This project will offer a fantastic opportunity to learn about databases, data structures, and time management in programming. The key-value pairs that can be stored in the database will be of string data type, and the expiration time will be represented in seconds. We will also provide detailed information on the input and output forms of the project, especially for the functions that store, retrieve, delete, and list key-value pairs.

### Implementation Approach

We will be using the Singleton design pattern to ensure that only one instance of the database is used throughout the program. This will prevent data inconsistency and ensure that all operations are performed on the same database instance. We will also be using the time package in Golang to handle expiration times. The key-value pairs that can be stored in the database will be of string data type, and the expiration time will be represented in seconds. The functions will return appropriate messages or data based on the operation performed. We will also provide detailed information on the input and output forms of the functions.
