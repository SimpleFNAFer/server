package repository

const (
	querySaveRequest = `
		insert into server.request
		(source_ip)
		values ($1);
`

	queryCountRequests = `
		select count(*) from server.request;
`

	queryCountLastSecondRequest = `
		select count(*) from server.request
		where timestamp >= now() - interval '1 second';
`
)
