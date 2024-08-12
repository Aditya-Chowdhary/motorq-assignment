-- name: GetVehicle :one
SELECT *
FROM vehicles
WHERE vin = @vin;

-- name: CreateVehicle :one
insert into vehicles (vin, org_id, manufacturer, make, model, year)
values (@vin, @org_id, @manufacturer, @make, @model, @year)
RETURNING *;