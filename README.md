# GoMeye
As a Meye player and a person who is always obsessed about following rules, I found that keeping your spent XP and stat points syncronized was a challenge, the chronicle master had it hard to check every player's stats and validate that those points were spent correctly.

GoMeye is the project meant to serve as the Meye Digital Assistand Backend, an initiative build in order to kill those discrepancies from the root, by allowing to the chronicle master setting xp assignations and to the players to spend those, automating the calculations that were made by hand and reducing to 0 the possible errors in this process.

There are three roles:
 - admin: role in charge of creating, updating and deleting users, this role can be used to reset an user secret
 - master: role in charge of assigning XP to the different PJs
 - player: role in charge of creating and updating own PJs

## Requirements
 - [Go, >= v1.18][1]
 - [Postgres, >= v14.3][2]
 - [golang-migrate CLI][3]

## Installation
 - Create a postgres database
 - Rename files ending in .env.example and remove .example.
 - set your different environment variables in the previous step files, DATABASE_URL value is needed in the next step
 - Run "migrate -database "{DATABASE_URL}" -path ./migrations up" (do not include brackets on DATABASE_URL) for each environment, migrations should be run in every environment
 - Run "GO_ENV=production go run ./src/main.go" this command will take environment variables defined in .env, to use environment variables from dev.env, omit GO_ENV=production

When migrations are run, an admin user is inserted with credentials username="admin" secret="league", It's recomended to update the secret with the PUT /api/gomeye/users/{userID} endpoint

## Proyect architecture
[Echo][5] was used as web framework, [Gorm][6] as ORM, [postgreSQL][2], [mongoDB][7] for PJs stats.  The packages are located in the src directory, every package is named as the containing folder, inside those folders there can be several files separated by entity.

 - Packages generalities
 - - api: package in charge containing the echo server, routing requests to different handlers and injecting dependencies to the components
 - - config: package in charge of providing configuration variables values, it exports env files definitions
 - - domain: exports the models, structs, methods(User, DomainError, etc) that are part of the domain logic, used by services
 - - handler: exports methods that are called by the different api endpoints, those methods usually validate inputs and call a service methods, those are wrapped by entity (User, Pj, XPAssignations). a dto.go file can be found in this package, this file contains the request structs needed in the inputs - outputs of the API 
 - - service: exports the different services, those are wrapped by entity (User, Pj, XPAssignations) and usually use structs and methods from domain package
 - - repository: exports methods to get, store, update and delete persisted data
 - - util: package with helper functions

## Errors
Every package has the responsability to map returned errors to DomainErrors, if an error is originated in a package, it should return a DomainError, if it is returning an error originated in other package, it relies on the fact that it was a DomainError and returns it directly.

## Tests
while testing, test.env will be used if LoadConfig(config.Test) is called. it's recomended to use a different database for testing as this may be truncated for testing purposes.

- go test ./...

### Pull requests guidelines
In order to upload a pull request please write to [Daniel Tamayo](mailto:dataech@gmail.com?subject=[Gomeye]%20PR%20Proposal) explaining the changes that would be inlcuded in case of approval and how those can improve the proyect, this Email will be replied as soon as possible to start a discussion (if needed) to refine the feature or with an OK to continue.

When the PR proposal is approved

- create a branch named "feature/added-feature" from the branch "development"
- the first commit should have a message like "{type}({scope}): {subject in imperative present}", The rest of the commits in the PR branch will be squahed into this first commit on merge. examples:
- - "feat(user): create base project structure; add endpoint to update an user"
- - "fix(pj): fix XPCalculations when updating a PJ"
- - "test(user): add integration tests for POST /users/token"
- - "docs(user): add openapi spec to users CRUD"
- - "refactor(handler): refactor validations to improve performance"
 If several things need to be in the subject please split them by ;



- push your changes and create a PR to the branch "development"

[1]:https://go.dev/doc/install
[2]:https://www.postgresql.org/download/
[3]:https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
[4]:dalenis@utp.edu.co
[5]:https://echo.labstack.com/
[6]:https://gorm.io/
[7]:https://www.mongodb.com/
