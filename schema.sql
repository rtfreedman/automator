create table if not exists tasks (
    id serial primary not null,
    jobName text unique not null,
    command text not null,
    lastRun timestamp with time zone not null default CURRENT_TIMESTAMP,
    timeRun time with time zone not null default CURRENT_TIME,
    interval int not null default 24
)