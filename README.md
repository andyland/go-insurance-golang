#Usage

##Setup

1. Create a postgres database on your machine 
```
$ createdb somedbname
```
2. Add the following to your ~/.bash_profile
```
export GO_INSURANCE_TEST_PORT="5000"
```
```export GO_INSURANCE_TEST_PGURL="postgres://yourusername:yourpassword@localhost:5432/somedbname?sslmode=disable"
```

## Endpoints

###/users/:username/appointments

#####POST

Creates an appointment

Requires:
* date (yyyy-mm-dd format)
* time\_of_day

#####GET

Returns an appointment for the user, or 404 if there isn't one


#####DELETE

Deletes the appointment for that user
