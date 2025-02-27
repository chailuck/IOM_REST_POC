create table cst_paty_addr_ext (
	addr_row_id char(15)  not null ,
	id_type char(2)  not null,
	id_numb char(20)  not null,
	addr_type char(2)  not null,
	ampr_name char(40)   ,
	buld_name char(40)   ,
	city_name char(40)   ,
	home_no char(20)   ,
	moo char(10)   ,
	strt_name char(40)   ,
	time_at_addr char(40)   ,
	tmbl_name char(10)   ,
	accm_type char(10)   ,
	zip_code char(10)   ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

primary key (addr_row_id)
    constraint cstp_paty_addr_ext_01
);

--==============================================================
-- Index: cstp_paty_addr_ext_01
--==============================================================

create unique index cstp_paty_addr_ext_01 on cst_paty_addr_ext (
	addr_row_id  ASC
);

--==============================================================
-- Index: cstn_paty_addr_ext_02
--==============================================================

create unique index cstn_paty_addr_ext_02 on cst_paty_addr_ext (
	id_numb, id_type, addr_type ASC
);

