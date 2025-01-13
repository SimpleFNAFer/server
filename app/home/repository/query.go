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
		with ip_times as (select source_ip,
								 date_trunc('second', request.timestamp) as scd,
								 count(*)                                as total
						  from server.request
						  where source_ip not in (select ip from server.blocked_ip)
						  group by source_ip, date_trunc('second', request.timestamp))
		select source_ip
		from ip_times
		where total > 30
		group by source_ip;
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
