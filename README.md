# Altea Data Storage

A simple API for file storage.

## Requirements
- go 1.18

## todo
- Make app configurable
  - host&port (currently 0.0.0.0:9000)
  - storage used (currently local, available InMemory)
- Add S3 storage
- Add SQL storage
- Add OpenApi (or RAML) Documentation

## Routes

### `GET /files/{path}`
List all files contained in folder

### `PUT /files/{path}`
Write the content of the body for the given path


### `DELETE /files/{path}`
Delete file/directory recursively

## defaults
Serves the public folder at the root of the project.