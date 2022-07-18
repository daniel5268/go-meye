# GoMeye
As a meye player and a person who is always obsessed about following rules, I found that keeping your spend XP and stat points syncronized was a challenge, the chronicle master had it hard to check every player stats and validate that those points were spend correctly and there were no discrepancies.

GoMeye is the project meant to serve as the Meye Digital Assistand Backend, an initiative build in order to kill those discrepancies from the root, by allowing to the chronicle master setting xp assignations and to the players to spend those assignations, automating the calculations that were made by hand and reducing to 0 the possible errors in this process.

It allows to manage users, pjs and xp assignations.

## Requirements
 - [Go, > v1.18][1]
 - [Postgres, > v14.3][2]
 - [golang-migrate CLI][3]

## Instalation
In order to start the installation process the requirements should be already installed
 - Create a postgres database
 - Rename files ending in .env.example and remove .example.
 - set your different environment variables in the previous step files, DATABASE_URL value is needed in the next step
 - Run "migrate -database "{DATABASE_URL}" -path ./migrations up" (do not include brackets on DATABASE_URL) for each environment
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

## Errors
Every package has the responsability to map returned errors to DomainErrors, if an error is originated in a package, it should return a DomainError, if it is returning an error originated in other package, it relies on the fact that it was a DomainError and returns it directly.

## Tests
while testing, test.env will be used if LoadConfig(config.Test) is called. it's recomended to use a different database for testing as this may be truncated for testing purposes.

- go test ./...

### Pull requests guidelines
In order to upload a pull request please write to dataech@gmail.com with subjet "gomeye PR proposal" explaining the changes that would be inlcuded in case of approval and how those can improve the proyect, this Email will be replied as soon as possible to start a discussion (if needed) to refine the feature or with an OK to continue.

When the PR proposal is approved

- create a branch named "feature/addedFeature" from the branch "development"
- the first commit should have a message like "<type>(<scope>): <subject in imperative presend>" example: "feat(user): create base project structure; add endpoint to update an user". If several things need to be in the subject please split them by ; The rest of the commits in the PR branch will be squahed into this first commit on merge.



- push your changes and create a PR to the branch "development"

[1]:https://go.dev/doc/install
[2]:https://www.postgresql.org/download/
[3]:https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
[4]:dalenis@utp.edu.co
[5]:https://echo.labstack.com/
[6]:https://gorm.io/
[7]:https://www.mongodb.com/
