CREATE TABLE IF NOT EXISTS public.organisations (
	org_id bigserial NOT NULL,
	"name" text NOT NULL,
	account text NOT NULL,
	website text NOT NULL,
	fuel_policy integer DEFAULT 1000 NOT NULL,
	fuel_set_by bigint DEFAULT currval('organisations_org_id_seq') NOT NULL, 
	speed_policy integer NOT NULL,
	speed_set_by bigint DEFAULT currval('organisations_org_id_seq') NOT NULL, 
	parent_id bigint NULL,
	CONSTRAINT organisations_pk PRIMARY KEY (org_id),
	CONSTRAINT organisations_organisations_fk FOREIGN KEY (parent_id) REFERENCES public.organisations(org_id) ON DELETE SET NULL ON UPDATE CASCADE
);