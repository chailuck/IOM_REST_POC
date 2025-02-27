create table cst_csad (
	ca_row_id char(15) not null,
	addr_row_id char(15) not null,
	cust_numb integer  not null,
	addr_type char(2)  not null,
	adr1 varchar(80)  not null,
	adr2 varchar(80)  not null,
	cnty_code char(4)  not null,
	addr_post_code char(8)  not null,
	pscd_seqn_numb smallint  not null,
	addr_pcxt char(5),
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

	primary key (addr_row_id)
		constraint cstp_csad_01
);

--==============================================================
-- Index: cstp_csad_01
--==============================================================
create unique index cstp_csad_01 on cst_csad (
	addr_row_id  ASC
);


--==============================================================
-- Index: cstn_csad_02
--==============================================================
create unique index cstn_csad_02 on cst_csad (
	cust_numb, addr_type ASC
);

--==============================================================
-- Index: cstn_csad_03
--==============================================================
create index cstn_csad_03 on cst_csad (
	ca_row_id, addr_type ASC
);

--==============================================================
-- Index: cstn_csad_04
--==============================================================
create index cstn_csad_04 on cst_csad (
	pscd_seqn_numb, addr_post_code, cnty_code ASC
);




