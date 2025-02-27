create table CST_CUST_ATTR (
	attr_row_id char(15) not null,
	cust_numb integer  not null,
	ca_row_id char(15) not null,
	ext_id varchar(70) ,
	bl_ext_id varchar(20)  ,
	attr_name varchar(20)  not null,
	attr_vlue varchar(40)  ,
	crtd_dttm datetime year to second  not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

primary key (attr_row_id)
    constraint cstp_cust_attr_01
);

--==============================================================
-- Index: cstp_cust_attr_01
--==============================================================
create unique index cstp_cust_attr_01 on cst_cust_attr (
	attr_row_id  ASC
);

--==============================================================
-- Index: cstn_cust_attr_02
--==============================================================
create unique index cstn_cust_attr_02 on cst_cust_attr (
	cust_numb, attr_name ASC
);
  
--==============================================================
-- Index: cstn_cust_attr_03
--==============================================================
create index cstn_cust_attr_03 on cst_cust_attr (
	attr_name ASC
);


--==============================================================
-- Index: cstn_cust_attr_04
--==============================================================
create index cstn_cust_attr_04 on cst_cust_attr (
	ext_id ASC
);

--==============================================================
-- Index: cstn_cust_attr_05
--==============================================================
create index cstn_cust_attr_05 on cst_cust_attr (
	bl_ext_id ASC
);

--==============================================================
-- Index: cstn_cust_attr_06
--==============================================================
create unique index cstn_cust_attr_06 on cst_cust_attr (
	ca_row_id, attr_name ASC
);