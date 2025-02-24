create table cst_cust_ext (
	ca_row_id char(15) not null,
	cust_numb integer  not null,
	bl_ext_id varchar(20)  ,
	ext_id varchar(70)  ,
	paty_row_id char(15)   ,
	parn_ca_row_id char(15)  ,
	name_type varchar(10)   ,
	agmt_refn_id varchar(30)   ,
	pymt_chnl_desc varchar(100)   ,
	pymt_chng_id varchar(30)  ,
	refr_id varchar(70) ,
	bl_arrg_id varchar(30)   ,
	prod_cnt integer   ,
	open_date datetime year to second   ,
	raw_acnt_id varchar(30)   ,
	crdt_clas char(4)  ,
	crdt_limt_wave_indc char(2)   ,
	pers_crdt_limt float   ,
	acct_sub_type char(6)   ,
	acct_prio char(6)   ,
	temp_crdt_limt float   ,
	pers_crdt_dttm datetime year to second   ,
	tax_id char(20)   ,
	bill_info_ext_id char(30)   ,
	bill_name_ext_id char(30)   ,
	bill_lang char(2)   ,
	bill_id_type char(2)  ,
	bill_id_numb char(20) ,
	py_telp_numb char(21)   ,
	offc_telp_numb char(21)   ,
	home_telp_numb char(21)   ,
	priv_telp_numb_flag char(1)   ,
	titl varchar(40,10)   ,
	frst_name nvarchar(80,40)  not null,
	last_name nvarchar(80,40)   ,
	bill_name_type varchar(10)   ,
	marl_stts char(1)   ,
	gndr char(1)   ,
	pymt_ext_id char(30)   ,
	pymt_mthd char(4)   ,
	pou_ext_id varchar(30)   ,
	pou_refr_id varchar(30)   ,
	pou_id varchar(10)   ,
	pou_nmae varchar(40)   ,
	pou_noof_idd integer   ,
	pou_noof_ir integer   ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

primary key (ca_row_id)
    constraint cstp_cust_ext_01
);

--==============================================================
-- Index: cstp_cust_ext_01
--==============================================================
create unique index cstp_cust_ext_01 on cst_cust_ext (
	ca_row_id  ASC
);

--==============================================================
-- Index: cstn_cust_ext_02
--==============================================================
create unique index cstn_cust_ext_02 on cst_cust_ext (
	cust_numb ASC
);

--==============================================================
-- Index: cstn_cust_ext_03
--==============================================================
create index cstn_cust_ext_03 on cst_cust_ext (
	parn_ca_row_id ASC
);

--==============================================================
-- Index: cstn_cust_ext_04
--==============================================================
create index cstn_cust_ext_04 on cst_cust_ext (
	bl_ext_id ASC
);

--==============================================================
-- Index: cstn_cust_ext_05
--==============================================================
create index cstn_cust_ext_05 on cst_cust_ext (
	ext_id ASC
);


--==============================================================
-- Index: cstn_cust_ext_06
--==============================================================
create index cstn_cust_ext_06 on cst_cust_ext (
	pt_row_id ASC
);
