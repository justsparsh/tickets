## Version

Go: 1.23.6
Cosmos SDK: 0.50.12

## Initializing

```
ignite chain serve
```

Take note of the location of the `ticketsd` binary. I will refer to it as `TICKETS_PATH`

## CRUD

`$NAME` means it is means to be a user input
`[--name=$NAME]` means this flag is optional

### Create a Ticket

```
TICKETS_PATH tx tickets create-ticket $NAME $EVENT --from alice
```

### List Tickets

```
TICKETS_PATH q tickets list-tickets [ --name=$NAME ] [ --event=$EVENT ]
```

### Fetch a Ticket

```
TICKETS_PATH q tickets show-ticket $ID
```

### Update a Ticket

```
TICKETS_PATH tx tickets update-ticket $NAME $EVENT $ID --from alice
```

## DELETE a Ticket

```
TICKETS_PATH tx tickets update-ticket $NAME $EVENT $ID --from alice
```
