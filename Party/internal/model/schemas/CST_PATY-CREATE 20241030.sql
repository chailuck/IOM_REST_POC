create table CST_PATY (
	paty_row_id char(15) not null ,
	id_type char(2) not null,
	id_numb char(20) not null,
	lang char(1)  not null,
	home_telp_numb char(21) ,
	titl nvarchar(40,10) ,
	frst_name nvarchar(80,40) not null,
	last_name nvarchar(80,40) ,
	marl_stts char(1)  ,
	gndr char(1)  ,
	name_type varchar(10)  ,
	cntc_lang char(3)  ,
	occp_code char(3)  ,
	time_in_buss varchar(10)  ,
	salr_levl char(1)  ,
	ntnt_code char(4)  ,
	init_time_in_addr varchar(10)  ,
	grde char(1)  ,
	id_exp_date date ,
	date_of_brth date ,
	lrge_cust_indc char(2)  ,
	sub_type varchar(5)  ,
	paty_type varchar(5)  ,
	src_cust_type varchar(5)  ,
	parn_titl nvarchar(40,10) ,
	parn_frst_name nvarchar(80,40)  ,
	parn_last_name nvarchar(80,40)  ,
	parn_cntc char(21)  ,
	parn_id char(21)  ,
	totl_subs integer  ,
	max_allw_only_subs integer  ,
	deal_pool varchar(20),
	crtd_dttm datetime year to second  not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

primary key (paty_row_id)
    constraint cstp_paty_01
);

--==============================================================
-- Index: cstp_paty_01
--==============================================================

create unique index cstp_paty_01 on cst_paty (
	paty_row_id  ASC
);

--==============================================================
-- Index: cstn_paty_02
--==============================================================

create unique index cstn_paty_02 on cst_paty (
	id_numb, id_type  ASC
);
 
--==============================================================
-- Index: cstn_paty_03
--==============================================================

create   index cstn_paty_03 on cst_paty (
	frst_name  ASC
);
 
--==============================================================
-- Index: cstn_paty_04
--==============================================================

create   index cstn_paty_04 on cst_paty (
	last_name  ASC
);
 
--==============================================================
-- Index: cstn_paty_05
--==============================================================

create   index cstn_paty_05 on cst_paty (
	paty_type  ASC
);
 
--==============================================================
-- Index: cstn_paty_06
--==============================================================

create   index cstn_paty_06 on cst_paty (
	sub_type ASC
);
 
--==============================================================
-- Index: cstn_paty_07
--==============================================================

create   index cstn_paty_07 on cst_paty (
	src_cust_type  ASC
);
