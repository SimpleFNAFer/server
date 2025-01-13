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
		select source_ip, count(*) as total from server.request
		where timestamp >= now() - interval '1 second'
		group by source_ip;
`

	queryGetFrequentIPs = `
		select source_ip
		from server.request 
		    join server.blocked_ip 
		        on request.source_ip = blocked_ip.ip
		where request.timestamp >= now() - interval '1 second'
		and blocked_ip.ip <> request.source_ip
		group by source_ip 
		having count(*) > 30;
`

	queryGetBlockedIPs = `
		select ip
		from server.blocked_ip;
`

	queryGetAdminPassword = `
		select password_hash 
		from server.admin
		where login = $1;
`

	queryBlockIP = `
		insert into server.blocked_ip
		(ip)
		values ($1)
		on conflict do nothing;
`

	queryUnblockIP = `
		delete 
		from server.blocked_ip 
		where ip = $1;
`
)
