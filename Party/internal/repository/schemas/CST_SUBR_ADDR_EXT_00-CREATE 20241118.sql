create table CST_SUBR_ADDR_EXT_00 (
	addr_row_id char(15) not null,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	addr_type char(2)  not null,
	subr_row_id char(15) not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
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
		constraint cstp_subr_addr_ext_00_01
);

--==============================================================
-- Index: cstp_subr_addr_ext_00_01
--==============================================================
create unique index cstp_subr_addr_ext_00_01 on cst_subr_addr_ext_00 (
	addr_row_id  ASC
);


--==============================================================
-- Index: cstn_subr_addr_ext_00_02
--==============================================================
create unique index cstn_subr_addr_ext_00_02 on cst_subr_addr_ext_00 (
	cust_numb, subr_numb, addr_type ASC
);

--==============================================================
-- Index: cstn_subr_addr_ext_00_03
--==============================================================
create index cstn_subr_addr_ext_00_03 on cst_subr_addr_ext_00 (
	ext_id ASC
);

--==============================================================
-- Index: cstn_subr_addr_ext_00_04
--==============================================================
create index cstn_subr_addr_ext_00_04 on cst_subr_addr_ext_00 (
	bl_ext_id ASC
);

--==============================================================
-- Index: cstn_subr_addr_ext_00_02
--==============================================================
create unique index cstn_subr_addr_ext_00_02 on cst_subr_addr_ext_00 (
	subr_row_id, addr_type ASC
);
