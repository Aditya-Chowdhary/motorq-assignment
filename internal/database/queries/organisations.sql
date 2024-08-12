-- name: GetAllOrganisations :many
SELECT *
FROM organisations;

-- name: CreateOrganisation :one
INSERT INTO organisations (name, account, website, fuel_policy, speed_policy, parent_id)
Values (@name, @account, @website, @fuel_policy, @speed_policy, @parent_id)
RETURNING *;

-- name: UpdateOrganisation :one
UPDATE organisations
SET account = @account,
website = @website,
fuel_policy = @fuel_policy,
speed_policy = @speed_policy
WHERE org_id = @org_id
RETURNING *;