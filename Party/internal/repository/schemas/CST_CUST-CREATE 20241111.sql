create table CST_CUST (
	ca_row_id char(15) not null,
	cust_numb integer  not null,
	paty_row_id char(15) ,
	blpd_code char(2)  not null,
	comp_code char(2)  not null,
	mgrt_from_cust integer,
	cust_stts char(1)  not null,
	titl varchar(40,10)   ,
	frst_name nvarchar(80,40)  not null,
	last_name nvarchar(80,40)   ,
	id_type char(2)  not null,
	id_numb char(20)  not null,
	old_id_numb char(20)   ,
	gndr char(1)   ,
	date_of_brth date   ,
	occp_code char(3)   ,
	ntnt_code char(4)   ,
	marl_stts char(1)   ,
	chld smallint,
	faml_memb smallint	, 
	salr_levl char(1)   ,
	incl_dict char(1)   ,
	good_stts char(1)   ,
	lang char(1)  not null,
	grup_head char(1)   ,
	leaf_cust char(1)   ,
	pret_cust_numb integer,  
	grup_code integer default 0 not null,
	grup_levl smallint default 0 not null,
	empl_flag char(1) default 'N'  ,
	rprt_levl_flag char(1) default 'N' not null,
	pmnt_levl_flag char(1) default 'N' not null,
	grup_subr_indc char(1)  not null,
	home_telp_numb char(21)   ,
	home_fax_numb char(21)   ,
	offc_telp_numb char(21)   ,
	offc_fax_numb char(21)   ,
	crcr_telp_numb varchar(21)   ,
	py_telp_numb char(21)   ,
	emal_addr varchar(40)   ,
	fcbk_id varchar(40)   ,
	fcbk_lgin varchar(40)   ,
	docm_addr_type char(2)   ,
	acct_type char(2)   ,
	acct_sub_type char(2)   ,
	jrst_prsn_flag char(1)   ,
	brnc_code varchar(150)   ,
	crtd_dttm datetime year to second  not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

primary key (ca_row_id)
    constraint cstp_cust_01
);

--==============================================================
-- Index: cstp_cust_01
--==============================================================
create unique index cstp_cust_01 on cst_cust (
	ca_row_id  ASC
);


--==============================================================
-- Index: cstn_cust_02
--==============================================================
create unique index cstn_cust_02 on cst_cust (
	cust_numb ASC
);


--==============================================================
-- Index: cstn_cust_03
--==============================================================
create  index cstn_cust_03 on cst_cust (
	pret_cust_numb ASC
);


--==============================================================
-- Index: cstn_cust_04
--==============================================================
create  index cstn_cust_04 on cst_cust (
	parn_ca_row_id ASC
);


--==============================================================
-- Index: cstn_cust_05
--==============================================================
create  index cstn_cust_05 on cst_cust (
	cust_stts ASC
);

--==============================================================
-- Index: cstn_cust_06
--==============================================================
create  index cstn_cust_06 on cst_cust (
	blpd_code ASC
);

--==============================================================
-- Index: cstn_cust_07
--==============================================================
create  index cstn_cust_07 on cst_cust (
	frst_name ASC
);

--==============================================================
-- Index: cstn_cust_08
--==============================================================
create  index cstn_cust_08 on cst_cust (
	last_name ASC
);

--==============================================================
-- Index: cstn_cust_09
--==============================================================
create  index cstn_cust_09 on cst_cust (
	id_numb, cust_stts ASC
);


--==============================================================
-- Index: cstn_cust_10
--==============================================================
create  index cstn_cust_10 on cst_cust (
	id_type ASC
);

--==============================================================
-- Index: cstn_cust_11
--==============================================================
create  index cstn_cust_11 on cst_cust (
	paty_row_id ASC
);
