-- Create keyspace
CREATE KEYSPACE IF NOT EXISTS am
WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

-- Use keyspace
USE am;

CREATE TABLE IF NOT EXISTS tokens (
    channel text,
    id uuid,
    secret text,
    groups text,
    created_at timestamp,
    expires_at timestamp,
    updated_at timestamp,
    
    PRIMARY KEY (groups, id)
);

-- Create
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('system',aeaaff1c-094f-4c7d-8371-8f6e1c153ecf,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','im',NULL);
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('system',2969c274-36f1-4f2f-b14b-9c7e13e246aa,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','am',NULL);
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('system',a90380b2-cf6f-4d16-83aa-2ab1b79d0483,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','am',NULL);
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('iservice',37b1381c-2bfd-458e-a8d1-754d1ee5ea2f,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','fm',NULL);
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('smartui',ab922259-9b80-412f-98d4-43edcd9c214b,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','fm',NULL);
INSERT INTO "iom"."tokens" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('smartui',b598836b-e43b-4595-9578-d19a21b93b75,'jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','fm',NULL);

-- ListAll
select * from tokens;
select * from tokens_by_secret;
-- GetByGroups
select * from tokens where groups = 'am';
-- GetById, DeleteByIdm, UpdateById
select * from tokens where groups = 'am' and id=2969c274-36f1-4f2f-b14b-9c7e13e246aa;
update tokens set secret='jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO' where groups = 'am' and id in(2969c274-36f1-4f2f-b14b-9c7e13e246aa);
update tokens set secret='jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO' where groups = 'fm' and id in(37b1381c-2bfd-458e-a8d1-754d1ee5ea2f,ab922259-9b80-412f-98d4-43edcd9c214b);
update tokens set secret='jPS6rTiaLLlWTO4dnYRML7G6pfhpgEzTPCt7FmzUK4dE1XO7WcA2Ka19CYb4-UwO' where groups = 'im' and id in(aeaaff1c-094f-4c7d-8371-8f6e1c153ecf);
delete from tokens where groups = 'system' and id in(468cfce9-eab7-11ef-a20c-00155d4cc564);
select * from tokens  where groups = 'system' and id in(468cfce9-eab7-11ef-a20c-00155d4cc564);


-- Create 
CREATE TABLE IF NOT EXISTS tokens_by_secret (
    channel text,
    id uuid,
    secret text,
    groups text,
    created_at timestamp,
    expires_at timestamp,
    updated_at timestamp,
    
    PRIMARY KEY (groups, secret)
);

INSERT INTO "iom"."tokens_by_secret" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('system',aeaaff1c-094f-4c7d-8371-8f6e1c153ecf,'api-key2','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','im',NULL);
INSERT INTO "iom"."tokens_by_secret" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('system',2969c274-36f1-4f2f-b14b-9c7e13e246aa,'api-key3','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','am',NULL);
INSERT INTO "iom"."tokens_by_secret" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('iservice',37b1381c-2bfd-458e-a8d1-754d1ee5ea2f,'api-key4','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','fm',NULL);
INSERT INTO "iom"."tokens_by_secret" ("channel","id","secret","created_at","expires_at","groups","updated_at") 
VALUES ('smartui',ab922259-9b80-412f-98d4-43edcd9c214b,'api-key5','2026-02-11 07:26:23.684+0000','2027-02-11 07:26:23.684+0000','fm',NULL);


-- GetBySecret
select * from tokens_2 where groups = 'fm' and secret='api-key5';


