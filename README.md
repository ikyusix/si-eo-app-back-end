# SIEO APP (Si Event Organizer App)

## Getting started
#### Setting up a project
- Move into your project directory: `cd ~/YOUR_PROJECTS_DIRECTORY`
- Clone this repository: `git clone https://git.enigmacamp.com/enigma-camp/class-turing-2/mini-project-organizer/kelompok-7/backend-go/si-EO_app.git`
- Move into the project director: `cd YOUR_PROJECT_NAME`
#### Working on the project
- Run the development task: `go run .`
    - Starting web server at port: [http://localhost:8586](http://localhost:8586)
- Use postman to test the API
> Note: **Event Organizer - EnigmaMiniProject.postman_collection.json**

### Database
By default, the back-end in configured to connect to a postgresSQL database
- Set `config.json`;
```json
"port":{
        "port": "8585"
    },
    "database":{
        "host": "localhost",
        "port": "5432",
        "user": "postgres",
        "pass": "root",
        "name": "sieo_db"
    }
``` 

### API routing
- USERS
    - GET ALL USERS
    - GET USER BY ID
    - EDIT USER BY ID
    - DELETE USER BY ID
    - USER LOGIN
    - USER SIGN UP
    - UPLOAD IMAGE PROFILE BY ID
- EO
    - GET ALL EO
    - INSERT EO
    - EDIT EO
- EVENT
    - GET ALL EVENT
    - INSERT EVENT
    - GET EVENT BY ID
### Changelog
- **v1.0**
> Note: Open testing application


