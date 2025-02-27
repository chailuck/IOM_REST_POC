create table CST_SPKD_ATTR_00 (
	attr_row_id char(15)   ,
	seqn_spkd integer  not null,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	pack_code char(8) not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
	attr_type char(8)  ,
	attr_name varchar(30)  ,
	attr_vlue varchar(150)  ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

	primary key (attr_row_id)
		constraint cstp_spkd_attr_00_01
);

--==============================================================
-- Index: cstp_spkd_attr_00_01
--==============================================================
create unique index cstp_spkd_attr_00_01 on cst_spkd_attr_00 (
	attr_row_id  ASC
);

--==============================================================
-- Index: cstn_spkd_attr_00_02
--==============================================================
create unique index cstn_spkd_attr_00_02 on cst_spkd_attr_00 (
	seqn_spkd, attr_name  ASC
);

--==============================================================
-- Index: cstn_spkd_attr_00_03
--==============================================================
create index cstn_spkd_attr_00_03 on cst_spkd_attr_00 (
	cust_numb, subr_numb, pack_code  ASC
);

--==============================================================
-- Index: cstn_spkd_attr_00_04
--==============================================================
create index cstn_spkd_attr_00_04 on cst_spkd_attr_00 (
	pack_code  ASC
);

--==============================================================
-- Index: cstn_spkd_attr_00_05
--==============================================================
create index cstn_spkd_attr_00_05 on cst_spkd_attr_00 (
	ext_id  ASC
);

--==============================================================
-- Index: cstn_spkd_attr_00_06
--==============================================================
create index cstn_spkd_attr_00_06 on cst_spkd_attr_00 (
	bl_ext_id  ASC
);