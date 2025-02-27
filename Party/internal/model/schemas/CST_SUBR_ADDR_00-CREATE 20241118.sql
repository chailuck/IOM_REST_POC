create table CST_SUBR_ADDR_00 (
	addr_row_id char(15) not null ,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	addr_type char(2)  not null,
	subr_row_id char(15) not null,
	adr1 varchar(80) not null,
	adr2 varchar(80)  ,
	cnty_code char(4)  not null,
	addr_post_code char(8)  not null,
	pscd_seqn_numb smallint  not null,
	addr_pcxt char(5)   ,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

	primary key (addr_row_id)
		constraint cstp_subr_addr_01

);


--==============================================================
-- Index: cstp_subr_addr_00_01
--==============================================================
create unique index cstp_subr_addr_00_01 on cst_subr_addr_00 (
	addr_row_id  ASC
);

--==============================================================
-- Index: cstn_subr_addr_00_02
--==============================================================
create unique index cstn_subr_addr_00_02 on cst_subr_addr_00 (
	cust_numb, subr_numb, addr_type  ASC
);

--==============================================================
-- Index: cstn_subr_addr_00_03
--==============================================================
create unique index cstn_subr_addr_00_03 on cst_subr_addr_00 (
	subr_row_id, addr_type ASC
);

--==============================================================
-- Index: cstn_subr_addr_00_04
--==============================================================
create index cstn_subr_addr_00_04 on cst_subr_addr_00 (
	pscd_seqn_numb, addr_post_code, cnty_code ASC
);
