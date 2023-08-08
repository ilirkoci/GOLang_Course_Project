## Project Setup

Create a directory ``task_management_system`` . This will be the HOME directory of the project.

Next, ``cd`` into this directory and run 

~~~go 
go mod init tms.zinkworks.com
~~~

The above command will enable modules for the project. The ``go mod init`` command expects a ``module path``, which is a unique identifier for the project. We will use ``tms.zinkworks.com`` as our module path.

You will se a ``go.mod`` file has been created and it look similar to this:

~~~go
module tms.zinkworks.com

go 1.20
~~~

### Managing Dependencies with ``go.mod`` in Go Projects

1. If a valid ``go.mod`` file is present at the root of your project directory, your project is recognized as a module.
2. When you're working within your project directory and use ``go get`` to download a dependency, the specific version of the dependency will be recorded in the go.mod file.
3. This approach ensures reproducible builds across various machines and environments because the exact versions are known and managed in the go.mod file.
4. When you run or build the code in your project, Go will utilize the precise dependencies listed in the go.mod file. In case the required dependencies are not already present on your local machine, Go will automatically download them along with any recursive dependencies.
5. Additionally, the go.mod file defines the module path (e.g., tms.zinkworks.com in our case), which acts as the root import path for the packages within your project.

### Directory Structure for the Project

Go to home directory and create two more directory inside it.

``model`` - this will contain all the models

``tms\api`` - this will contain all the APIs

At this point you will see below directory structure along will few additional files that will be explained later.

![Alt text](image.png)

### Setting a HTTP Server

See ``main.go`` file

### Setting a PostgreSQL Server

Install PostgreSQL Server from here:
https://sbp.enterprisedb.com/getfile.jsp?fileid=1258629

Set the password for ``postgres`` user and change this accordingly.
~~~go
flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:ambrish@localhost/tms?sslmode=disable", "PostgreSQL Connection")
~~~

So the format is:
postgres://username:password@localhost/databasename?sslmode=disable

After installing PostgreSQl, run SQL commands from these files:

1. create_table.sql
2. insert_table.sql

### Integrating ``httprouter`` package

When developing a Go API with endpoints, you may encounter limitations in the standard ``http.ServeMux`` router. This default router lacks features such as routing requests to different handlers based on request methods (e.g., GET, POST) and handling clean URLs with interpolated parameters.

To overcome these limitations, we'll integrate the popular ``httprouter`` package into our application. ``httprouter`` is a well-tested and stable router that meets our requirements. Moreover, it offers excellent performance due to its use of a radix tree for URL matching. If you plan to build a REST API for public use, ``httprouter`` is a reliable choice.

#### Comparing httprouter with Java EE and Spring for Building Go APIs

**Familiarity with Servlet-like Routing:** Developers familiar with Java EE or Spring are accustomed to working with powerful routing mechanisms for handling different HTTP methods and clean URLs with interpolated parameters. httprouter provides similar functionality, making it easy for them to adapt and work with in a Go environment.

**Enhanced Functionality:** While the Go standard library's http.ServeMux has certain limitations, httprouter overcomes these limitations by offering more advanced features like routing requests based on HTTP methods (GET, POST, etc.) and supporting clean URLs with interpolated parameters. This allows developers to build more sophisticated APIs and handle various types of requests efficiently.

We will use ``go get`` command for downloading the latest release of httprouter package.

~~~go
go get github.com/julienschmidt/httprouter@v1
~~~

This will also update the ``go.mod`` file and add the ``httprouter`` dependency.

~~~go
module tms.zinkworks.com

go 1.20

require github.com/julienschmidt/httprouter v1.3.0
~~~

### Defining the ``Task`` Model

See ``task.go`` file

## Running the Project

~~~go
go run .\tms\api
~~~

## Endpoints

See ``routes.go``

### Updating ``go.sum`` file

To update the ``go.sum`` file after removing dependencies, you can use below command:

~~~go
go mod tidy
~~~

This command will remove any entries from the ``go.sum`` file that are no longer needed based on your updated ``go.mod`` file.

## Use Case - Task Assignment:

a). Allow users to assign tasks to specific team members or collaborators.

b). Add an endpoint to retrieve tasks assigned to a particular user.

### To Be Implemented

1. You'll need to modify the existing task table to include a new column for storing the user ID of the team member to whom the task is assigned.
2. You'll need to add a new API endpoint that allows users to assign tasks to specific team members. This endpoint should receive the task ID and the user ID of the assignee as input.
3. You'll also need to add an endpoint that allows users to retrieve tasks assigned to a particular team member. This endpoint should receive the user ID as input and return a list of tasks assigned to that user.
4. You'll need to update the Task details wherever applicable and it should show one addition field Assigned To.


## Use Case - Task Comments:

a). Allow users to add comments to tasks to provide updates or additional information.

b). Implement an endpoint to retrieve task comments and display them in the task details view.

### To Be Implemented

1. Create a new table to store task comments. The table could be named task_comment and could have the following columns: id (primary key), task_id (foreign key to reference the task the comment belongs to), comment (the actual comment text), and created_at (timestamp for when the comment was added).
2. Add a new API endpoint that allows users to add comments to a task. This endpoint should receive the task ID and the comment text as input, and then insert the comment into the task_comment table.
3. Add a new API endpoint that allows users to retrieve all comments for a specific task. This endpoint should receive the task ID as input and return a list of comments associated with that task.

## Unit tests:

1. Add unit tests to all your code you have created so far.
2. Aim to achieve 80% code coverage.

## Integration tests:

1. Add integration tests to test your API functionality.

### Tips
- Have a clear understanding of what your API endpoint should do before writing your test.
- Include all aspects of functionality including any input validation and error handling.

## Docker deployment:

Create your Dockerfile. Don't forget to:
- Specify your parent image
- Create the app directory for your app's source code.
- Copy everything from your root directory into your work directory.
- Install Go dependencies.
- Build your go app with any configuration.
- Expose your container port(s).
- Execute the command to run your go app when the container starts.
