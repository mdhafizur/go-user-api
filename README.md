To start your Go application locally, follow these steps:

Install Go:
Make sure you have Go installed on your machine. You can download and install it from the official Go website: https://golang.org/dl/

Clone the Project:
Clone your Go project to your local machine using a version control tool like Git. Navigate to the directory where you want to store the project and run:

bash
Copy code
git clone <repository-url>
Replace <repository-url> with the URL of your Git repository.

Navigate to the Project Directory:
Change your current working directory to the root of your Go project:

bash
Copy code
cd user-api
Create a .env file:
Create a .env file in the project root and set the necessary environment variables. For example:

makefile
Copy code
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
MONGO_INITDB_DATABASE=godb
API_KEY=secretkey
DB_URL=mongodb://admin:admin@localhost:27017/godb
Adjust the values based on your MongoDB configuration.

Install Dependencies:
Run the following command to download and install the project dependencies:

bash
Copy code
go mod download
Run the Application:
Once the dependencies are installed, you can run your Go application using the following command:

bash
Copy code
go run cmd/user-api/main.go
This command starts the Go application, and it should be accessible at http://localhost:8080.

Remember to replace the MongoDB connection URL in the .env file with your actual MongoDB configuration.




#### To rerun tests without using the cache, you can use the -count flag with the go test command. The -count flag specifies the number of times to run the tests. By setting it to 1, you ensure that the tests are not cached. Here's the command:
```bash
go test -count=1 ./internal/app/test/...
```