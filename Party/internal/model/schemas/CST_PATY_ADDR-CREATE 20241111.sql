create table cst_paty_addr (
	addr_row_id char(15)   ,
	paty_row_id char(15)   ,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
	id_type char(2)  not null,
	id_numb char(20)  not null,
	addr_type char(2)  not null,
	adr1 varchar(80)  not null,
	adr2 varchar(80)   ,
	cnty_code char(4)  not null,
	addr_post_code char(8)  not null,
	pscd_seqn_numb smallint  ,
	addr_pcxt char(5)   ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

primary key (addr_row_id)
    constraint cstp_paty_addr_01
);

--==============================================================
-- Index: cstp_paty_addr_01
--==============================================================

create unique index cstp_paty_addr_01 on cst_paty_addr (
	addr_row_id  ASC
);

--==============================================================
-- Index: cstn_paty_addr_02
--==============================================================

create unique index cstn_paty_addr_02 on cst_paty_addr (
	id_numb, id_type, addr_type ASC
);

--==============================================================
-- Index: cstn_paty_addr_03
--==============================================================

create index cstn_paty_addr_03 on cst_paty_addr (
	bl_ext_id ASC
);

--==============================================================
-- Index: cstn_paty_addr_04
--==============================================================

create index cstn_paty_addr_04 on cst_paty_addr (
	ext_id ASC
);


