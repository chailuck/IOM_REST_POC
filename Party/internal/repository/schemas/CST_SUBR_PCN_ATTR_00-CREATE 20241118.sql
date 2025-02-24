create table cst_subr_attr_00 (
	attr_row_id char(15) not null ,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,	
	subr_row_id char(15)   ,
	attr_type char(8)  ,
	attr_name varchar(30) not null,
	attr_vlue varchar(150)  ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,
	
	primary key (attr_row_id)
		constraint cstp_subr_attr_01

);


--==============================================================
-- Index: cstp_subr_attr_00_01
--==============================================================
create unique index cstp_subr_attr_00_01 on cst_subr_attr_00 (
	attr_row_id  ASC
);

--==============================================================
-- Index: cstp_subr_attr_00_02
--==============================================================
create unique index cstp_subr_attr_00_02 on cst_subr_attr_00 (
	cust_numb, subr_numb, attr_name  ASC
);

--==============================================================
-- Index: cstp_subr_attr_00_03
--==============================================================
create index cstp_subr_attr_00_03 on cst_subr_attr_00 (
	ext_id  ASC
);

--==============================================================
-- Index: cstp_subr_attr_00_04
--==============================================================
create index cstp_subr_attr_00_04 on cst_subr_attr_00 (
	bl_ext_id  ASC
);

--==============================================================
-- Index: cstp_subr_attr_00_05
--==============================================================
create index cstp_subr_attr_00_05 on cst_subr_attr_00 (
	attr_name  ASC
);

--==============================================================
-- Index: cstp_subr_attr_00_06
--==============================================================
create unique index cstp_subr_attr_00_06 on cst_subr_attr_00 (
	subr_row_id, attr_name  ASC
);
--==============================================================
-- Index: cstp_subr_attr_00_07
--==============================================================
create index cstp_subr_attr_00_07 on cst_subr_attr_00 (
	attr_type  ASC
);
