-- name: GetAllOrganisations :many
SELECT *
FROM organisations;

-- name: GetOrganisation :one
SELECT *
FROM organisations
WHERE org_id = @org_id;

-- name: GetOrganisationWithChild :many
with recursive query_test
as (
	select org_id, "name", account, website, fuel_policy, fuel_set_by, speed_policy, speed_set_by, parent_id
	from organisations o
	where o.org_id = @org_id
	union all 	
	select o.org_id, o."name", o.account, o.website, o.fuel_policy, o.fuel_set_by, o.speed_policy, o.speed_set_by, o.parent_id 
	from organisations o 
	join query_test q on o.parent_id = q.org_id)
select * from query_test;

-- name: CreateOrganisation :one
INSERT INTO organisations (name, account, website, fuel_policy, speed_policy, parent_id)
Values (@name, @account, @website, @fuel_policy, @speed_policy, @parent_id)
RETURNING *;

-- name: CreateOrganisationThroughParent :one
INSERT INTO organisations (name, account, website, fuel_policy, fuel_set_by, speed_policy, parent_id)
Values (@name, @account, @website, @fuel_policy, @parent_id, @speed_policy, @parent_id)
RETURNING *;

-- name: UpdateOrganisation :one
UPDATE organisations
SET account = @account,
website = @website,
fuel_policy = @fuel_policy,
speed_policy = @speed_policy
WHERE org_id = @org_id
RETURNING *;

-- name: UpdateFuelPolicy :execrows
with recursive query_test
as (
	select org_id, "name", account, website, fuel_policy, fuel_set_by, speed_policy, speed_set_by, parent_id
	from organisations o
	where org_id = @org_id
	union all 	
	select o.org_id, o."name", o.account, o.website, o.fuel_policy, o.fuel_set_by, o.speed_policy, o.speed_set_by, o.parent_id 
	from organisations o 
	join query_test q on o.parent_id = q.org_id)
update organisations 
set fuel_policy = @fuel_policy,
fuel_set_by = @org_id
where org_id in (select org_id from query_test);