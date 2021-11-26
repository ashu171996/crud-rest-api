# Crud-rest-api
Golang CRUD Rest API using MySql database and gorilla MUX.

## Local setup
Install mysql and go 1.16v and execute the following query in mysql :-

`CREATE TABLE `projects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `country_code` varchar(25) DEFAULT NULL,
  `manager_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);`

## How to execute
After creating table we need to pass environment variable values into .env file by creating a new file .env in current directory and copy paste the variables from env-dist file.

At last run these go command in order to build file and run binary.

In currenct directory open terminal and type `go build`
Which will build the code and execute it by typing `./crud-rest-api`