CREATE TABLE IF NOT EXISTS public.vehicles (
	vehicle_id bigserial NOT NULL,
	vin text NOT NULL,
	org_id bigint NOT NULL,
	manufacturer text NOT NULL,
	make text NOT NULL,
	model text NOT NULL,
	"year" integer NOT NULL,
	CONSTRAINT vehicles_pk PRIMARY KEY (vehicle_id),
	CONSTRAINT vehicles_unique UNIQUE (vin),
	CONSTRAINT vehicles_organisations_fk FOREIGN KEY (org_id) REFERENCES public.organisations(org_id) ON DELETE CASCADE ON UPDATE CASCADE
);