# automator
## Postgres Schema Description
```
    id serial primary not null, # primary key
    jobName text unique not null, # name of the job running
    command text not null, # command to run
    lastRun timestamp with time zone not null default CURRENT_TIMESTAMP, # the last timestamp this command was run
    timeRun time with time zone not null default CURRENT_TIME, # the time of day to run the command
    interval int not null default 24 # how frequently (in hours) we want to run this command
```