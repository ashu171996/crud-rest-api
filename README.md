# Crud-rest-api
Golang CRUD Rest API using MySql database and gorilla MUX.

## Local setup
Install docker into your system then run these steps into current directory:- 
  * `docker build -t crud-rest-api-1 .`
  * `docker run -d -p 52000:3306 --name crud-rest-api-1 -e MYSQL_ROOT_PASSWORD=supersecret crud-rest-api-1`

  Then check if your database is created or not
  * `docker exec -it crud-rest-api-1 bash`
  * `mysql -u root -p`
  Then type password and verify if it is created or not.


## How to execute
After creating table we need to pass environment variable values into .env file by creating a new file .env in current directory and copy paste the variables from env-dist file.

NOTE: mysql image is running on PORT=52000, user=root, databaseName=company

At last run these go command in order to build file and run binary.

In currenct directory open terminal and type `go build`
Which will build the code and execute it by typing `./crud-rest-api`
