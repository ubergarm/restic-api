restic-api
===
*NOTE* this is pre-alpha, it will change and I may force push. You've been warned!

A RESTish API wrapper around [restic](https://github.com/restic/restic) CLI.

## Run
```bash
go run restic-api.go
```

## Test
```bash
http localhost:8080
```

## Performance
```bash
docker run --rm -it williamyeh/wrk -t4 -c1000 -d30s http://$DOCKER_HOST_IP:8080
```

## Todo
- [x] Proof Of Concept
- [ ] Use actual environment vars
- [ ] Implement listing snapshots
- [ ] Implement downloading files
- [ ] Dynamically specify password/repo in headers?
- [ ] Document Usage

## References
[restic](https://github.com/restic/restic)
