# repository-design-pattern

The repository design pattern that has following characteristics

* Easy to register new repositories with minimal repetitive code
* Single database connection (aws session) across all repositories
* Parent struct that contains shared field
* Each repository must implement standard Restful methods, put can contain any additional functions

