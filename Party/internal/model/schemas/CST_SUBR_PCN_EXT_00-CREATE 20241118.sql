create table cst_subr_pcn_ext_00 (
	subr_row_id char(15) not null ,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
	subr_ref_id varchar(30)  ,
	acct_ref_id varchar(30)  ,
	pmnt_chnl_prim varchar(15)  ,
	true_subr_type varchar(10)  ,
	subr_id varchar(15)  ,
	true_subr_stts varchar(5)  ,
	subr_crlm float  ,
	totl_obli float  ,
	totl_futr_prce float  ,
	home_telp_numb char(21) ,
	name_type varchar(10)  ,
	sale_id varchar(20)  ,
	subr_genr_lang char(2)  ,
	subr_sms_indc char(1)  ,
	subr_msim_indc char(1)  ,
	rel_subr varchar(30)  ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,

	primary key (subr_row_id)
		constraint cstp_subr_pcn_ext_00_01
);

--==============================================================
-- Index: cstp_subr_pcn_ext_00_01
--==============================================================
create unique index cstp_subr_pcn_ext_00_01 on cst_subr_pcn_ext_00 (
	subr_row_id  ASC
);

--==============================================================
-- Index: cstp_subr_pcn_ext_00_02
--==============================================================
create unique index cstp_subr_pcn_ext_00_02 on cst_subr_pcn_ext_00 (
	cust_numb, subr_numb  ASC
);

--==============================================================
-- Index: cstp_subr_pcn_ext_00_03
--==============================================================
create index cstp_subr_pcn_ext_00_03 on cst_subr_pcn_ext_00 (
	ext_id  ASC
);

--==============================================================
-- Index: cstp_subr_pcn_ext_00_04
--==============================================================
create index cstp_subr_pcn_ext_00_04 on cst_subr_pcn_ext_00 (
	bl_ext_id  ASC
);


